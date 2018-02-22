// Webhook receiver
package receiver

import (
	"github.com/kataras/iris"
)

type ReceiverInterface interface {
	// Get receiver id
	GetId() string

	// Receive from context
	Receive(iris.Context) error

	// Check receiver context
	Match(iris.Context) bool
}

var receivers = make(map[string]ReceiverInterface)

func RegisterReceiver(r ReceiverInterface) {
	receivers[r.GetId()] = r
}

func GetReceiver(id string) ReceiverInterface {
	return receivers[id]
}
