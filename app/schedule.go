package app

import (
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
)

func (a *App) startSchedule() {
	c := cron.New()

	for _, h := range a.config.Hooks {
		if h.Schedule == "" {
			continue
		}

		hookId := h.Id

		log.WithFields(log.Fields{
			"HookId":   hookId,
			"Schedule": h.Schedule,
		}).Debug("Scheduled")

		c.AddFunc(h.Schedule, func() {
			a.RunHook(hookId, nil)
		})
	}

	c.Start()
}
