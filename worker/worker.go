package worker

import (
	"log"
	"time"
	"math/rand"
)

type EchoWorker struct {
	N int
	In <- chan int64
}

func (w *EchoWorker) Serve () {
	for {
		select {
		case i := <-w.In:
			w.handle(i)

		// 5秒后超时
		case <-time.After(5 * time.Second):
			break
		}
	}
}

func (w *EchoWorker) handle (i int64) {
	log.Println("echo:", i)
	time.Sleep(2 * time.Second)
	r := rand.Intn(3)
	if r == 0 {
		panic("panic")
	}
}

func (w *EchoWorker) Workers () int {
	return w.N
}

type GenDataWorker struct {
	N int
	Out chan int64
}

func (w *GenDataWorker) Serve () {
	for {
		i := time.Now().Unix()
		log.Println("gen", i)
		w.Out <- i
	}
}

func (w *GenDataWorker) Workers () int {
	return w.N
}
