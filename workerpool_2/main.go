package main

import (
	"design-pattern/workerpool_2/pool"
	"design-pattern/workerpool_2/work"
	"log"
)

const WORKER_COUNT = 5
const JOB_COUNT = 100

func main() {
	log.Println("starting application...")
	collector := pool.StartDispatcher(WORKER_COUNT) // start up worker pool

	for i, job := range work.CreateJobs(JOB_COUNT) {
		collector.Work <- pool.Work{Job: job, ID: i}
	}
}
