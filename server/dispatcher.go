package server

//Dispatcher Keeps the things together, a controller.
type Dispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	WorkerPool      chan chan Job
	MaxWorkers      int
	Workers         []Worker
	JobQueue        chan Job // A buffered channel that we can send work requests on.
	processFunction ProcessPayload
}

//NewDispatcher Create a new instance
func NewDispatcher(maxWorkers, maxQueue int, p ProcessPayload) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	jobq := make(chan Job, maxQueue)
	return &Dispatcher{
		WorkerPool:      pool,
		MaxWorkers:      maxWorkers,
		JobQueue:        jobq,
		processFunction: p,
	}
}

//Run start listening and hire the workers
func (d *Dispatcher) Run() {
	// starting n number of workers
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(d.WorkerPool, d.processFunction)
		worker.ID = "worker " + string(i+1)
		d.Workers = append(d.Workers, worker)
		worker.Start()
	}

	go d.dispatch()
}

//Stop Why you ever want to stop?
func (d *Dispatcher) Stop() {
	for _, w := range d.Workers {
		w.Stop()
	}
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			// a job request has been received
			go func(job Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool

				// dispatch the job to the worker job channel
				jobChannel <- job
			}(job)
		}
	}
}

//DispatchJob send a job to the workers
func (d *Dispatcher) DispatchJob(j *Job) {
	go func() {
		// Push the work onto the queue.
		d.JobQueue <- *j
	}()
}
