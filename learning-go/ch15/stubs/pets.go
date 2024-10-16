package stubs

type User struct {
	Name string
	Age  int
}

type Pet struct {
	Name   string
	Weight int
}

type Entities interface {
	GetUser(id string) (User, error)
	GetPets(userId string) ([]Pet, error)
}

type Logic struct {
	Entities Entities
}

func (l Logic) GetPetNames(userId string) ([]string, error) {
	pets, err := l.Entities.GetPets(userId)
	if err != nil {
		return nil, err
	}
	petNames := make([]string, len(pets))
	for i, p := range pets {
		petNames[i] = p.Name
	}
	return petNames, nil
}
