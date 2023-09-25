package sprite

type Water struct {
	Sprite
}

func NewWater(p int) *Water {
	w := &Water{
		Sprite: Sprite{
			sign: "W",
		},
	}
	w.setPosition(p)
	return w
}
