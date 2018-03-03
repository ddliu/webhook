package github

import (
	"github.com/ddliu/webhook/context"
	"github.com/ddliu/webhook/receiver"
	"net/http"
)

type ReceiverGithub struct{}

func (r *ReceiverGithub) GetId() string {
	return "github"
}

func (r *ReceiverGithub) Receive(c *context.Context, req *http.Request) (*context.Context, error) {
	return nil, nil
}

func (r *ReceiverGithub) Match(c *context.Context, req *http.Request) bool {
	return false
}

func init() {
	receiver.RegisterReceiver(&ReceiverGithub{})
}
