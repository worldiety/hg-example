package positions

import "fmt"

type LoadingError interface {
	error
	LoadingError() bool
	// NotFoundError | InfrastructureError
}

type SavingError interface {
	error
	SavingError() bool
	InvoiceDraft() bool
	// InfrastructureError
}

type NotFoundError CollectionID

func (e NotFoundError) LoadingError() bool {
	return true
}

func (e NotFoundError) ID() CollectionID {
	return CollectionID(e)
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("not found: %s", e.String())
}

func (e NotFoundError) String() string {
	return string(e)
}

type InfrastructureError struct {
	Cause error
}

func (e InfrastructureError) LoadingError() bool {
	return true
}

func (e InfrastructureError) SavingError() bool {
	return true
}

func (e InfrastructureError) Error() string {
	return fmt.Sprintf("infrastructure error: %v", e.Cause)
}

func (e InfrastructureError) InvoiceDraft() bool {
	return true
}

// Unwrap is usually unwanted to hide implementation details
//func (e InfrastructureError) Unwrap() error {
//	return e.Cause
//}

// Repository models logical transaction states (the aggregate root) for
// different (typically I/O based) implementations.
type Repository interface {
	Load(id CollectionID) (Positions, LoadingError)
	Save(positions Positions) SavingError
}

// Loader is an alternative but more functional approach for modelling a kind of repository.
// The advantage is, that dependency requirements are absolutely exact and testing requires no mocks.
// It allows easy currying and avoids structs for modelling behavior.
type Loader func(id CollectionID) (Positions, LoadingError)

// Saver is an alternative but more functional approach for modelling a kind of repository.
// The advantage is, that dependency requirements are perfectly exact and testing requires no mocks.
// It allows easy currying and avoids structs for modelling behavior.
type Saver func(positions Positions) SavingError
