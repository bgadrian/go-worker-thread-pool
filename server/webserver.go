package server

import (
	"encoding/json"
	"io"
	"net/http"
)

//StartServer start an web server, which can be used to control the jobs
func StartServer(d *Dispatcher) {
	//enclosure so we can send the dispatcher, avoiding global variables
	wrapper := func(w http.ResponseWriter, r *http.Request) {
		payloadHandler(d, w, r)
	}

	http.Handle("/", http.FileServer(http.Dir("./client")))
	http.HandleFunc("/job", wrapper)
	http.ListenAndServe(":8080", nil)
}

//payloadHandler Adding jobs/tasks from web requests
func payloadHandler(d *Dispatcher, w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var content Payload
	defer r.Body.Close()

	err := json.NewDecoder(io.LimitReader(r.Body, 2048)).Decode(&content)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	d.DispatchJob(&Job{content})

	w.WriteHeader(http.StatusOK)
}
