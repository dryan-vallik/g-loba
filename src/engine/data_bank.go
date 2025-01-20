package engine

import "loba/src/loba"

type CardSlot struct {
	Card loba.Card

	// Contains all the probabilities of each player having the card
	Probabilities []float32
	Owner         CardOwner
}

type DataBank struct {
	Cards map[loba.Card]CardSlot
}
