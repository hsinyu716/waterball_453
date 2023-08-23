package card

type Card interface {
	GenerateDeck() []Card
	Translate() string
}
