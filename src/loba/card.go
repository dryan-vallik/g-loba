package loba

type Suite uint8

const (
	SuiteNull Suite = iota
	Clubs
	Hearts
	Spades
	Diamonds
)

type Rank uint8

const (
	RankNull Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	Suite Suite
	Rank Rank
	IsJoker bool
}
