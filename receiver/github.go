package receiver

import (
	"github.com/kataras/iris"
)

type Github struct{}

func (r *Github) GetId() string {
	return "github"
}

func (r *Github) Receive(c iris.Context) error {
	return nil
}

func (r *Github) Match(c iris.Context) bool {
	return true
}

func init() {
	RegisterReceiver(&Github{})
}
