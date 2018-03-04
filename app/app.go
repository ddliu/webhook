// Application
package app

import (
	"encoding/json"
	"errors"
	"github.com/ddliu/webhook/contact"
	"github.com/ddliu/webhook/context"
	"github.com/ddliu/webhook/notifier"
	"github.com/ddliu/webhook/receiver"
	"github.com/ddliu/webhook/task"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type App struct {
	config      Config
	contactBook contact.ContactBook
	notifiers   []notifier.Notifier
	appContext  *context.Context
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

	_app.init()

	return _app
}

func (a *App) init() {
	if a.config.Verbose {
		log.SetLevel(log.DebugLevel)
	}

	c := context.New(nil)

	// Contact
	cb := &contact.ContactBook{}
	for _, contactConfig := range a.config.Contacts {
		ct := &contact.Contact{}
		ct.Properties = make(map[string]interface{})
		for contactKey, contactValue := range contactConfig {
			switch contactKey {
			case "Id":
				ct.Id = contactValue.(string)
			case "Groups":
				var groups []string
				var groupsRaw = contactValue.([]interface{})
				for _, v := range groupsRaw {
					groups = append(groups, v.(string))
				}
				ct.Groups = groups
			default:
				ct.Properties[contactKey] = contactValue
			}
		}
		cb.AddContact(ct)
	}

	c.SetValue("contact_book", cb)
	c.SetValue("vars", a.config.Vars)

	a.appContext = c

	// Notifier
	for _, notifierConfig := range a.config.Notifiers {
		id := notifierConfig["Type"].(string)
		if n, ok := notifier.GetNotifier(id); ok {
			cc := &context.Context{}
			cc.SetValue(".", notifierConfig)
			n.Config(cc)
		}
	}
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

	log.Debug("Received " + receiverType)

	rc, err := r.Receive(c, req)
	if err != nil {
		log.WithFields(log.Fields{
			"Type":    receiverType,
			"Request": c.GetValue("request"),
		}).Error("Receive failed")
		return err
	}

	if rc != nil {
		log.WithFields(log.Fields{
			"Type":     receiverType,
			"Receiver": rc.GetValue("."),
		}).Debug("Received")
	}

	c.SetValue("receiver", rc)

	for k, v := range hookConfig.Conditions {
		if c.GetValue(k) != v {
			return errors.New("Condition not matched")
		}
	}

	return nil
}

func (a *App) runHook(hookId string, req *http.Request) error {
	ctx := context.New(nil)
	requestContext, err := buildContextFromRequest(req)
	if err != nil {
		return err
	}

	ctx.SetValue("request", requestContext)
	ctx.SetValue("app", a.appContext)

	hookConfig := a.config.getHookConfigById(hookId)
	if hookConfig == nil {
		return errors.New("Hook not exist")
	}

	if err := a.receiveHook(hookConfig, ctx, req); err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"HookId":   hookId,
		"HookType": hookConfig.Type,
	}).Debug("Received hook")

	taskRunner := task.NewTaskRunner(ctx)
	for _, taskConfig := range hookConfig.Tasks {
		t := task.GetTaskById(taskConfig.Type)
		if t == nil {
			return errors.New("Unknown task: " + taskConfig.Type)
		}

		input := context.New(nil)
		input.SetValue(".", taskConfig.Params)
		var item task.TaskItem
		item.Input = input
		item.Task = t
		item.SaveAs = taskConfig.SaveAs

		taskRunner.Add(item)
	}

	return taskRunner.Run()
}

var _app *App

func GetApp() *App {
	return _app
}
