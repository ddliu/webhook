package gitlab

import (
	"github.com/ddliu/webhook/context"
	"github.com/ddliu/webhook/receiver"
	"net/http"
)

type ReceiverGitlab struct{}

func (r *ReceiverGitlab) GetId() string {
	return "gitlab"
}

func (r *ReceiverGitlab) Receive(c *context.Context, req *http.Request) (*context.Context, error) {
	return nil, nil
}

func (r *ReceiverGitlab) Match(c *context.Context, req *http.Request) bool {
	return false
}

func init() {
	receiver.RegisterReceiver(&ReceiverGitlab{})
}
