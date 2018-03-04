package app

// App configuration
type Config struct {
	Listen    string
	Verbose   bool
	Vars      map[string]interface{}
	Hooks     []HookConfig
	Contacts  []ContactConfig
	Notifiers []NotifierConfig
}

type HookConfig struct {
	Id         string
	Type       string
	Conditions map[string]interface{}
	Tasks      []TaskConfig
	Schedule   string
}

type TaskConfig struct {
	Type   string
	SaveAs string
	Params map[string]interface{}
}

type ContactConfig map[string]interface{}

type NotifierConfig map[string]interface{}

func (c Config) getHookConfigById(hookId string) *HookConfig {
	for _, h := range c.Hooks {
		if hookId == h.Id {
			return &h
		}
	}

	return nil
}
