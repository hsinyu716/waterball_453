package common

type WaterBall struct {
}

func (w *WaterBall) Attack(attacker *Hero, attacked *Hero) {
	attackHp := attacker.GetHp() * 1 / 2
	attacked.Damage(attackHp)
}
