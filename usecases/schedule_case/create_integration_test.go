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
		title    string
		seed     func() fixtures.CreateDefaultOutput
		expected func(scenario fixtures.CreateDefaultOutput) schedule_gateway.PaginateData
	}

	tests := []testCase{
		{
			title: "happy path",
			seed: func() fixtures.CreateDefaultOutput {
				scenario := fixtures.Preference.CreateDefault(t, nil)
				fixtures.Schedule.Create(
					t,
					schedule_case.CreateInput{
						PetID:        scenario.PetID,
						HostID:       scenario.HostID,
						MonthYear:    time.Date(2023, 0, 0, 0, 0, 0, 0, time.UTC),
						DaysOfMonth:  fixtures.Preference.AllDaysOfMonth,
						Notes:        "Notes",
						FemaleInHeat: nil,
					},
					scenario.TutorToken,
				)

				return scenario
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
		},

		{
			title: "should not schedule if tutor is not the owner of the pet",
			seed: func() fixtures.CreateDefaultOutput {
				scenario := fixtures.Preference.CreateDefault(t, nil)
				NOT_TUTOR_PET := fixtures.Pet.Create(t, nil, scenario.HostToken)
				fixtures.Schedule.Create(
					t,
					schedule_case.CreateInput{
						PetID:        NOT_TUTOR_PET,
						HostID:       scenario.HostID,
						MonthYear:    time.Date(2023, 0, 0, 0, 0, 0, 0, time.UTC),
						DaysOfMonth:  fixtures.Preference.AllDaysOfMonth,
						Notes:        "Notes",
						FemaleInHeat: nil,
					},
					scenario.TutorToken,
				)

				return scenario
			},
			expected: func(scenario fixtures.CreateDefaultOutput) schedule_gateway.PaginateData {
				return schedule_gateway.PaginateData{}
			},
		},

		{
			title: "should not schedule if availability not meet tutor needs",
			seed: func() fixtures.CreateDefaultOutput {
				var from1To5 uint32 = 0b11111
				scenario := fixtures.Preference.CreateDefault(t, &preference_case.CreateInput{
					OnlyVaccinated:          True(),
					AcceptElderly:           True(),
					AcceptOnlyNeuteredMales: False(),
					AcceptFemales:           True(),
					DaysOfMonth:             from1To5,
					AcceptFemaleInHeat:      True(),
					AcceptPuppies:           True(),
					AcceptMales:             True(),
					PetWeight:               fixtures.Preference.AllPetWeight,
				})

				var from1To3 uint32 = 0b111
				fixtures.Schedule.Create(
					t,
					schedule_case.CreateInput{
						PetID:        scenario.PetID,
						HostID:       scenario.HostID,
						MonthYear:    time.Date(2023, 0, 0, 0, 0, 0, 0, time.UTC),
						DaysOfMonth:  from1To3,
						FemaleInHeat: nil,
					},
					scenario.TutorToken,
				)

				return scenario
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
		},

		{
			title: "should schedule if availability is greater than tutor needs",
			seed: func() fixtures.CreateDefaultOutput {
				var from1To5 uint32 = 0b11111
				scenario := fixtures.Preference.CreateDefault(t, &preference_case.CreateInput{
					OnlyVaccinated:          True(),
					AcceptElderly:           True(),
					AcceptOnlyNeuteredMales: False(),
					AcceptFemales:           True(),
					DaysOfMonth:             from1To5,
					AcceptFemaleInHeat:      True(),
					AcceptPuppies:           True(),
					AcceptMales:             True(),
					PetWeight:               fixtures.Preference.AllPetWeight,
				})

				var from1To3 uint32 = 0b111
				fixtures.Schedule.Create(
					t,
					schedule_case.CreateInput{
						PetID:        scenario.PetID,
						HostID:       scenario.HostID,
						MonthYear:    time.Date(2023, 0, 0, 0, 0, 0, 0, time.UTC),
						DaysOfMonth:  from1To3,
						FemaleInHeat: nil,
					},
					scenario.TutorToken,
				)

				return scenario
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
		},

		{
			title: "should not schedule if pet is large and host only accepts small pets",
			seed: func() fixtures.CreateDefaultOutput {
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

				fixtures.Schedule.Create(
					t,
					schedule_case.CreateInput{
						PetID:        scenario.PetID,
						HostID:       scenario.HostID,
						MonthYear:    time.Date(2023, 0, 0, 0, 0, 0, 0, time.UTC),
						DaysOfMonth:  fixtures.Preference.AllDaysOfMonth,
						FemaleInHeat: nil,
					},
					scenario.TutorToken,
				)

				return scenario
			},
			expected: func(scenario fixtures.CreateDefaultOutput) schedule_gateway.PaginateData {
				return schedule_gateway.PaginateData{}
			},
		},

		{
			title: "should not schedule if pet is small and host only accepts large pets",
			seed: func() fixtures.CreateDefaultOutput {
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

				fixtures.Schedule.Create(
					t,
					schedule_case.CreateInput{
						PetID:        scenario.PetID,
						HostID:       scenario.HostID,
						MonthYear:    time.Date(2023, 0, 0, 0, 0, 0, 0, time.UTC),
						DaysOfMonth:  fixtures.Preference.AllDaysOfMonth,
						FemaleInHeat: nil,
					},
					scenario.TutorToken,
				)

				return scenario
			},
			expected: func(scenario fixtures.CreateDefaultOutput) schedule_gateway.PaginateData {
				return schedule_gateway.PaginateData{}
			},
		},

		{
			title: "should not schedule if pet is non-neutered male and host only accepts neutered pets",
			seed: func() fixtures.CreateDefaultOutput {
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

				fixtures.Schedule.Create(
					t,
					schedule_case.CreateInput{
						PetID:        scenario.PetID,
						HostID:       scenario.HostID,
						MonthYear:    time.Date(2023, 0, 1, 0, 0, 0, 0, time.UTC),
						DaysOfMonth:  fixtures.Preference.AllDaysOfMonth,
						FemaleInHeat: nil,
					},
					scenario.TutorToken,
				)

				return scenario
			},
			expected: func(scenario fixtures.CreateDefaultOutput) schedule_gateway.PaginateData {
				return schedule_gateway.PaginateData{}
			},
		},

		{
			title: "should not schedule if pet is male and host does not accept males even if neutered",
			seed: func() fixtures.CreateDefaultOutput {
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

				fixtures.Schedule.Create(
					t,
					schedule_case.CreateInput{
						PetID:        scenario.PetID,
						HostID:       scenario.HostID,
						MonthYear:    time.Date(2023, 0, 1, 0, 0, 0, 0, time.UTC),
						DaysOfMonth:  fixtures.Preference.AllDaysOfMonth,
						FemaleInHeat: nil,
					},
					scenario.TutorToken,
				)

				return scenario
			},
			expected: func(scenario fixtures.CreateDefaultOutput) schedule_gateway.PaginateData {
				return schedule_gateway.PaginateData{}
			},
		},

		{
			title: "should not schedule if pet is female and host does not accept females even if neutered",
			seed: func() fixtures.CreateDefaultOutput {
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

				fixtures.Schedule.Create(
					t,
					schedule_case.CreateInput{
						PetID:        scenario.PetID,
						HostID:       scenario.HostID,
						MonthYear:    time.Date(2023, 0, 1, 0, 0, 0, 0, time.UTC),
						DaysOfMonth:  fixtures.Preference.AllDaysOfMonth,
						FemaleInHeat: nil,
					},
					scenario.TutorToken,
				)

				return scenario
			},
			expected: func(scenario fixtures.CreateDefaultOutput) schedule_gateway.PaginateData {
				return schedule_gateway.PaginateData{}
			},
		},

		{
			title: "should not schedule if pet is not vaccinated and host only accepts vaccinated pets",
			seed: func() fixtures.CreateDefaultOutput {
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

				fixtures.Schedule.Create(
					t,
					schedule_case.CreateInput{
						PetID:        scenario.PetID,
						HostID:       scenario.HostID,
						MonthYear:    time.Date(2023, 0, 1, 0, 0, 0, 0, time.UTC),
						DaysOfMonth:  fixtures.Preference.AllDaysOfMonth,
						FemaleInHeat: nil,
					},
					scenario.TutorToken,
				)

				return scenario
			},
			expected: func(scenario fixtures.CreateDefaultOutput) schedule_gateway.PaginateData {
				return schedule_gateway.PaginateData{}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			fixtures.CleanDatabase()
			scenario := test.seed()

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
