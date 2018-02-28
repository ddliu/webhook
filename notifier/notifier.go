package notifier

import (
	"github.com/ddliu/webhook/contact"
	"github.com/ddliu/webhook/context"
)

type Notifier interface {
	GetId() string
	Config(*context.Context)
	Notify(*contact.Contact, string, string) error
	IsMatch(*contact.Contact) bool
}

var notifiers = make(map[string]Notifier)

func RegisterNotifier(n Notifier) {
	notifiers[n.GetId()] = n
}

func GetNotifier(id string) (Notifier, bool) {
	n, ok := notifiers[id]
	return n, ok
}

func MatchNotifiers(c *contact.Contact) []Notifier {
	var result []Notifier
	for _, v := range notifiers {
		if v.IsMatch(c) {
			result = append(result, v)
		}
	}

	return result
}
