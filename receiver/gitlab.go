package receiver

import (
	"github.com/kataras/iris"
)

type Gitlab struct{}

func (r *Gitlab) GetId() string {
	return "gitlab"
}

func (r *Gitlab) Receive(c iris.Context) error {
	return nil
}

func (r *Gitlab) Match(c iris.Context) bool {
	return true
}

func init() {
	RegisterReceiver(&Gitlab{})
}
