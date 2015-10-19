package app

import (
	"github.com/lunny/tango"
	"net"
	"net/http"
	"strings"
	"time"
)

type Web struct {
	*tango.Tango
	Host     string
	Port     string
	Protocol string

	listener net.Listener
}

func NewWeb(host, port, protocol string) *Web {
	return &Web{
		Host:     host,
		Port:     port,
		Protocol: protocol,
	}
}

func (w *Web) IsHTTPS() bool {
	return strings.ToLower(w.Protocol) == "https"
}

func (w *Web) Listen() error {
	w.Tango = tango.New([]tango.Handler{
		tango.Return(),
		tango.Param(),
		tango.Contexts(),
	}...)

	addr := w.Host + ":" + w.Port
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	w.listener = listener
	server := &http.Server{Addr: addr, Handler: w.Tango}
	go server.Serve(tcpKeepAliveListener{listener.(*net.TCPListener)})
	return nil
}

func (w *Web) Close() error {
	if w.listener != nil {
		return w.listener.Close()
	}
	return nil
}

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
