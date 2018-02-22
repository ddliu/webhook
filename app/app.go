// Application
package app

import (
	"encoding/json"
	"errors"
	"github.com/ddliu/webhook/receiver"
	"github.com/ddliu/webhook/task"
	"github.com/kataras/iris"
	"io/ioutil"
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

func (a *App) receiveHook(hookConfig *HookConfig, c iris.Context) error {
	receiverType := hookConfig.Type
	if receiverType == "" {
		receiverType = "auto"
	}

	r := receiver.GetReceiver(receiverType)
	if r == nil {
		return errors.New("Unknown receiver: " + receiverType)
	}

	return r.Receive(c)
}

func (a *App) runHook(hookId string, c iris.Context) error {
	hookConfig := a.config.getHookConfigById(hookId)
	if hookConfig == nil {
		return errors.New("Hook not exist")
	}

	if err := a.receiveHook(hookConfig, c); err != nil {
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
