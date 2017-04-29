package server

import "github.com/sirupsen/logrus"

func (s *Server) Stop() error {
	logrus.Info("stopping server")
	defer func() {
		s.stopChan <- true
	}()

	return nil
}
