package task

import (
	"github.com/ddliu/webhook/context"
	log "github.com/sirupsen/logrus"
	"time"
)

type TaskItem struct {
	Task  TaskInterface
	Input *context.Context
}

type TaskRunner struct {
	appContext     *context.Context
	requestContext *context.Context
	tasks          []TaskItem
}

func NewTaskRunner(appContext *context.Context, requestContext *context.Context) *TaskRunner {
	return &TaskRunner{
		appContext:     appContext,
		requestContext: requestContext,
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
		startTime := time.Now()
		log.WithFields(log.Fields{
			"Type": task.Task.GetId(),
		}).Debug("Begin task")
		err := task.Task.Run(t.appContext, t.requestContext, task.Input)
		spentTime := time.Now().Sub(startTime)
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
	Run(*context.Context, *context.Context, *context.Context) error
}

var registeredTasks = make(map[string]TaskInterface)

func registerTask(t TaskInterface) {
	registeredTasks[t.GetId()] = t
}

func GetTaskById(taskId string) TaskInterface {
	return registeredTasks[taskId]
}
