package pet_gateway

type PetGateway interface {
	Create(input CreateInput) (string, error)
	Patch(filter PatchFilter, values PatchValues) (bool, error)
	List(input ListInput) ([]ListOutput, error)
	Delete(id string) (bool, error)
	GetByID(id string) (*GetByIDOutput, error)
}

type CreateInput struct {
	UserID    string
	Species   string
	Name      string
	Breed     string
	Birthdate string
	Gender    string
	Weight    uint8
	Neutered  bool
}

type PatchFilter struct {
	ID        *string
	Species   *string
	Name      *string
	Breed     *string
	Birthdate *string
	Gender    *string
	Weight    *uint8
}

type PatchValues struct {
	Species   string
	Name      string
	Breed     string
	Birthdate string
	Gender    string
	Weight    uint8
}

type ListInput struct {
	TutorID *string
}

type ListOutput struct {
	Name      string
	Breed     string
	Birthdate string
	Gender    string
	Weight    uint8
	Species   string
}

type GetByIDOutput struct {
	Name      string
	Breed     string
	Birthdate string
	Gender    string
	Weight    uint8
	Species   string
	Neutered  bool
}
