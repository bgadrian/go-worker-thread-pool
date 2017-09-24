package main

import (
	"encoding/json"
	"net/http"
)

//payloadTest Adding jobs/tasks from other sources
func payloadTest(p Payload, d *Dispatcher) {
	work := Job{Payload: p}

	// Push the work onto the queue.
	d.JobQueue <- work
}

//payloadHandler Adding jobs/tasks from web requests
func payloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Read the body into a string for json decoding
	// var content = &PayloadCollection{}
	// err := json.NewDecoder(io.LimitReader(r.Body, MaxLength)).Decode(&content)

	var content Payload
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(content)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// let's create a job with the payload
	// work := Job{Payload: content}

	// Push the work onto the queue.
	// JobQueue <- work

	w.WriteHeader(http.StatusOK)
}
