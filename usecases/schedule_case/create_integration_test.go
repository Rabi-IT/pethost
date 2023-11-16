package schedule_case_test

import (
	"net/http"
	"pethost/fixtures"
	"pethost/frameworks/database/gateways/schedule_gateway"
	"pethost/usecases/pet_case"
	"pethost/usecases/pet_case/pet"
	"pethost/usecases/preference_case"
	"pethost/usecases/schedule_case"
	"pethost/usecases/schedule_case/schedule_status"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func True() *bool {
	t := true
	return &t
}

func False() *bool {
	t := false
	return &t
}

func Test_Integration_Create(t *testing.T) {
	type testCase struct {
		title      string
		statusCode int
		seed       func() (fixtures.CreateDefaultOutput, int)
		expected   func(scenario fixtures.CreateDefaultOutput) schedule_gateway.PaginateData
	}

	tests := []testCase{
		{
			title: "happy path",
			seed: func() (fixtures.CreateDefaultOutput, int) {
				scenario := fixtures.Preference.CreateDefault(t, nil)
				_, statusCode := fixtures.Schedule.Create(
					t,
					schedule_case.CreateInput{
						PetIDs: []string{scenario.PetID},
						HostID: scenario.HostID,
						Dates: []schedule_gateway.CreateDate{
							{
								MonthYear:   time.Date(2023, 0, 0, 0, 0, 0, 0, time.UTC),
								DaysOfMonth: fixtures.Preference.AllDaysOfMonth,
							},
						},
						Notes: "Notes",
					},
					scenario.TutorToken,
				)

				return scenario, statusCode
			},
			expected: func(scenario fixtures.CreateDefaultOutput) schedule_gateway.PaginateData {
				return schedule_gateway.PaginateData{
					MonthYear:   time.Date(2023, 0, 0, 0, 0, 0, 0, time.UTC),
					DaysOfMonth: fixtures.Preference.AllDaysOfMonth,
					Notes:       "Notes",
					TutorID:     scenario.TutorID,
					PetID:       scenario.PetID,
					Status:      schedule_status.Open,
				}
			},
			statusCode: http.StatusCreated,
		},

		{
			title: "should not schedule if tutor is not the owner of the pet",
			seed: func() (fixtures.CreateDefaultOutput, int) {
				scenario := fixtures.Preference.CreateDefault(t, nil)
				NOT_TUTOR_PET := fixtures.Pet.Create(t, nil, scenario.HostToken)
				_, statusCode := fixtures.Schedule.Create(
					t,
					schedule_case.CreateInput{
						PetIDs: []string{NOT_TUTOR_PET},
						HostID: scenario.HostID,
						Dates: []schedule_gateway.CreateDate{
							{
								MonthYear:   time.Date(2023, 0, 0, 0, 0, 0, 0, time.UTC),
								DaysOfMonth: fixtures.Preference.AllDaysOfMonth,
							},
						},
						Notes: "Notes",
					},
					scenario.TutorToken,
				)

				return scenario, statusCode
			},
			expected: func(scenario fixtures.CreateDefaultOutput) schedule_gateway.PaginateData {
				return schedule_gateway.PaginateData{}
			},
			statusCode: http.StatusNotFound,
		},

		{
			title: "should not schedule if availability not meet tutor needs",
			seed: func() (fixtures.CreateDefaultOutput, int) {
				var HOST_AVAILABILITY uint32 = 0b00111
				scenario := fixtures.Preference.CreateDefault(t, &preference_case.CreateInput{
					OnlyVaccinated:          True(),
					AcceptElderly:           True(),
					AcceptOnlyNeuteredMales: False(),
					AcceptFemales:           True(),
					DaysOfMonth:             HOST_AVAILABILITY,
					AcceptFemaleInHeat:      True(),
					AcceptPuppies:           True(),
					AcceptMales:             True(),
					PetWeight:               fixtures.Preference.AllPetWeight,
				})

				var TUTOR_NEEDS uint32 = 0b11111
				_, statusCode := fixtures.Schedule.Create(
					t,
					schedule_case.CreateInput{
						PetIDs: []string{scenario.PetID},
						HostID: scenario.HostID,
						Dates: []schedule_gateway.CreateDate{
							{
								MonthYear:   time.Date(2023, 0, 0, 0, 0, 0, 0, time.UTC),
								DaysOfMonth: TUTOR_NEEDS,
							},
						},
					},
					scenario.TutorToken,
				)

				return scenario, statusCode
			},
			expected: func(scenario fixtures.CreateDefaultOutput) schedule_gateway.PaginateData {
				return schedule_gateway.PaginateData{}
			},
			statusCode: http.StatusNotFound,
		},

		{
			title: "should schedule if availability is greater than tutor needs",
			seed: func() (fixtures.CreateDefaultOutput, int) {
				var HOST_AVAILABILITY uint32 = 0b11111
				scenario := fixtures.Preference.CreateDefault(t, &preference_case.CreateInput{
					OnlyVaccinated:          True(),
					AcceptElderly:           True(),
					AcceptOnlyNeuteredMales: False(),
					AcceptFemales:           True(),
					DaysOfMonth:             HOST_AVAILABILITY,
					AcceptFemaleInHeat:      True(),
					AcceptPuppies:           True(),
					AcceptMales:             True(),
					PetWeight:               fixtures.Preference.AllPetWeight,
				})

				var TUTOR_NEEDS uint32 = 0b00111
				_, statusCode := fixtures.Schedule.Create(
					t,
					schedule_case.CreateInput{
						PetIDs: []string{scenario.PetID},
						HostID: scenario.HostID,
						Dates: []schedule_gateway.CreateDate{
							{
								MonthYear:   time.Date(2023, 0, 0, 0, 0, 0, 0, time.UTC),
								DaysOfMonth: TUTOR_NEEDS,
							},
						},
					},
					scenario.TutorToken,
				)

				return scenario, statusCode
			},
			expected: func(scenario fixtures.CreateDefaultOutput) schedule_gateway.PaginateData {
				var from1To3 uint32 = 0b111
				return schedule_gateway.PaginateData{
					MonthYear:   time.Date(2023, 0, 0, 0, 0, 0, 0, time.UTC),
					TutorID:     scenario.TutorID,
					DaysOfMonth: from1To3,
					PetID:       scenario.PetID,
					Status:      schedule_status.Open,
				}
			},
			statusCode: http.StatusCreated,
		},

		{
			title: "should not schedule if pet is large and host only accepts small pets",
			seed: func() (fixtures.CreateDefaultOutput, int) {
				scenario := fixtures.Preference.CreateDefault(t, &preference_case.CreateInput{
					OnlyVaccinated:          True(),
					AcceptElderly:           True(),
					AcceptOnlyNeuteredMales: False(),
					AcceptFemales:           True(),
					DaysOfMonth:             fixtures.Preference.AllDaysOfMonth,
					AcceptFemaleInHeat:      True(),
					AcceptPuppies:           True(),
					AcceptMales:             True(),
					PetWeight:               fixtures.Preference.OnlySmallPets,
				})

				newPetWeight := pet_case.PatchValues{Weight: fixtures.Pet.VeryLargePet}
				response, status := fixtures.Pet.Patch(t, scenario.PetID, newPetWeight, scenario.TutorToken)
				require.Equal(t, "OK", response)
				require.Equal(t, http.StatusOK, status)

				_, statusCode := fixtures.Schedule.Create(
					t,
					schedule_case.CreateInput{
						PetIDs: []string{scenario.PetID},
						HostID: scenario.HostID,
						Dates: []schedule_gateway.CreateDate{
							{
								MonthYear:   time.Date(2023, 0, 0, 0, 0, 0, 0, time.UTC),
								DaysOfMonth: fixtures.Preference.AllDaysOfMonth,
							},
						},
					},
					scenario.TutorToken,
				)

				return scenario, statusCode
			},
			expected: func(scenario fixtures.CreateDefaultOutput) schedule_gateway.PaginateData {
				return schedule_gateway.PaginateData{}
			},
			statusCode: http.StatusNotFound,
		},

		{
			title: "should not schedule if pet is small and host only accepts large pets",
			seed: func() (fixtures.CreateDefaultOutput, int) {
				scenario := fixtures.Preference.CreateDefault(t, &preference_case.CreateInput{
					OnlyVaccinated:          True(),
					AcceptElderly:           True(),
					AcceptOnlyNeuteredMales: False(),
					AcceptFemales:           True(),
					DaysOfMonth:             fixtures.Preference.AllDaysOfMonth,
					AcceptFemaleInHeat:      True(),
					AcceptPuppies:           True(),
					AcceptMales:             True(),
					PetWeight:               fixtures.Preference.OnlyLargePets,
				})

				newPetWeight := pet_case.PatchValues{Weight: fixtures.Pet.SmallPet}
				response, status := fixtures.Pet.Patch(t, scenario.PetID, newPetWeight, scenario.TutorToken)
				require.Equal(t, "OK", response)
				require.Equal(t, http.StatusOK, status)

				_, statusCode := fixtures.Schedule.Create(
					t,
					schedule_case.CreateInput{
						PetIDs: []string{scenario.PetID},
						HostID: scenario.HostID,
						Dates: []schedule_gateway.CreateDate{
							{
								MonthYear:   time.Date(2023, 0, 0, 0, 0, 0, 0, time.UTC),
								DaysOfMonth: fixtures.Preference.AllDaysOfMonth,
							},
						},
					},
					scenario.TutorToken,
				)

				return scenario, statusCode
			},
			expected: func(scenario fixtures.CreateDefaultOutput) schedule_gateway.PaginateData {
				return schedule_gateway.PaginateData{}
			},
			statusCode: http.StatusNotFound,
		},

		{
			title: "should not schedule if pet is non-neutered male and host only accepts neutered pets",
			seed: func() (fixtures.CreateDefaultOutput, int) {
				ACCEPT_ONLY_NEUTERED_MALES := True()
				scenario := fixtures.Preference.CreateDefault(t, &preference_case.CreateInput{
					OnlyVaccinated:          True(),
					AcceptElderly:           True(),
					AcceptOnlyNeuteredMales: ACCEPT_ONLY_NEUTERED_MALES,
					AcceptFemales:           True(),
					DaysOfMonth:             fixtures.Preference.AllDaysOfMonth,
					AcceptFemaleInHeat:      True(),
					AcceptPuppies:           True(),
					AcceptMales:             True(),
					PetWeight:               fixtures.Preference.AllPetWeight,
				})

				newPet := pet_case.PatchValues{
					Neutered: False(),
					Gender:   pet.Male,
				}
				response, status := fixtures.Pet.Patch(t, scenario.PetID, newPet, scenario.TutorToken)
				require.Equal(t, "OK", response)
				require.Equal(t, http.StatusOK, status)

				_, statusCode := fixtures.Schedule.Create(
					t,
					schedule_case.CreateInput{
						PetIDs: []string{scenario.PetID},
						HostID: scenario.HostID,
						Dates: []schedule_gateway.CreateDate{
							{
								MonthYear:   time.Date(2023, 0, 0, 0, 0, 0, 0, time.UTC),
								DaysOfMonth: fixtures.Preference.AllDaysOfMonth,
							},
						},
					},
					scenario.TutorToken,
				)

				return scenario, statusCode
			},
			expected: func(scenario fixtures.CreateDefaultOutput) schedule_gateway.PaginateData {
				return schedule_gateway.PaginateData{}
			},
			statusCode: http.StatusNotFound,
		},

		{
			title: "should not schedule if pet is male and host does not accept males even if neutered",
			seed: func() (fixtures.CreateDefaultOutput, int) {
				ACCEPT_MALES := False()
				NEUTERED := True()
				scenario := fixtures.Preference.CreateDefault(t, &preference_case.CreateInput{
					OnlyVaccinated:          True(),
					AcceptElderly:           True(),
					AcceptOnlyNeuteredMales: True(),
					AcceptFemales:           True(),
					DaysOfMonth:             fixtures.Preference.AllDaysOfMonth,
					AcceptFemaleInHeat:      True(),
					AcceptPuppies:           True(),
					AcceptMales:             ACCEPT_MALES,
					PetWeight:               fixtures.Preference.AllPetWeight,
				})

				newPet := pet_case.PatchValues{
					Neutered: NEUTERED,
					Gender:   pet.Male,
				}
				response, status := fixtures.Pet.Patch(t, scenario.PetID, newPet, scenario.TutorToken)
				require.Equal(t, "OK", response)
				require.Equal(t, http.StatusOK, status)

				_, statusCode := fixtures.Schedule.Create(
					t,
					schedule_case.CreateInput{
						PetIDs: []string{scenario.PetID},
						HostID: scenario.HostID,
						Dates: []schedule_gateway.CreateDate{
							{
								MonthYear:   time.Date(2023, 0, 0, 0, 0, 0, 0, time.UTC),
								DaysOfMonth: fixtures.Preference.AllDaysOfMonth,
							},
						},
					},
					scenario.TutorToken,
				)

				return scenario, statusCode
			},
			expected: func(scenario fixtures.CreateDefaultOutput) schedule_gateway.PaginateData {
				return schedule_gateway.PaginateData{}
			},
			statusCode: http.StatusNotFound,
		},

		{
			title: "should not schedule if pet is female and host does not accept females even if neutered",
			seed: func() (fixtures.CreateDefaultOutput, int) {
				ACCEPT_FEMALES := False()
				NEUTERED := True()
				scenario := fixtures.Preference.CreateDefault(t, &preference_case.CreateInput{
					OnlyVaccinated:          True(),
					AcceptElderly:           True(),
					AcceptOnlyNeuteredMales: True(),
					AcceptFemales:           ACCEPT_FEMALES,
					DaysOfMonth:             fixtures.Preference.AllDaysOfMonth,
					AcceptFemaleInHeat:      True(),
					AcceptPuppies:           True(),
					AcceptMales:             True(),
					PetWeight:               fixtures.Preference.AllPetWeight,
				})

				newPet := pet_case.PatchValues{
					Neutered: NEUTERED,
					Gender:   pet.Female,
				}
				response, status := fixtures.Pet.Patch(t, scenario.PetID, newPet, scenario.TutorToken)
				require.Equal(t, "OK", response)
				require.Equal(t, http.StatusOK, status)

				_, statusCode := fixtures.Schedule.Create(
					t,
					schedule_case.CreateInput{
						PetIDs: []string{scenario.PetID},
						HostID: scenario.HostID,
						Dates: []schedule_gateway.CreateDate{
							{
								MonthYear:   time.Date(2023, 0, 0, 0, 0, 0, 0, time.UTC),
								DaysOfMonth: fixtures.Preference.AllDaysOfMonth,
							},
						},
					},
					scenario.TutorToken,
				)

				return scenario, statusCode
			},
			expected: func(scenario fixtures.CreateDefaultOutput) schedule_gateway.PaginateData {
				return schedule_gateway.PaginateData{}
			},
			statusCode: http.StatusNotFound,
		},

		{
			title: "should not schedule if host does not accept female in heat and it is in heat",
			seed: func() (fixtures.CreateDefaultOutput, int) {
				ACCEPT_FEMALES := True()
				ACCEPT_FEMALES_IN_HEAT := False()
				NEUTERED := False()

				scenario := fixtures.Preference.CreateDefault(t, &preference_case.CreateInput{
					OnlyVaccinated:          True(),
					AcceptElderly:           True(),
					AcceptOnlyNeuteredMales: True(),
					AcceptFemales:           ACCEPT_FEMALES,
					DaysOfMonth:             fixtures.Preference.AllDaysOfMonth,
					AcceptFemaleInHeat:      ACCEPT_FEMALES_IN_HEAT,
					AcceptPuppies:           True(),
					AcceptMales:             True(),
					PetWeight:               fixtures.Preference.AllPetWeight,
				})

				newPet := pet_case.PatchValues{
					Gender:   pet.Female,
					Neutered: NEUTERED,
				}
				response, status := fixtures.Pet.Patch(t, scenario.PetID, newPet, scenario.TutorToken)
				require.Equal(t, "OK", response)
				require.Equal(t, http.StatusOK, status)

				FEMALES_IN_HEAT := map[string]bool{scenario.PetID: true}
				_, statusCode := fixtures.Schedule.Create(
					t,
					schedule_case.CreateInput{
						PetIDs:        []string{scenario.PetID},
						HostID:        scenario.HostID,
						FemalesInHeat: FEMALES_IN_HEAT,
						Dates: []schedule_gateway.CreateDate{
							{
								MonthYear:   time.Date(2023, 0, 0, 0, 0, 0, 0, time.UTC),
								DaysOfMonth: fixtures.Preference.AllDaysOfMonth,
							},
						},
					},
					scenario.TutorToken,
				)

				return scenario, statusCode
			},
			expected: func(scenario fixtures.CreateDefaultOutput) schedule_gateway.PaginateData {
				return schedule_gateway.PaginateData{}
			},
			statusCode: http.StatusNotFound,
		},

		{
			title: "should not schedule if pet is not vaccinated and host only accepts vaccinated pets",
			seed: func() (fixtures.CreateDefaultOutput, int) {
				ONLY_VACCINATED := True()
				VACCINATED := False()
				scenario := fixtures.Preference.CreateDefault(t, &preference_case.CreateInput{
					OnlyVaccinated:          ONLY_VACCINATED,
					AcceptElderly:           True(),
					AcceptOnlyNeuteredMales: True(),
					AcceptFemales:           True(),
					DaysOfMonth:             fixtures.Preference.AllDaysOfMonth,
					AcceptFemaleInHeat:      True(),
					AcceptPuppies:           True(),
					AcceptMales:             True(),
					PetWeight:               fixtures.Preference.AllPetWeight,
				})

				newPet := pet_case.PatchValues{
					Vaccinated: VACCINATED,
				}

				response, status := fixtures.Pet.Patch(t, scenario.PetID, newPet, scenario.TutorToken)
				require.Equal(t, "OK", response)
				require.Equal(t, http.StatusOK, status)

				_, statusCode := fixtures.Schedule.Create(
					t,
					schedule_case.CreateInput{
						PetIDs: []string{scenario.PetID},
						HostID: scenario.HostID,
						Dates: []schedule_gateway.CreateDate{
							{
								MonthYear:   time.Date(2023, 0, 0, 0, 0, 0, 0, time.UTC),
								DaysOfMonth: fixtures.Preference.AllDaysOfMonth,
							},
						},
					},
					scenario.TutorToken,
				)

				return scenario, statusCode
			},
			expected: func(scenario fixtures.CreateDefaultOutput) schedule_gateway.PaginateData {
				return schedule_gateway.PaginateData{}
			},
			statusCode: http.StatusNotFound,
		},

		{
			title: "should not schedule if pet is elderly and host does not accept elderly pets",
			seed: func() (fixtures.CreateDefaultOutput, int) {
				ACCEPT_ELDERLY := False()
				scenario := fixtures.Preference.CreateDefault(t, &preference_case.CreateInput{
					OnlyVaccinated:          True(),
					AcceptElderly:           ACCEPT_ELDERLY,
					AcceptOnlyNeuteredMales: True(),
					AcceptFemales:           True(),
					DaysOfMonth:             fixtures.Preference.AllDaysOfMonth,
					AcceptFemaleInHeat:      True(),
					AcceptPuppies:           True(),
					AcceptMales:             True(),
					PetWeight:               fixtures.Preference.AllPetWeight,
				})

				PET_BIRTHDATE := time.Now().AddDate(-pet.ElderlyAge, 0, 0)
				newPet := pet_case.PatchValues{
					Birthdate: PET_BIRTHDATE,
				}

				response, status := fixtures.Pet.Patch(t, scenario.PetID, newPet, scenario.TutorToken)
				require.Equal(t, "OK", response)
				require.Equal(t, http.StatusOK, status)

				_, statusCode := fixtures.Schedule.Create(
					t,
					schedule_case.CreateInput{
						PetIDs: []string{scenario.PetID},
						HostID: scenario.HostID,
						Dates: []schedule_gateway.CreateDate{
							{
								MonthYear:   time.Date(2023, 0, 0, 0, 0, 0, 0, time.UTC),
								DaysOfMonth: fixtures.Preference.AllDaysOfMonth,
							},
						},
					},
					scenario.TutorToken,
				)

				return scenario, statusCode
			},
			expected: func(scenario fixtures.CreateDefaultOutput) schedule_gateway.PaginateData {
				return schedule_gateway.PaginateData{}
			},
			statusCode: http.StatusNotFound,
		},

		{
			title: "should not schedule if pet is a puppy and host does not accept puppies",
			seed: func() (fixtures.CreateDefaultOutput, int) {
				ACCEPT_PUPPY := False()
				scenario := fixtures.Preference.CreateDefault(t, &preference_case.CreateInput{
					OnlyVaccinated:          True(),
					AcceptElderly:           True(),
					AcceptOnlyNeuteredMales: True(),
					AcceptFemales:           True(),
					DaysOfMonth:             fixtures.Preference.AllDaysOfMonth,
					AcceptFemaleInHeat:      True(),
					AcceptPuppies:           ACCEPT_PUPPY,
					AcceptMales:             True(),
					PetWeight:               fixtures.Preference.AllPetWeight,
				})

				PET_BIRTHDATE := time.Now().AddDate(-pet.PuppieAge, 0, 1)
				newPet := pet_case.PatchValues{
					Birthdate: PET_BIRTHDATE,
				}

				response, status := fixtures.Pet.Patch(t, scenario.PetID, newPet, scenario.TutorToken)
				require.Equal(t, "OK", response)
				require.Equal(t, http.StatusOK, status)

				_, statusCode := fixtures.Schedule.Create(
					t,
					schedule_case.CreateInput{
						PetIDs: []string{scenario.PetID},
						HostID: scenario.HostID,
						Dates: []schedule_gateway.CreateDate{
							{
								MonthYear:   time.Date(2023, 0, 0, 0, 0, 0, 0, time.UTC),
								DaysOfMonth: fixtures.Preference.AllDaysOfMonth,
							},
						},
					},
					scenario.TutorToken,
				)

				return scenario, statusCode
			},
			expected: func(scenario fixtures.CreateDefaultOutput) schedule_gateway.PaginateData {
				return schedule_gateway.PaginateData{}
			},
			statusCode: http.StatusNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			fixtures.CleanDatabase()
			scenario, statusCode := test.seed()
			require.Equal(t, test.statusCode, statusCode)

			response := schedule_gateway.PaginateOutput{
				Data: []schedule_gateway.PaginateData{},
			}

			fixtures.Get(t, fixtures.GetInput{
				Query: schedule_case.PaginateFilter{
					Status: schedule_status.Open,
				},
				URI:      fixtures.Schedule.URI,
				Response: &response,
				Token:    scenario.HostToken,
			})

			expected := test.expected(scenario)
			if len(response.Data) == 0 {
				require.Equal(t, expected, schedule_gateway.PaginateData{})
			} else {
				require.Equal(t, expected, response.Data[0])
			}
		})
	}
}
