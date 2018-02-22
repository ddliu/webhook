package receiver

import (
	"github.com/kataras/iris"
)

// Unknown receiver can receive any context
type Unknown struct {
}

func (r *Unknown) GetId() string {
	return "unknown"
}

func (r *Unknown) Receive(c iris.Context) error {
	return nil
}

func (r *Unknown) Match(c iris.Context) bool {
	return true
}

func init() {
	RegisterReceiver(&Unknown{})
}
