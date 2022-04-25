package main

import (
	"time"

	"github.com/Avielyo10/edge-api/internal/common/logs"
	"github.com/Avielyo10/edge-api/internal/edge/domain/image"
	"github.com/Avielyo10/edge-api/internal/update/domain/update"
	log "github.com/sirupsen/logrus"
	"go.uber.org/atomic"
)

func main() {
	// init logger
	logs.Init()
	// Set up a new queue.
	queue := update.NewUpdateQueue()

	var maxWorkers uint32 = 10
	numOfUpdates := *atomic.NewUint32(0)

	for { // keep running until the queue is closed.
		for numOfUpdates.Load() < maxWorkers { // as long as there are less than max workers, keep spawning new workers.
			numOfUpdates.Inc() // increment the number of updates.
			go work(queue, &numOfUpdates)
		}
	}
}

// Workflow:
// 1. Get a job from the queue.
// 2. Check for updates.
// 2.1. If error, rollback.
// 2.2. If updated successfully - return.
// 2.3. If no error, continue.
// 2.3.1. If timeout, rollback.
// 2.3.2. If no timeout, continue.
// 3. Decriment the number of updates.
func work(queue *update.Queue, numOfUpdates *atomic.Uint32) {
	job := queue.Get() // block here if no updates are available in the queue.
	defer numOfUpdates.Dec()
	if err := job.CheckForUpdate(); err != nil {
		log.WithField("error", err).Error("error while checking for updates, rolling back")
		err = job.Rollback()
		if err != nil {
			log.WithField("error", err).Error("error while rolling back")
		}
	} else if job.IsSuccessful() {
		return
	} else {
		select {
		case <-job.(*image.Image).Done():
			err = job.Rollback()
			if err != nil {
				log.WithField("error", err).Error("error while rolling back")
			}
		case <-time.After(1 * time.Minute):
			queue.Add(job)
		}
	}
}
