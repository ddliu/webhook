package receiver

import (
	"github.com/ddliu/webhook/context"
	"net/http"
)

type Gitlab struct{}

func (r *Gitlab) GetId() string {
	return "gitlab"
}

func (r *Gitlab) Receive(c *context.Context, req *http.Request) error {
	return nil
}

func (r *Gitlab) Match(c *context.Context, req *http.Request) bool {
	return true
}

func init() {
	RegisterReceiver(&Gitlab{})
}
