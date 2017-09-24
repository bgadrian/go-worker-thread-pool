package main

import (
	"errors"
	"flag"
	"log"
	"math/rand"
	"time"
)

func main() {
	MaxWorker := flag.Uint("MAX_WORKERS", 10, "max nr of workers")
	MaxQueue := flag.Uint("MAX_QUEUE", 100, "max nr of jobs in queue")
	flag.Parse()

	process := func(j Job) error {
		if j.Payload.magic == "error" {
			return errors.New("error")
		}
		log.Printf("processed %v\n", j.Payload.magic)
		time.Sleep(3 + time.Duration(rand.Intn(3)))
		return nil
	}

	dispatcher := NewDispatcher(int(*MaxWorker), int(*MaxQueue), process)
	dispatcher.Run()
}
