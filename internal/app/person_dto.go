package app

type PersonDTO struct {
	Email     string `json:"Email"`
	Phone     string `json:"Phone"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
}

func FromDTO(person *PersonDTO) Person {
	return Person{
		Email:     person.Email,
		Phone:     person.Phone,
		FirstName: person.FirstName,
		LastName:  person.LastName,
	}
}
