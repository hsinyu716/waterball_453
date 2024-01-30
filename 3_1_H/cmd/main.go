package main

import (
	"cosmos.youtube/internal/common"
)

func main() {
	fireBall := common.ISubscriberObserver(common.NewFireBall())
	waterBall := common.ISubscriberObserver(common.NewWaterBall())

	pewDiePie := new(common.Channel)
	pewDiePie.SetName("PewDiePie")
	waterSoftware := new(common.Channel)
	waterSoftware.SetName("水球軟體學院")
	waterBall.Subscribe(waterSoftware)
	waterBall.Subscribe(pewDiePie)
	fireBall.Subscribe(waterSoftware)
	fireBall.Subscribe(pewDiePie)

	video := common.NewVideo("C1M1S2", "這個世界正是物件導向的呢！", 240)
	waterSoftware.Upload(video)
	video = common.NewVideo("Hello guys", "Clickbait", 30)
	pewDiePie.Upload(video)
	video = common.NewVideo("C1M1S3", "物件 vs. 類別", 60)
	waterSoftware.Upload(video)
	video = common.NewVideo("Minecraft", "Let’s play Minecraft", 1800)
	pewDiePie.Upload(video)
}
