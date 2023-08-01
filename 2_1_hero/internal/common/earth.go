package common

type Earth struct {
}

func (w *Earth) Attack(attacker *Hero, attacked *Hero) {
	for i := 0; i < 3; i++ {
		attackHp := 440
		attacked.Damage(attackHp)
	}
}
