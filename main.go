package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"github.com/potterhe/go-sample/worker"
)

func main() {

	ch := make(chan int64, 5)

	w := worker.EchoWorker{N: 2, In: ch}
	w2 := worker.GenDataWorker{N: 1, Out: ch}

	master := worker.NewMaster()
	master.Register(&w)
	master.Register(&w2)
	master.Serve()

	channelSignal := make(chan os.Signal)
	signal.Notify(channelSignal, os.Interrupt, syscall.SIGTERM)
	log.Println("started")
	<-channelSignal
	log.Println("end")
	os.Exit(0)
}
