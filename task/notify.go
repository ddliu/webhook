package task

type Notify struct {
}

func (t *Notify) GetId() string {
	return "notify"
}

func (t *Notify) Run(c *TaskContext, i TaskInput) error {
	return nil
}

func init() {
	registerTask(&Notify{})
}
