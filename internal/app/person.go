package app

type Person struct {
	Id        string `db:"id"`
	Email     string `db:"email"`
	Phone     string `db:"phone"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}

func (p *Person) GetId() string {
	return p.Id
}

func (p *Person) GetEmail() string {
	return p.Email
}

func (p *Person) GetPhone() string {
	return p.Phone
}

func (p *Person) GetFirstName() string {
	return p.FirstName
}

func (p *Person) GetLastName() string {
	return p.LastName
}
