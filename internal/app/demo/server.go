package demo

import (
	"github.com/jmz331/nats-grpc/internal/pkg/logger"
	"github.com/nats-io/nats.go"
)

type Server struct {
	nc *nats.Conn

	subs []*nats.Subscription
}

func (s *Server) Start() error {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return err
	}
	sub, err := nc.Subscribe("help", func(msg *nats.Msg) {
		logger.Info("receive message:%v", string(msg.Data))
		err := nc.Publish(msg.Reply, []byte("I can help!"+string(msg.Data)))
		if err != nil {
			return
		}
	})
	if err != nil {
		return err
	}
	s.subs = append(s.subs, sub)
	s.nc = nc
	return nil
}

func (s *Server) Stop() {
	for _, sub := range s.subs {
		if err := sub.Unsubscribe(); err != nil {
			logger.Info("unable to unsubscribe subscribe:%s, err:%+v", sub.Subject, err)
		}
	}
	if err := s.nc.Drain(); err != nil {
		logger.Info("unable to drain nats connection, err:%+v", err)
	}
}

func NewServer() *Server {
	return &Server{
		subs: make([]*nats.Subscription, 0),
	}
}
