package utils

import "fmt"

const (
	DataNil            = "沒有元素可移動"
	DataMoveToNil      = "元素移動"
	DataSameType       = "同類不移動"
	DataHeroStrengthen = "Hero 遇水/冰強化10滴血, Water/Ice消失"
	DataHeroWeaken     = "Hero 遇火弱化10滴血, Fire消失"
	DataHeroDead       = "Hero 血量為0, 死亡"
	DataConflict       = "冰水火不容"
)

func MsgPrint(v string) {
	fmt.Println(v)
}
