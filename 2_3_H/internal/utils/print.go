package utils

import "fmt"

const (
	DataNil       = "沒有元素"
	DataSameType  = "同類不移動"
	DataHeroWater = "Hero 遇水補血10滴, Water消失"
	DataHeroFire  = "Hero 遇火失血10滴, Fire消失"
	DataHeroDead  = "Hero 死亡"
	DataWaterFire = "水火不容"
)

func MsgPrint(v string) {
	fmt.Println(v)
}
