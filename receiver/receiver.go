// Webhook receiver
package receiver

import (
	"github.com/ddliu/webhook/context"
	"net/http"
)

type ReceiverInterface interface {
	// Get receiver id
	GetId() string

	// Receive from context
	Receive(*context.Context, *http.Request) (*context.Context, error)

	// Check receiver context
	Match(*context.Context, *http.Request) bool
}

var receivers = make(map[string]ReceiverInterface)

func RegisterReceiver(r ReceiverInterface) {
	receivers[r.GetId()] = r
}

func GetReceiver(id string) ReceiverInterface {
	return receivers[id]
}
