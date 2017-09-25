package main

import (
	"errors"
	"flag"
	"log"
	"math/rand"
	"time"

	"github.com/bgadrian/go-worker-thread-pool/server"
)

func main() {
	MaxWorker := flag.Uint("MAX_WORKERS", 10, "max nr of workers")
	MaxQueue := flag.Uint("MAX_QUEUE", 100, "max nr of jobs in queue")
	flag.Parse()

	//every payload (request to /job) from the client is sent here
	processator := func(w *server.Worker, j server.Job) error {
		text := j.Payload.Magic
		if text == "error" {
			return errors.New("error")
		}
		log.Printf("processing '%v' by %v\n", text, w.ID)
		time.Sleep(3 + time.Duration(rand.Intn(3)))
		log.Printf("done processing '%v' by %v\n", text, w.ID)
		return nil
	}

	log.Println("open http://localhost:8080 in your browser & keep this process open.")
	dispatcher := server.NewDispatcher(int(*MaxWorker), int(*MaxQueue), processator)
	dispatcher.Run()
	server.StartServer(dispatcher)
}
