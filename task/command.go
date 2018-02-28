package task

import (
	"github.com/ddliu/webhook/context"
)

type Command struct {
}

func (t *Command) GetId() string {
	return "command"
}

func (t *Command) Run(appContext *context.Context, requestContext *context.Context, inputContext *context.Context) error {
	return nil
}

func init() {
	registerTask(&Command{})
}
