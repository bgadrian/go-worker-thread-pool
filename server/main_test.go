package main

import "testing"
import "time"
import "sync"
import "sync/atomic"

func TestMain(t *testing.T) {
	var sent, processed uint64

	process := func(p Job) error {
		atomic.AddUint64(&processed, 1)
		return nil
	}

	dispatcher := NewDispatcher(2, 10, process)
	dispatcher.Run()
	wg := sync.WaitGroup{}

	spam := func() {
		for i := 0; i < 5; i++ {
			payloadTest(Payload{"1"}, dispatcher)
			atomic.AddUint64(&sent, 1)
			time.Sleep(time.Microsecond)
		}
		wg.Done()
	}

	wg.Add(2)
	go spam()
	go spam()

	wg.Wait()
	dispatcher.Stop()
	time.Sleep(time.Second)

	if sent != processed {
		t.Errorf("not all jobs were processed %v left",
			sent-processed)
	}
}
