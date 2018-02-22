package receiver

import (
	"github.com/ddliu/webhook/context"
	"net/http"
)

// Auto receiver can detect and choose receivers automaticly
type Auto struct {
}

func (r *Auto) GetId() string {
	return "auto"
}

func (r *Auto) Receive(c *context.Context, req *http.Request) error {
	return r.GetMatchedReceiver(c, req).Receive(c, req)
}

func (r *Auto) Match(c *context.Context, req *http.Request) bool {
	return true
}

func (r *Auto) GetMatchedReceiver(c *context.Context, req *http.Request) ReceiverInterface {
	for id, receiver := range receivers {
		if receiver.Match(c, req) {
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
