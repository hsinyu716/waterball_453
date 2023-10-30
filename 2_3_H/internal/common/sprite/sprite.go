package sprite

import (
	"fmt"
	"reflect"
)

type Sprite struct {
	sign     string
	position int
}

type ISprite interface {
	setPosition(p int)
	GetPosition() int
	ChangePosition(p int)
	Remove(spritePositions *[]ISprite)
	Move(toSprite ISprite, spritePositions *[]ISprite)
}

func (s *Sprite) Remove(spritePositions *[]ISprite) {
	fmt.Println("s.position", s.position)
	(*spritePositions)[s.position] = nil
}

func (s *Sprite) setPosition(p int) {
	s.position = p
}

func (s *Sprite) GetPosition() int {
	return s.position
}

func (s *Sprite) ChangePosition(p int) {
	s.position = p
}

func (s *Sprite) Move(toSprite ISprite, spritePositions *[]ISprite) {
	position := toSprite.GetPosition()
	(*spritePositions)[position] = s
	s.Remove(spritePositions)
	(*spritePositions)[position].ChangePosition(position)
}

type TypeSprite string

const (
	HeroSprite  TypeSprite = "*sprite.Hero"
	WaterSprite TypeSprite = "*sprite.Water"
	FireSprite  TypeSprite = "*sprite.Fire"
	IceSprite   TypeSprite = "*sprite.Ice"
	NilSprite   TypeSprite = "*sprite.Nil"
)

var (
	ConflictType = []string{
		string(FireSprite),
		string(WaterSprite),
		string(IceSprite),
	}

	WeakenType = []string{
		string(FireSprite),
	}

	StrengthenType = []string{
		string(WaterSprite),
		string(IceSprite),
	}
)

func StrengthenOrWeaken(fromSprite, toSprite ISprite, spriteType []string) bool {
	match := false
	for _, hpType := range spriteType {
		match = reflect.TypeOf(fromSprite).String() == string(HeroSprite) && reflect.TypeOf(toSprite).String() == hpType ||
			reflect.TypeOf(fromSprite).String() == hpType && reflect.TypeOf(toSprite).String() == string(HeroSprite)
		if match {
			break
		}
	}
	return match
}
