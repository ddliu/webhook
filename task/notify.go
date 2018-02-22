package task

import (
	"github.com/ddliu/webhook/context"
)

type Notify struct {
}

func (t *Notify) GetId() string {
	return "notify"
}

func (t *Notify) Run(c *context.Context, i *context.Context) error {
	return nil
}

func init() {
	registerTask(&Notify{})
}
