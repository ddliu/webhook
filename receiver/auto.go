package receiver

import (
	"github.com/ddliu/webhook/context"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// Auto receiver can detect and choose receivers automaticly
type Auto struct {
}

func (r *Auto) GetId() string {
	return "auto"
}

func (r *Auto) Receive(c *context.Context, req *http.Request) (*context.Context, error) {
	return r.GetMatchedReceiver(c, req).Receive(c, req)
}

func (r *Auto) Match(c *context.Context, req *http.Request) bool {
	return true
}

func (r *Auto) GetMatchedReceiver(c *context.Context, req *http.Request) (receiver ReceiverInterface) {
	defer func() {
		log.Debug("Receiver detected automatically: " + receiver.GetId())
	}()
	for id, receiver := range receivers {
		if id == "auto" || id == "unknown" {
			continue
		}

		if receiver.Match(c, req) {
			return receiver
		}
	}

	receiver = GetReceiver("unknown")

	return
}

func init() {
	RegisterReceiver(&Auto{})
}
