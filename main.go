package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sync"

	"golang.org/x/net/websocket"
)

type Server struct {
	mu    sync.RWMutex
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Printf("new incoming connection from client: %s\n", ws.RemoteAddr())

	s.mu.Lock()
	s.conns[ws] = true
	s.mu.Unlock()

	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {

	defer func() {
		ws.Close()
		s.mu.Lock()
		delete(s.conns, ws)
		s.mu.Unlock()
	}()

	buf := make([]byte, 1024)
	for {

		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("error reading: %v", err)
			break
		}

		msg := buf[:n]
		cloneMsg := bytes.Clone(msg)
		s.broadcoast(cloneMsg)
	}
}

func (s *Server) broadcoast(b []byte) {

	s.mu.RLock()
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Printf("write error: %v\n", err)
			}
		}(ws)
	}
	s.mu.RUnlock()
}

func main() {
	server := NewServer()
	http.Handle("/ws", websocket.Handler(server.handleWS))
	http.ListenAndServe(":3000", nil)
}
