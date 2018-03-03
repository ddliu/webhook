package task

import (
	"github.com/ddliu/webhook/context"
	log "github.com/sirupsen/logrus"
	"time"
)

type TaskItem struct {
	Task   TaskInterface
	Input  *context.Context
	SaveAs string
}

type TaskRunner struct {
	ctx   *context.Context
	tasks []TaskItem
}

func NewTaskRunner(ctx *context.Context) *TaskRunner {
	return &TaskRunner{
		ctx: ctx,
	}
}

func (t *TaskRunner) Add(item TaskItem) {
	t.tasks = append(t.tasks, item)
}

func (t *TaskRunner) Run() error {
	for _, task := range t.tasks {
		startTime := time.Now()
		log.WithFields(log.Fields{
			"Type": task.Task.GetId(),
		}).Debug("Begin task")

		t.ctx.SetValue("task.input", task.Input)

		output, err := task.Task.Run(t.ctx)
		spentTime := time.Now().Sub(startTime)

		t.ctx.SetValue("task.output.last", output)
		if task.SaveAs != "" {
			t.ctx.SetValue("task.output."+task.SaveAs, output)
		}

		if err != nil {
			log.WithFields(log.Fields{
				"Type":      task.Task.GetId(),
				"SpentTime": spentTime,
			}).Error("Run task failed: " + err.Error())

			return err
		} else {
			log.WithFields(log.Fields{
				"Type":      task.Task.GetId(),
				"SpentTime": spentTime,
			}).Info("Run task success")
		}
	}

	return nil
}

type TaskInterface interface {
	GetId() string
	Run(*context.Context) (*context.Context, error)
}

var registeredTasks = make(map[string]TaskInterface)

func RegisterTask(t TaskInterface) {
	registeredTasks[t.GetId()] = t
}

func GetTaskById(taskId string) TaskInterface {
	return registeredTasks[taskId]
}
