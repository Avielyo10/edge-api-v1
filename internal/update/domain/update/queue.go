package update

// Queue is a queue for updates.
type Queue struct {
	queue chan UpdatesInterface
}

// NewUpdateQueue returns a new update queue.
func NewUpdateQueue() *Queue {
	return &Queue{queue: make(chan UpdatesInterface)}
}

// Add adds an update to the queue.
func (uq *Queue) Add(update UpdatesInterface) {
	uq.queue <- update
}

// Get returns the next update in the queue.
func (uq *Queue) Get() UpdatesInterface {
	return <-uq.queue
}

// Close closes the update queue.
func (uq *Queue) Close() {
	close(uq.queue)
}
