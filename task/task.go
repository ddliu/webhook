package task

import (
	// "github.com/ddliu/webhook/tpl"
	"github.com/ddliu/webhook/context"
)

type TaskItem struct {
	Task  TaskInterface
	Input *context.Context
}

type TaskRunner struct {
	context *context.Context
	tasks   []TaskItem
}

func NewTaskRunner(ctx *context.Context) *TaskRunner {
	return &TaskRunner{
		context: ctx,
	}
}

func (t *TaskRunner) Add(task TaskInterface, input *context.Context) {
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
	Run(*context.Context, *context.Context) error
}

var registeredTasks = make(map[string]TaskInterface)

func registerTask(t TaskInterface) {
	registeredTasks[t.GetId()] = t
}

func GetTaskById(taskId string) TaskInterface {
	return registeredTasks[taskId]
}
