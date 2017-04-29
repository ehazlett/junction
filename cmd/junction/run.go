package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/codegangsta/cli"
	"github.com/ehazlett/junction/server"
	"github.com/ehazlett/junction/utils"
	"github.com/ehazlett/junction/version"
	"github.com/sirupsen/logrus"
)

func getHostname() string {
	h, err := os.Hostname()
	if err != nil {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}
	return h
}

func getLocalIP() string {
	ip, err := utils.GetDefaultIP()
	if err != nil {
		return "127.0.0.1"
	}

	return ip
}

var runCommand = cli.Command{
	Name:   "run",
	Usage:  "start server",
	Action: runAction,
	Before: func(c *cli.Context) error {
		if c.GlobalBool("debug") {
			logrus.SetFormatter(&logrus.TextFormatter{})
			logrus.SetLevel(logrus.DebugLevel)
		}
		return nil
	},
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "node-name, n",
			Usage: "Name of node",
			Value: getHostname(),
		},
		cli.StringFlag{
			Name:  "listen-addr, l",
			Usage: "Listen address",
			Value: fmt.Sprintf("%s:8080", getLocalIP()),
		},
		cli.StringFlag{
			Name:  "control-socket, c",
			Usage: "Plugin control socket",
			Value: "unix:///var/run/junction.sock",
		},
		cli.StringFlag{
			Name:  "bind-addr, b",
			Usage: "Bind address for node communication",
			Value: fmt.Sprintf("%s:7946", getLocalIP()),
		},
		cli.StringFlag{
			Name:  "advertise-addr, a",
			Usage: "Advertise address for node communication",
			Value: fmt.Sprintf("%s:7946", getLocalIP()),
		},
		cli.StringFlag{
			Name:  "join-addr, j",
			Usage: "Address of node to join (optional)",
			Value: "",
		},
		cli.DurationFlag{
			Name:  "node-timeout, t",
			Usage: "Timeout for nodes",
			Value: time.Second * 10,
		},
		cli.BoolFlag{
			Name:  "cluster-debug",
			Usage: "Enable lower level debug for cluster communication",
		},
	},
}

func runAction(c *cli.Context) error {
	logrus.Info(version.FullVersion())

	cfg := &server.Config{
		Name:          c.String("node-name"),
		ListenAddr:    c.String("listen-addr"),
		ControlSocket: c.String("control-socket"),
		BindAddr:      c.String("bind-addr"),
		AdvertiseAddr: c.String("advertise-addr"),
		JoinAddr:      c.String("join-addr"),
		NodeTimeout:   c.Duration("node-timeout"),
		Debug:         c.Bool("cluster-debug"),
	}

	srv, err := server.NewServer(cfg)
	if err != nil {
		return err
	}

	// handle interrupt and switch to original ns
	interruptChan := make(chan os.Signal, 2)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-interruptChan
		srv.Stop()
	}()

	logrus.WithFields(logrus.Fields{
		"version":       version.Version(),
		"controlSocket": c.String("control-socket"),
		"listenAddr":    c.String("listen-addr"),
	}).Info(version.Name())

	if err := srv.Run(); err != nil {
		return err
	}

	return nil
}
