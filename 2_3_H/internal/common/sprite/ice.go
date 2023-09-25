package sprite

type Ice struct {
	Sprite
}

func NewIce(p int) *Ice {
	f := &Ice{
		Sprite: Sprite{
			sign: "I",
		},
	}
	f.setPosition(p)
	return f
}
