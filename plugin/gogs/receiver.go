package Gogs

import (
	"github.com/ddliu/webhook/context"
	"github.com/ddliu/webhook/receiver"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"net/http"
	"strings"
)

// See: https://gogs.io/docs/features/webhook
type ReceiverGogs struct{}

func (r *ReceiverGogs) GetId() string {
	return "gogs"
}

func (r *ReceiverGogs) Receive(c *context.Context, req *http.Request) (*context.Context, error) {
	rc := context.New(nil)

	rc.SetValue("event", c.GetValue("request.headers.X-Gogs-Event"))
	rc.SetValue("signature", c.GetValue("request.headers.X-Gogs-Signature"))

	ref := cast.ToString(c.GetValue("request.payload.ref"))
	branch := strings.TrimPrefix(ref, "refs/heads/")
	rc.SetValue("branch", branch)
	return rc, nil
}

func (r *ReceiverGogs) Match(c *context.Context, req *http.Request) bool {
	log.Debug(c.GetValue("request.headers"))
	return c.Exist("request.headers.X-Gogs-Delivery")
}

func init() {
	receiver.RegisterReceiver(&ReceiverGogs{})
}
