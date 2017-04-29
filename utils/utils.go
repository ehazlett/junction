package utils

import (
	"net"
	"net/url"
	"strings"
)

// GetProtoHost returns a protocol, host and error from the specified address
// TCP will be used as the default protocol if one is not specified
func GetProtoHost(addr string) (string, string, error) {
	// assume and prepend tcp if missing
	if strings.Index(addr, "://") == -1 {
		addr = "tcp://" + addr
	}
	u, err := url.Parse(addr)
	if err != nil {
		return "", "", err
	}

	proto := "tcp"
	host := u.Host
	if strings.Index(u.Scheme, "unix") >= 0 {
		proto = "unix"
		host = u.Path
	}

	return proto, host, nil
}

// GetLocalIP returns the non-loopback outbound IP for the node
func GetDefaultIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().String()
	i := strings.LastIndex(localAddr, ":")

	return localAddr[0:i], nil
}
