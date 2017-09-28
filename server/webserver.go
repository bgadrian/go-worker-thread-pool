package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

//StartServer start an web server, which can be used to control the jobs
func StartServer(d *Dispatcher) {
	//enclosure so we can send the dispatcher, avoiding global variables
	wrapper := func(w http.ResponseWriter, r *http.Request) {
		log.Println("received /job request")

		payloadHandler(d, w, r)
	}

	http.Handle("/", http.FileServer(http.Dir("./client")))
	http.HandleFunc("/job", wrapper)
	http.HandleFunc("/ws", websocketHandler)
	http.ListenAndServe(":8080", nil)
}

//payloadHandler Adding jobs/tasks from web requests
func payloadHandler(d *Dispatcher, w http.ResponseWriter, r *http.Request) {
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

//Use a mutex fo production, and GorillaWebsocket
var wsList []*websocket.Conn

//WorkerUpdate worker status update msg to clients
type WorkerUpdate struct {
	WorkerID string `json:"WorkerID"`
	Status   string `json:"Status"`
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	wsList = append(wsList, ws)
}

//clientStream send status update for all workers to all clients
func clientStream(w *Worker, status string) {
	m := WorkerUpdate{w.ID, status}
	// log.Println("stream to clients", wsList, m)

	for _, ws := range wsList {
		err := ws.WriteJSON(m)

		if err != nil {
			//remove connection from wsList
			log.Println(err)
		}
	}
}
