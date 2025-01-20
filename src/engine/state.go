package engine

import "math"

// A CardOwner represents which player owns a card. Negative values are reserved for places that are not players, like the deck (OwnerDeck: -125), the discard pile (OwnerDiscard: -126), or owned by ourselves (OwnerSelf: -127).
// OwnerUnknown is the default, and is equal to 0
// Numbers 1 and above are free and can be used to specify which player owns the card.
type CardOwner = int8

const (
	OwnerSelf CardOwner = iota + math.MinInt8
	OwnerDiscard
	OwnerDeck
)

const (
	OwnerUnknown CardOwner = 0 // Default
)
