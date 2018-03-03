package receiver

import (
	"github.com/ddliu/webhook/context"
	"net/http"
)

// Unknown receiver can receive any context
type Unknown struct {
}

func (r *Unknown) GetId() string {
	return "unknown"
}

func (r *Unknown) Receive(c *context.Context, req *http.Request) (*context.Context, error) {
	return nil, nil
}

func (r *Unknown) Match(c *context.Context, req *http.Request) bool {
	return true
}

func init() {
	RegisterReceiver(&Unknown{})
}
