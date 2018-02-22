package task

import (
	// "github.com/ddliu/webhook/tpl"
	"github.com/spf13/cast"
)

type TaskContext struct {
}

type TaskInput struct {
	v map[string]interface{}
}

func NewTaskInput(v map[string]interface{}) TaskInput {
	return TaskInput{
		v: v,
	}
}

func (c *TaskInput) GetString(k string) string {
	return cast.ToString(c.v[k])
}

func (c *TaskInput) GetInt(k string) int {
	return cast.ToInt(c.v[k])
}

type TaskItem struct {
	Task  TaskInterface
	Input TaskInput
}

type TaskRunner struct {
	context *TaskContext
	tasks   []TaskItem
}

func NewTaskRunner(ctx *TaskContext) *TaskRunner {
	return &TaskRunner{
		context: ctx,
	}
}

func (t *TaskRunner) Add(task TaskInterface, input TaskInput) {
	t.tasks = append(t.tasks, TaskItem{
		Task:  task,
		Input: input,
	})
}

func (t *TaskRunner) Run() error {
	for _, task := range t.tasks {
		if err := task.Task.Run(t.context, task.Input); err != nil {
			return err
		}
	}

	return nil
}

type TaskInterface interface {
	GetId() string
	Run(*TaskContext, TaskInput) error
}

var registeredTasks = make(map[string]TaskInterface)

func registerTask(t TaskInterface) {
	registeredTasks[t.GetId()] = t
}

func GetTaskById(taskId string) TaskInterface {
	return registeredTasks[taskId]
}
