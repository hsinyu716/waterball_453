package common

type Subscriber struct {
	name string
}

type ISubscriberObserver interface {
	Notify(video *Video)
	SetName(name string)
	GetName() string
	Subscribe(channel *Channel)
	UnSubscribe(channel *Channel)
}

func (s *Subscriber) SetName(name string) {
	s.name = name
}

func (s *Subscriber) GetName() string {
	return s.name
}
