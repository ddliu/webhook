package receiver

import (
	"github.com/ddliu/webhook/context"
	"net/http"
)

type Github struct{}

func (r *Github) GetId() string {
	return "github"
}

func (r *Github) Receive(c *context.Context, req *http.Request) error {
	return nil
}

func (r *Github) Match(c *context.Context, req *http.Request) bool {
	return true
}

func init() {
	RegisterReceiver(&Github{})
}
