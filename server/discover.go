package server

import "github.com/ehazlett/libdiscover"

func (s *Server) startDiscover() error {
	cfg := &libdiscover.Config{
		Name:          s.config.Name,
		BindAddr:      s.config.BindAddr,
		AdvertiseAddr: s.config.AdvertiseAddr,
		JoinAddr:      s.config.JoinAddr,
		NodeTimeout:   s.config.NodeTimeout,
		Debug:         s.config.Debug,
		EventHandler:  s.eventHandler,
	}

	d, err := libdiscover.NewDiscover(cfg)
	if err != nil {
		return err
	}
	if err := d.Run(); err != nil {
		return err
	}

	s.discover = d

	return nil
}
