package receiver

import (
	"github.com/ddliu/webhook/context"
	"net/http"
)

// See: https://gogs.io/docs/features/webhook
type Gogs struct{}

func (r *Gogs) GetId() string {
	return "gogs"
}

func (r *Gogs) Receive(c *context.Context, req *http.Request) error {
	return nil
}

func (r *Gogs) Match(c *context.Context, req *http.Request) bool {
	return c.Exist("headers.X-Gogs-Delivery")
}

func init() {
	RegisterReceiver(&Gogs{})
}
