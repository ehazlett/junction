package server

import (
	"encoding/json"

	"github.com/ehazlett/libdiscover"
	"github.com/sirupsen/logrus"
)

func (s *Server) eventHandler(e libdiscover.Event) error {
	logrus.WithFields(logrus.Fields{
		"name":    e.Name,
		"payload": string(e.Payload),
		"event":   e.Name,
	}).Debug("cluster event")

	switch e.Name {
	case NodeJoinEvent:
	case NodeLeaveEvent:
		return s.handleLeaveEvent(e)
	case NodeInfoEvent:
		return s.handleInfoEvent(e)
	}

	return nil
}

func getNodeConfig(data []byte) (*Config, error) {
	var cfg *Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

func (s *Server) handleInfoEvent(e libdiscover.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	cfg, err := getNodeConfig(e.Payload)
	if err != nil {
		return err
	}
	// don't add self
	if cfg.Name == s.config.Name {
		return nil
	}

	return nil
}

func (s *Server) handleLeaveEvent(e libdiscover.Event) error {
	cfg, err := getNodeConfig(e.Payload)
	if err != nil {
		return err
	}
	logrus.WithFields(logrus.Fields{
		"node": cfg.Name,
	}).Debug("node left the cluster")
	return nil
}
