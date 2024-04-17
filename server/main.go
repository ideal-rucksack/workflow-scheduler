package main

import (
	"github.com/go-co-op/gocron/v2"
	"log"
	"time"
)

func main() {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	_, err = scheduler.NewJob(
		gocron.DurationJob(2*time.Second),
		gocron.NewTask(
			func() {
				log.Println("任务1")
			}),
	)

	_, err = scheduler.NewJob(
		gocron.DurationJob(2*time.Second),
		gocron.NewTask(
			func() {
				log.Println("任务2")
			}),
	)

	if err != nil {
		panic(err)
	}

	scheduler.Start()

	_, err = scheduler.NewJob(
		gocron.DurationJob(2*time.Second),
		gocron.NewTask(
			func() {
				log.Println("任务3")
			}),
	)

	select {
	case <-time.After(10 * time.Second):
	}

	err = scheduler.Shutdown()
	if err != nil {
		panic(err)
	}
}
