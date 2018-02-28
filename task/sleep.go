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

func (s *Sleep) Run(appContext *context.Context, requestContext *context.Context, inputContext *context.Context) error {
	v, _ := inputContext.GetValue("DurationMS")
	ms := cast.ToInt(v)

	if ms <= 0 {
		ms = 1000
	}

	time.Sleep(time.Millisecond * time.Duration(ms))
	return nil
}

func init() {
	registerTask(&Sleep{})
}
