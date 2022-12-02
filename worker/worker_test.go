package worker

import (
	"fmt"
	"testing"
	"time"
)

type Broker struct {
	Job string
}

type Broker2 struct {
	Job string
}

func (broker *Broker) Run() {
	count := 0
	for true {
		count++
		if count >= 30 {
			break
		}

		if count == 10 {
			panic(fmt.Errorf("%s: restart", broker.Job))
		}

		time.Sleep(time.Second * 2)
		fmt.Println(fmt.Sprintf("%s:%d", broker.Job, count))
	}
}

func (broker *Broker2) Run() {
	count := 0
	for true {
		count++
		if count >= 30 {
			break
		}

		if count == 10 {
			panic(fmt.Errorf("%s: restart", broker.Job))
		}

		time.Sleep(time.Second * 2)
		fmt.Println(fmt.Sprintf("%s:%d - job other", broker.Job, count))
	}
}

func TestNew(t *testing.T) {

	jobA := &Broker{Job: "JobA"}
	jobB := &Broker2{Job: "JobB"}

	worker := New()
	worker.Add("Job A", jobA)
	worker.Add("Job B", jobB)
	go worker.Start()

	time.Sleep(time.Hour)
}
