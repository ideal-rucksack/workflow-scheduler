package main

import (
	"github.com/go-co-op/gocron/v2"
	"log"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	job, err := scheduler.NewJob(
		gocron.DurationJob(2*time.Second),
		gocron.NewTask(
			func(a string, b int) {
				command := exec.Command("bash", "-c", "echo "+a+" 和 "+strconv.Itoa(b))
				output, err := command.CombinedOutput()
				if err != nil {
					log.Printf("Failed to execute script: %s, error: %s", "bash -c echo "+a+" 和 "+strconv.Itoa(b), err)
				}
				log.Printf("Output: %s\n", output)
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
	case <-time.After(10 * time.Second):
	}

	err = scheduler.Shutdown()
	if err != nil {
		panic(err)
	}
}
