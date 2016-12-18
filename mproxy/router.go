package mproxy

import (
	"fmt"
	"github.com/siddontang/go-mysql/server"
	"net"
)

// Router holds a router app instance
type Router struct {
	conf          Config
	connPoolCache *ConnPoolCache
	port          string
	s             net.Listener
}

// NewRouter constructs a Router instance
func NewRouter(conf Config) (*Router, error) {
	port := fmt.Sprintf(":%d", conf.Auth.Port)
	fmt.Println("Listening to port: ", port)
	listenSocket, err := net.Listen("tcp", port)
	if err != nil {
		return nil, err
	}

	return &Router{
		conf:          conf,
		connPoolCache: NewConnPoolCache(conf),
		port:          port,
		s:             listenSocket,
	}, nil
}

// Start the mysql connection handler
func (r *Router) Start() {
	authUser := r.conf.Auth.Username
	authPass := r.conf.Auth.Password

	for {
		if c, err := r.s.Accept(); err == nil {
			handler := &ProxyHandler{
				connPoolCache: r.connPoolCache,
			}
			go handleTCPConn(c, authUser, authPass, handler)
		}
	}
}

func handleTCPConn(c net.Conn, authUser string, authPass string, handler *ProxyHandler) {
	defer handler.ReturnConn()
	defer c.Close()

	var e error

	if conn, authError := server.NewConn(c, authUser, authPass, handler); authError == nil {
		for {
			if e = conn.HandleCommand(); e != nil {
				return
			}
		}
	}
}
