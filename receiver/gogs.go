package receiver

import (
	"github.com/kataras/iris"
)

type Gogs struct{}

func (r *Gogs) GetId() string {
	return "gogs"
}

func (r *Gogs) Receive(c iris.Context) error {
	return nil
}

func (r *Gogs) Match(c iris.Context) bool {
	return true
}

func init() {
	RegisterReceiver(&Gogs{})
}
