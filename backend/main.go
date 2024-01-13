package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"todo/pkg/rest"
)

func main() {

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	r := rest.NewRestServer()
	r.StartServer()

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	<-done
}
