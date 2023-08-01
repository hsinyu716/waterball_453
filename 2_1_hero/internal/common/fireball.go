package common

type FireBall struct {
}

func (w *FireBall) Attack(attacker *Hero, attacked *Hero) {
	for i := 0; i < 3; i++ {
		attackHp := 50
		attacked.Damage(attackHp)
	}
}
