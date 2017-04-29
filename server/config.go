package server

import "time"

type Config struct {
	Name          string
	ListenAddr    string
	ControlSocket string
	BindAddr      string
	AdvertiseAddr string
	JoinAddr      string
	NodeTimeout   time.Duration
	Debug         bool
}
