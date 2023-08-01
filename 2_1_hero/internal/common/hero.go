package common

import "fmt"

type Hero struct {
	name       string
	hp         int
	attackType AttackStrategy
}

func NewHero(name string, hp int, attack AttackStrategy) *Hero {
	return &Hero{
		name:       name,
		hp:         hp,
		attackType: attack,
	}
}

type IHero interface {
	GetName() string
	SetName(name string)
	GetHp() int
	SetHp(hp int)
	Damage(damage int)
	IsDead() bool
	SetAttackType(attack AttackStrategy)
	Attack(attacked *Hero)
}

func (h *Hero) GetName() string {
	return h.name
}

func (h *Hero) SetName(name string) {
	h.name = name
}

func (h *Hero) GetHp() int {
	return h.hp
}

func (h *Hero) SetHp(hp int) {
	if hp < 0 {
		hp = 0
	}
	h.hp = hp
}

func (h *Hero) Damage(damage int) {
	h.SetHp(h.GetHp() - damage)
	fmt.Println(fmt.Sprintf("%s 受到攻擊 -%v HP, 剩 %v HP", h.name, damage, h.GetHp()))
}

func (h *Hero) IsDead() bool {
	return h.hp <= 0
}

func (h *Hero) SetAttackType(attack AttackStrategy) {
	h.attackType = attack
}

func (h *Hero) Attack(attacked *Hero) {
	h.attackType.Attack(h, attacked)
}
