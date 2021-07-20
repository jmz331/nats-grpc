package main

import (
	"github.com/google/uuid"
	"github.com/jmz331/nats-grpc/internal/pkg/logger"
	"github.com/nats-io/nats.go"
	"sync"
	"time"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	defer nc.Drain()

	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			send(nc)
			wg.Done()
		}()
	}
	wg.Wait()
}

func send(nc *nats.Conn) {
	id := uuid.New().String()
	request, err := nc.Request("help", []byte(id), time.Second*10)
	if err != nil {
		panic(err)
	}
	logger.Info("%s receive:%s", id, string(request.Data))
}
