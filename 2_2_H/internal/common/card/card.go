package card

type Card interface {
	InitDeck() []Card
	Translate() string
}
