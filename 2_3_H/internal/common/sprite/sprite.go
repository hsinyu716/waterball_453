package sprite

type Sprite struct {
	sign     string
	position int
}

type ISprite interface {
	setPosition(p int)
	GetPosition() int
	Remove(spritePositions *[]ISprite)
}

func (s *Sprite) Remove(spritePositions *[]ISprite) {
	(*spritePositions)[s.position] = nil
}

func (s *Sprite) setPosition(p int) {
	s.position = p
}

func (s *Sprite) GetPosition() int {
	return s.position
}

type TypeSprite string

const (
	HeroSprite  TypeSprite = "*sprite.Hero"
	WaterSprite TypeSprite = "*sprite.Water"
	FireSprite  TypeSprite = "*sprite.Fire"
	IceSprite   TypeSprite = "*sprite.Ice"
)
