package main

import (
	"github.com/jmz331/nats-grpc/internal/app/demo"
	"github.com/jmz331/nats-grpc/internal/pkg/logger"
	"os"
	"os/signal"
)

func main() {
	srv := demo.NewServer()
	err := srv.Start()
	if err != nil {
		panic(err)
		return
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	logger.Info("server stopped")
}
