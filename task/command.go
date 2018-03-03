package task

import (
	"github.com/ddliu/webhook/context"
)

type Command struct {
}

func (t *Command) GetId() string {
	return "command"
}

func (t *Command) Run(ctx *context.Context) (*context.Context, error) {
	return nil, nil
}

func init() {
	RegisterTask(&Command{})
}
