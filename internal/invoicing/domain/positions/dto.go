package positions

// EuroCents represents just some money.
type EuroCents int

// A Position on an invoice.
type Position struct {
	No    int
	Title string
	Value EuroCents
}
