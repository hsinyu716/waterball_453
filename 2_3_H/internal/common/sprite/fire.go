package sprite

type Fire struct {
	Sprite
}

func NewFire(p int) *Fire {
	f := &Fire{
		Sprite: Sprite{
			sign: "F",
		},
	}
	f.setPosition(p)
	return f
}
