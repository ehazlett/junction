package server

import (
	"net"
	"time"

	"github.com/ehazlett/junction/api"
	"github.com/ehazlett/junction/utils"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	randomIntervalMax = 5
)

func (s *Server) Run() error {
	if err := s.startDiscover(); err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	api.RegisterJunctionManagerServer(grpcServer, s)

	// setup listeners
	for _, addr := range []string{s.config.ControlSocket, s.config.ListenAddr} {
		proto, host, err := utils.GetProtoHost(addr)
		if err != nil {
			return err
		}
		l, err := net.Listen(proto, host)
		if err != nil {
			return err
		}
		defer l.Close()

		go func() {
			if err := grpcServer.Serve(l); err != nil {
				logrus.Error(err)
			}
		}()
	}

	// heartbeat
	go func() {
		t := time.NewTicker(s.config.NodeTimeout / 2)
		for _ = range t.C {
			if err := s.heartbeat(); err != nil {
				logrus.Error(err)
			}
		}
	}()
	// send initial heartbeat
	if err := s.heartbeat(); err != nil {
		logrus.Error(err)
	}

	<-s.stopChan

	if err := s.discover.Stop(); err != nil {
		return err
	}

	return nil
}
