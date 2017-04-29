package server

import (
	"sync"

	"github.com/ehazlett/libdiscover"
)

const (
	NodeJoinEvent  = "node-join"
	NodeLeaveEvent = "node-leave"
	NodeInfoEvent  = "node-info"
)

type Server struct {
	config   *Config
	stopChan chan bool
	discover *libdiscover.Discover
	mu       *sync.Mutex
}

func NewServer(cfg *Config) (*Server, error) {
	ch := make(chan bool)
	return &Server{
		stopChan: ch,
		config:   cfg,
		mu:       &sync.Mutex{},
	}, nil
}
