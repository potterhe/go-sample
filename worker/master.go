package worker

import (
	"time"
	"log"
)

type Worker interface {
	Serve()
	Workers() int
}

type Master struct {
	shutdown bool
	workers []Worker
	exit chan Worker
}

func NewMaster() *Master {
	m := &Master{
		shutdown: false,
		workers: make([]Worker, 0),
		exit: make(chan Worker, 10),
	}
	return m
}

func (m *Master) Register (w Worker) {
	m.workers = append(m.workers, w)
}

func (m *Master) Serve () {

	go m.monitorWorker()

	for _, w := range m.workers {

		for i := 0; i < w.Workers(); i++ {

			m.spawn(w)
		}
	}
}

func (m *Master) spawn (w Worker) {

	go func() {

		defer func() {
			if x := recover(); x != nil {
			}
			m.exit <- w
		}()

		w.Serve()
	}()
}

func (m *Master) monitorWorker () {

	for {

		select {
		case w := <-m.exit:
			log.Println("recive panic, then spawn", w)
			m.spawn(w)
		case <-time.After(5 * time.Second):
			break
		}
	}
}
