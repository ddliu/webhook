package task

import (
	"time"
)

type Sleep struct {
}

func (s *Sleep) GetId() string {
	return "sleep"
}

func (s *Sleep) Run(c *TaskContext, i TaskInput) error {
	ms := i.GetInt("DurationMS")
	if ms <= 0 {
		ms = 1000
	}

	println("Sleep...")

	time.Sleep(time.Millisecond * time.Duration(ms))
	return nil
}

func init() {
	registerTask(&Sleep{})
}
