// Application
package app

import (
	"encoding/json"
	"errors"
	"github.com/ddliu/webhook/context"
	"github.com/ddliu/webhook/receiver"
	"github.com/ddliu/webhook/task"
	"io/ioutil"
	"net/http"
)

type App struct {
	config Config
}

func NewApp(configPath string) *App {
	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	var c Config
	err = json.Unmarshal(b, &c)
	if err != nil {
		panic(err)
	}

	_app = &App{
		config: c,
	}

	return _app
}

func (a *App) Start() {
	startServer()
}

func (a *App) receiveHook(hookConfig *HookConfig, c *context.Context, req *http.Request) error {
	receiverType := hookConfig.Type
	if receiverType == "" {
		receiverType = "auto"
	}

	r := receiver.GetReceiver(receiverType)
	if r == nil {
		return errors.New("Unknown receiver: " + receiverType)
	}

	return r.Receive(c, req)
}

func (a *App) runHook(hookId string, req *http.Request) error {
	ctx := &context.Context{}
	buildContextFromRequest(ctx, req)

	hookConfig := a.config.getHookConfigById(hookId)
	if hookConfig == nil {
		return errors.New("Hook not exist")
	}

	if err := a.receiveHook(hookConfig, ctx, req); err != nil {
		return err
	}

	taskRunner := task.NewTaskRunner(nil)
	for _, taskConfig := range hookConfig.Tasks {
		t := task.GetTaskById(taskConfig.Type)
		if t == nil {
			return errors.New("Unknown task: " + taskConfig.Type)
		}

		taskRunner.Add(t, task.NewTaskInput(taskConfig.Params))
	}

	return taskRunner.Run()
}

var _app *App

func GetApp() *App {
	return _app
}
