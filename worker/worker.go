package worker

import (
	"errors"
	"fmt"
	"time"
)

type Job interface {
	Run()
}

type Worker struct {
	entries       []*Entry
	nextID        int
	channelWorker chan *Entry
}

type Entry struct {
	id    int
	name  string
	job   Job
	error error
}

func New() *Worker {
	return &Worker{
		entries:       nil,
		nextID:        0,
		channelWorker: make(chan *Entry, 0),
	}
}

func (worker *Worker) Start() {

	for _, entries := range worker.entries {
		time.Sleep(time.Second)
		go worker.runEntry(entries)
	}

	for errChannel := range worker.channelWorker {
		fmt.Println(fmt.Sprintf("Worker:Restart: JobID: %d, JobName: %s", errChannel.id, errChannel.name))
		time.Sleep(time.Second)
		go worker.runEntry(errChannel)
	}

}

func (worker *Worker) Add(nameJob string, job Job) {
	worker.nextID++
	worker.entries = append(worker.entries, &Entry{
		id:   worker.nextID,
		name: nameJob,
		job:  job,
	})

	worker.channelWorker = make(chan *Entry, len(worker.entries))
}

func (worker *Worker) runEntry(entry *Entry) {
	defer func() {
		if r := recover(); r != nil {
			entryChannel := &Entry{
				id:   entry.id,
				name: entry.name,
				job:  entry.job,
			}

			if err, ok := r.(error); ok {
				entryChannel.error = err
			} else {
				entryChannel.error = errors.New("not found error")
			}

			fmt.Println(fmt.Sprintf("Worker:Stop: JobID: %d, JobName: %s, Error: %v", entryChannel.id, entryChannel.name, entryChannel.error))
			worker.channelWorker <- entryChannel
		}
	}()

	fmt.Println(fmt.Sprintf("Worker:Start: JobID: %d, JobName: %s", entry.id, entry.name))
	entry.job.Run()
}
