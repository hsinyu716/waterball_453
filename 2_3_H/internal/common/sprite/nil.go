package sprite

type Nil struct {
	Sprite
}

func NewNil(p int) *Nil {
	n := &Nil{
		Sprite: Sprite{
			sign: "N",
		},
	}
	n.setPosition(p)
	return n
}
