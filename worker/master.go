package worker

import (
	"time"
	"log"
	"sync"
)

type Worker interface {
	Serve()
	Workers() int
	ShutDown()
}

type Master struct {
	shutdown bool
	workers []Worker
	exit chan Worker
	wait sync.WaitGroup
}

func NewMaster() *Master {
	m := &Master{
		shutdown: false,
		workers: make([]Worker, 0),
		exit: make(chan Worker, 10),
		wait: sync.WaitGroup{},
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

func (m *Master) ShutDown () {
	m.shutdown = true
	for _, w := range m.workers {
		w.ShutDown()
	}
}

func (m *Master) Wait () {
	m.wait.Wait()
}

func (m *Master) spawn (w Worker) {

	go func() {
		m.wait.Add(1)

		defer func() {
			if x := recover(); x != nil {
			}
			m.wait.Done()
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
			if !m.shutdown {
				m.spawn(w)
			}
		case <-time.After(5 * time.Second):
			break
		}
	}
}
