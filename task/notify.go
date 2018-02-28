package task

import (
	"errors"
	"github.com/ddliu/webhook/contact"
	"github.com/ddliu/webhook/context"
	"github.com/ddliu/webhook/notifier"
)

type Notify struct {
}

func (t *Notify) GetId() string {
	return "notify"
}

type NotifyInput struct {
	Receiver string
	Title    string
	Content  string
}

func (t *Notify) Run(appContext *context.Context, requestContext *context.Context, inputContext *context.Context) error {
	var input NotifyInput

	if err := inputContext.Unmarshal(&input); err != nil {
		return err
	}

	v, _ := appContext.GetValue("contact_book")
	cb := v.(*contact.ContactBook)
	to := cb.GetById(input.Receiver)
	if to == nil {
		return errors.New("Receiver does not exist: " + input.Receiver)
	}

	notifiers := notifier.MatchNotifiers(to)
	if len(notifiers) == 0 {
		return errors.New("No one to send")
	}

	for _, n := range notifiers {
		err := n.Notify(to, input.Title, input.Content)
		if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	registerTask(&Notify{})
}
