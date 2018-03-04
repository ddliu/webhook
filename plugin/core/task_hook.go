package core

import (
	"github.com/ddliu/webhook/app"
	"github.com/ddliu/webhook/context"
	"github.com/ddliu/webhook/task"
	"github.com/spf13/cast"
)

type Hook struct {
}

func (s *Hook) GetId() string {
	return "hook"
}

func (s *Hook) Run(ctx *context.Context) (*context.Context, error) {
	calledHook := cast.ToString(ctx.GetContext("task.input.HookId"))
	app.GetApp().RunHook(calledHook, nil)

	return nil, nil
}

func init() {
	task.RegisterTask(&Hook{})
}
