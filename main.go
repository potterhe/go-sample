package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"github.com/potterhe/go-sample/worker"
)

func main() {

	ch := make(chan int64, 10)

	w := worker.EchoWorker{N: 2, In: ch}
	w2 := worker.GenDataWorker{N: 1, Out: ch}

	master := worker.NewMaster()
	master.Register(&w)
	master.Register(&w2)
	master.Serve()

	notifier := make(chan os.Signal, 1)
	signal.Notify(notifier, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-notifier
		log.Printf("received %v signal, shutting down...", sig)
		master.ShutDown()
	}()

	master.Wait()
	os.Exit(0)
}
