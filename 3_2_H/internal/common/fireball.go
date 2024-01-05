package common

import "fmt"

type FireBall struct {
	Subscriber
}

func NewFireBall() *FireBall {
	f := &FireBall{}
	f.SetName("火球")
	return f
}

func (f *FireBall) Notify(video *Video) {
	if video.length <= 60 {
		fmt.Println(fmt.Sprintf("%s 解除訂閱了 %s", f.name, video.channel.name))
		f.UnSubscribe(video.channel)
	}
}

func (f *FireBall) Subscribe(channel *Channel) {
	channel.Subscribe(f)
}

func (f *FireBall) UnSubscribe(channel *Channel) {
	channel.UnSubscribe(f)
}
