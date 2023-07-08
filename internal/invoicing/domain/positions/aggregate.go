package positions

import "golang.org/x/exp/slices"

type CollectionID string

func (id CollectionID) Undefined() bool {
	return id == ""
}

// Positions is the aggregate root, for our positions.
type Positions struct {
	id        CollectionID
	positions []Position
}

func NewPositions(id CollectionID, pos []Position) *Positions {
	return &Positions{
		id:        id,
		positions: pos,
	}
}

func (p *Positions) All() []Position {
	return slices.Clone(p.positions) // return defensive copy
}

func (p *Positions) Add(pos Position) {
	p.positions = append(p.positions, pos)
	p.enumerate()
}

// ensure pre and postvariante
func (p *Positions) enumerate() {
	for i, position := range p.positions {
		position.No = i
	}
}
