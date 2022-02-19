package bus

import (
	"fmt"

	"github.com/infraboard/eventbox/apps/event"
)

var (
	publisher  Publisher
	subscriber Subscriber
)

// P todo
func P() Publisher {
	return publisher
}

// S todo
func S() Subscriber {
	return subscriber
}

// Pub bus为全局对象
func Pub(e *event.Event) error {
	if publisher == nil {
		return fmt.Errorf("publisher not initail")
	}

	return publisher.Pub(e.Type.String(), e)
}

// SetPublisher 设置pub
func SetPublisher(p Publisher) {
	publisher = p
}

// Sub bus为全局对象
func Sub(t event.Type, h EventHandler) error {
	if subscriber == nil {
		return fmt.Errorf("subscriber not initial")
	}

	return subscriber.Sub(t.String(), h)
}

// SetSubscriber 设置sub
func SetSubscriber(s Subscriber) {
	subscriber = s
}
