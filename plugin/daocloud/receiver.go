package daocloud

import (
	"github.com/ddliu/webhook/context"
	"github.com/ddliu/webhook/receiver"
	"github.com/spf13/cast"
	"net/http"
	"strings"
)

// See: http://docs.daocloud.io/api/#webhook
type ReceiverDaocloud struct{}

func (r *ReceiverDaocloud) GetId() string {
	return "daocloud"
}

func (r *ReceiverDaocloud) Receive(c *context.Context, req *http.Request) (*context.Context, error) {
	rc := context.New(nil)
	rc.SetValue("ref", c.GetValue("request.payload.ref"))
	rc.SetValue("status", c.GetValue("request.payload.status"))

	return rc, nil
}

func (r *ReceiverDaocloud) Match(c *context.Context, req *http.Request) bool {
	return strings.HasPrefix(cast.ToString(c.GetValue("request.payload.image")), "daocloud.io")
}

func init() {
	receiver.RegisterReceiver(&ReceiverDaocloud{})
}
