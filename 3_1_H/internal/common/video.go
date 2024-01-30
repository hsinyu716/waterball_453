package common

type Video struct {
	channel     *Channel
	title       string
	description string
	length      int64
	Liker       []ISubscriberObserver
}

func NewVideo(title string, desc string, len int64) *Video {
	return &Video{
		title:       title,
		description: desc,
		length:      len,
	}
}

func (v *Video) SetChannel(channel *Channel) {
	v.channel = channel
}

func (v *Video) Like(subscriber ISubscriberObserver) {
	v.Liker = append(v.Liker, subscriber)
}
