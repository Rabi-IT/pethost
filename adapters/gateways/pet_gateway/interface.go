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
	Size      string
	Birthdate string
	Gender    string
	Weight    string
}

type PatchFilter struct {
	ID        *string
	Species   *string
	Name      *string
	Breed     *string
	Size      *string
	Birthdate *string
	Gender    *string
	Weight    *string
}

type PatchValues struct {
	Species   string
	Name      string
	Breed     string
	Size      string
	Birthdate string
	Gender    string
	Weight    string
}

type ListInput struct {
	Name      *string
	Breed     *string
	Size      *string
	Birthdate *string
	Gender    *string
	Weight    *string
	Species   *string
}

type ListOutput struct {
	Name      string
	Breed     string
	Size      string
	Birthdate string
	Gender    string
	Weight    string
	Species   string
}

type GetByIDOutput struct {
	Name      string
	Breed     string
	Size      string
	Birthdate string
	Gender    string
	Weight    string
	Species   string
}
