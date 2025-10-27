package author

type Author struct {
	FirstName string
	LastName string
}

func New(firstName, lastName string) (*Author, error) {
	return &Author{FirstName: firstName, LastName: lastName}, nil
}

func (a *Author) FullName() string {
	return a.FirstName + " " + a.LastName
}
