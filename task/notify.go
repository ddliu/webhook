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

func (t *Notify) Run(ctx *context.Context) (*context.Context, error) {
	var input NotifyInput

	inputContext := ctx.GetContext("task.input")

	if err := inputContext.Unmarshal(&input); err != nil {
		return nil, err
	}

	v := ctx.GetValue("app.contact_book")
	cb := v.(*contact.ContactBook)
	to := cb.GetById(input.Receiver)
	if to == nil {
		return nil, errors.New("Receiver does not exist: " + input.Receiver)
	}

	notifiers := notifier.MatchNotifiers(to)
	if len(notifiers) == 0 {
		return nil, errors.New("No one to send")
	}

	title := ctx.Tpl(input.Title)
	content := ctx.Tpl(input.Content)

	var err error
	for _, n := range notifiers {
		e := n.Notify(to, title, content)
		if e != nil {
			err = e
		}
	}

	return nil, err
}

func init() {
	RegisterTask(&Notify{})
}
