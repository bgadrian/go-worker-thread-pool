package main

import (
	"math/rand"
	"time"
)

//Payload the data to be processed, in this example is a string
type Payload struct {
	magic string
}

//ProcessPayload what to do with it?
type ProcessPayload func(p Job) error

// Job represents the task/job to be run, with the payload
type Job struct {
	Payload Payload
}

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool         chan chan Job
	JobChannel         chan Job
	quit               chan bool
	dispatcherJobQueue chan Job
	process            ProcessPayload
}

//NewWorker workers are the foundation of our queue system
func NewWorker(workerPool chan chan Job, p ProcessPayload) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
		process:    p,
	}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			//this is the worker's way to say "I'm free! give me a job!"
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				// we have received a work request.
				w.process(job)

				//simulating a very long time to process
				//so we can understand the process
				time.Sleep(time.Duration(rand.Intn(5)+3) * time.Second)

			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
