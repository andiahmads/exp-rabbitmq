package pool

import "log"

var WorkerChannel = make(chan chan Work)

type Collector struct {
	Work chan Work
	End  chan bool
}

func StartDispatcher(workerCount int) Collector {
	var i int
	var workers []Worker
	input := make(chan Work) // channel to received work
	end := make(chan bool)   // channel to spin down workers
	collector := Collector{Work: input, End: end}

	for i < workerCount {
		i++
		log.Println("Startring Worker: ", i) // log count worker
		worker := Worker{
			ID:            i,
			Channel:       make(chan Work),
			WorkerChannel: WorkerChannel,
			End:           make(chan bool),
		}

		worker.Start()
		workers = append(workers, worker)

	}

	// start collector
	go func() {
		for {
			select {
			case <-end:
				for _, w := range workers {
					w.Stop()
				}
				return
			case work := <-input:
				worker := <-WorkerChannel // wait for available channel
				worker <- work            // dispatch work to worker
			}

		}
	}()

	return collector
}
