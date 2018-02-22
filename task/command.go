package task

type Command struct {
}

func (t *Command) GetId() string {
	return "command"
}

func (t *Command) Run(c *TaskContext, i TaskInput) error {
	return nil
}

func init() {
	registerTask(&Command{})
}
