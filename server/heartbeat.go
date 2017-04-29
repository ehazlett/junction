package server

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

// heartbeat broadcasts to the cluster the local GRPC connection info
func (s *Server) heartbeat() error {
	logrus.WithFields(logrus.Fields{
		"node":    s.config.Name,
		"address": s.config.ListenAddr,
	}).Debug("heartbeat")
	data, err := json.Marshal(s.config)
	if err != nil {
		return err
	}

	return s.discover.SendEvent(NodeInfoEvent, data, false)
}
