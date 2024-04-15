package main

import (
	"github.com/go-co-op/gocron/v2"
	"time"
)

func main() {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	job, err := scheduler.NewJob(
		gocron.DurationJob(10*time.Second),
		gocron.NewTask(
			func(a string, b int) {
				println(a, b)
			},
			"hello",
			123),
	)

	if err != nil {
		panic(err)
	}

	println(job.ID().ID())

	scheduler.Start()

	select {
	case <-time.After(time.Minute):
	}

	err = scheduler.Shutdown()
	if err != nil {
		panic(err)
	}
}
