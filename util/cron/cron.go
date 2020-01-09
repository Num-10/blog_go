package cron

import (
	"blog_go/controller"
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func CronSetup()  {
	fmt.Println("cron starting...")

	c := cron.New()

	c.AddFunc("0 0 2 * * *", func() {
		//model.TestCreateUser()
		controller.StatisticsViewCount()
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <- t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
