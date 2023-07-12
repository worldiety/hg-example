package helloworld

import "fmt"

type Person struct {
	Firstname string
	Age       int
}

type PersonValidationError interface {
	error
	PersonValidationError() bool
}

type UnexpectedNameError string

func (u UnexpectedNameError) Error() string {
	return fmt.Sprintf("unexpected name: %s", string(u))
}

func (u UnexpectedNameError) PersonValidationError() bool { return true }

type InvalidAgeError int

func (i InvalidAgeError) Error() string {
	return fmt.Sprintf("invalid age: %s", int(i))
}

func (i InvalidAgeError) PersonValidationError() bool { return true }

func PersonCheck(person Person) PersonValidationError {
	if person.Age < 18 || person.Age > 120 {
		return InvalidAgeError(person.Age)
	}

	if person.Firstname != "Torben" {
		return UnexpectedNameError(person.Firstname)
	}

	return nil
}
