package app

// App configuration
type Config struct {
	Listen string
	Hooks  []HookConfig
}

type HookConfig struct {
	Id    string
	Type  string
	Tasks []TaskConfig
}

type TaskConfig struct {
	Type   string
	Params map[string]interface{}
}

func (c Config) getHookConfigById(hookId string) *HookConfig {
	for _, h := range c.Hooks {
		if hookId == h.Id {
			return &h
		}
	}

	return nil
}
