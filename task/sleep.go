package task

import (
	"github.com/ddliu/webhook/context"
	"github.com/spf13/cast"
	"time"
)

type Sleep struct {
}

func (s *Sleep) GetId() string {
	return "sleep"
}

func (s *Sleep) Run(ctx *context.Context) (*context.Context, error) {
	v := ctx.GetValue("task.input.DurationMS")
	ms := cast.ToInt(v)

	if ms <= 0 {
		ms = 1000
	}

	time.Sleep(time.Millisecond * time.Duration(ms))
	return nil, nil
}

func init() {
	RegisterTask(&Sleep{})
}
