package task

import (
	"github.com/ddliu/webhook/context"
)

type Command struct {
}

func (t *Command) GetId() string {
	return "command"
}

func (t *Command) Run(c *context.Context, i *context.Context) error {
	return nil
}

func init() {
	registerTask(&Command{})
}
