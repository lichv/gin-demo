package crontabs

import (
	"github.com/robfig/cron"
	"log"
)


func Setup() *cron.Cron {
	c := cron.New()
	c.AddFunc("0 0 * * * *", func() {
		log.Println("Run models.CleanAllTag...")
	})

	c.Run()
	return c
}
