package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

const (
	EnvHost = "HOST"
	EnvPort = "PORT"
)

func serve(l net.Listener, errCh chan struct{}) {
	for {
		conn, err := l.Accept()
		if err != nil {
			close(errCh)
			return
		}
		conn.Close()
	}
}

func main() {
	// build address
	host := strings.TrimSpace(os.Getenv(EnvHost))
	port := strings.TrimSpace(os.Getenv(EnvPort))
	if host == "" {
		host = "0.0.0.0"
	}
	if port == "" {
		port = "3000"
	}
	addr := fmt.Sprintf("%s:%s", host, port)

	// Serve
	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer l.Close()

	log.Printf("listening at %s\n", addr)

	errCh := make(chan struct{})

	go serve(l, errCh)

	term := make(chan os.Signal, 1)

	signal.Notify(term, syscall.SIGTERM)
	signal.Notify(term, os.Interrupt)

	select {
	case <-term:
	case <-errCh:
	}

	log.Println("exited")
}
