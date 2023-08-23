package common

type Card interface {
	Translate() string
	GenerateDeck() []Card
}
