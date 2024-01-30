package common

import "fmt"

type WaterBall struct {
	Subscriber
}

func NewWaterBall() *WaterBall {
	w := &WaterBall{}
	w.SetName("水球")
	return w
}

func (w *WaterBall) Subscribe(channel *Channel) {
	channel.Subscribe(w)
}

func (w *WaterBall) UnSubscribe(channel *Channel) {
	channel.UnSubscribe(w)
}

func (w *WaterBall) Notify(video *Video) {
	if video.length >= 180 {
		fmt.Println(fmt.Sprintf("%s 對影片 `%s` 按讚", w.name, video.title))
		video.Like(w)
	}
}
