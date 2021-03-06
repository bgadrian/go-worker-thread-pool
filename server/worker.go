package server

import (
	"time"
)

//Payload the data to be processed, in this example is a string
type Payload struct {
	Magic string
}

//ProcessPayload what to do with it?
type ProcessPayload func(w *Worker, p Job) error

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
	ID                 string //for debuging purposes
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
func (w *Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			//this is the worker's way to say "I'm free! give me a job!"
			w.WorkerPool <- w.JobChannel
			clientStream(w, "IDLE")

			select {
			case job := <-w.JobChannel:
				// we have received a work request.
				clientStream(w, "received: "+job.Payload.Magic)
				time.Sleep(2 * time.Second) //fake time

				clientStream(w, "processing: "+job.Payload.Magic)
				err := w.process(w, job)

				if err == nil {
					clientStream(w, "finished: "+job.Payload.Magic)
				} else {
					clientStream(w, "failed: "+job.Payload.Magic)
				}

				time.Sleep(2 * time.Second) //fake time

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
