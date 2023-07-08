package receiver

type Gender int

func (g Gender) Valid() bool {
	return g >= M && g <= maxGender
}

const (
	Undefined Gender = iota
	M
	W
	D
	maxGender = D
)

type Name string

func (n Name) Valid() bool {
	return n != "" && len(n) < 100
}

type Zip string

func (z Zip) Valid() bool {
	return len(z) > 2 && len(z) < 10
}

type Title int

func (t Title) Valid() bool {
	return t >= None && t <= maxTitle
}

const (
	None Title = iota
	Dr
	Prof
	maxTitle = Prof
)

type ContactPerson struct {
	Title     Title
	Firstname Name
	Lastname  Name
	Gender    Gender
}

type Address struct {
	Street    Name
	City      Name
	Zip       Zip
	Housename Name
	Floor     Name
}

type Company struct {
	Name    Name
	Address Address
}
