package engine

import "loba/src/loba"

type CardPredicts struct {
	Card loba.Card

	// Contains all the probabilities of each player having the card
	Probabilities []float32
	Owner         CardOwner
}

type DataBank struct {
	CardPredictions map[loba.Card]CardPredicts
	CardsOnDeck     uint16
	CardsOnDiscard  uint16
	JokerCount      uint16
	Stats           []PlayerStats
}
