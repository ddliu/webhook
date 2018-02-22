package receiver

import (
	"github.com/kataras/iris"
)

// Auto receiver can detect and choose receivers automaticly
type Auto struct {
}

func (r *Auto) GetId() string {
	return "auto"
}

func (r *Auto) Receive(c iris.Context) error {
	return r.GetMatchedReceiver(c).Receive(c)
}

func (r *Auto) Match(c iris.Context) bool {
	return true
}

func (r *Auto) GetMatchedReceiver(c iris.Context) ReceiverInterface {
	for id, receiver := range receivers {
		if receiver.Match(c) {
			return receiver
		}

		if id == "auto" || id == "unknown" {
			continue
		}
	}

	return GetReceiver("unknown")
}

func init() {
	RegisterReceiver(&Auto{})
}
