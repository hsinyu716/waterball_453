package common

import "fmt"

type Channel struct {
	name       string
	videos     []*Video
	subscriber []ISubscriberObserver
}

type IChannel interface {
	Subscribe(subscriber ISubscriberObserver)
	UnSubscribe(subscriber ISubscriberObserver)
	Notify()
	Upload(video *Video)
}

func (c *Channel) SetName(name string) {
	c.name = name
}

func (c *Channel) Subscribe(subscriber ISubscriberObserver) {
	fmt.Println(fmt.Sprintf("%s 訂閱了 %s", subscriber.GetName(), c.name))
	c.subscriber = append(c.subscriber, subscriber)
}

func (c *Channel) UnSubscribe(subscriber ISubscriberObserver) {
	find := -1
	for i, s := range c.subscriber {
		if s == subscriber {
			find = i
		}
	}
	if find > -1 {
		c.subscriber = append(c.subscriber[:find], c.subscriber[find+1:]...)
	}
}

func (c *Channel) Notify() {
	for _, s := range c.subscriber {
		s.Notify(c.videos[len(c.videos)-1])
	}
}

func (c *Channel) Upload(video *Video) {
	video.SetChannel(c)
	fmt.Println(fmt.Sprintf("頻道 %s 上架了一則新影片 `%s`", c.name, video.title))
	c.videos = append(c.videos, video)
	c.Notify()
}
