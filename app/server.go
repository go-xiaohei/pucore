package app

import (
	"github.com/lunny/tango"
	"gopkg.in/inconshreveable/log15.v2"
	"net"
	"net/http"
	"time"
)

type Server struct {
	*tango.Tango
	listener net.Listener

	Address string
	Domain  string
	// todo : TLS support
	IsTLS bool
}

func newServer() *Server {
	server := &Server{
		Tango: tango.New([]tango.Handler{
			tango.Return(),
			tango.Param(),
			tango.Contexts(),
		}...),
	}
	return server
}

func (s *Server) Start() {
	// closable server
	ln, err := net.Listen("tcp", s.Address)
	if err != nil {
		log15.Crit("Server.start.fail", "error", err)
	}
	s.listener = ln
	go func() {
		log15.Info("Server.start." + s.Address)
		server := &http.Server{Addr: s.Address, Handler: s.Tango}
		if err := server.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)}); err != nil {
			log15.Crit("Server.start.fail", "error", err)
		}
	}()
}

func (s *Server) Stop() {
	s.listener.Close()
}

// copy from pkg net/http
type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (c net.Conn, err error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(3 * time.Minute)
	return tc, nil
}
