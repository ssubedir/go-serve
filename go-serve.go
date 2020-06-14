package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Flags struct {
	tls    bool
	port   string
	path   string
	readT  int
	writeT int
	idleT  int
}

func main() {

	fs := &Flags{}

	// default http server config flags

	flag.StringVar(&fs.path, "path", "www", "Files Path")
	flag.BoolVar(&fs.tls, "tls", false, "TLS")
	flag.StringVar(&fs.port, "port", "9000", "Server Port")
	flag.IntVar(&fs.readT, "rt", 5, "Read Timeout")
	flag.IntVar(&fs.writeT, "wt", 10, "Write Timeout")
	flag.IntVar(&fs.idleT, "it", 120, "Idle Timeout")

	flag.Parse()

	s := ServerConfig(fs)

	go func() {
		log.Println("Starting http file server on port " + fs.port)
		err := s.ListenAndServe()
		if err != nil {
			log.Panic("Error starting server")
		}
	}()

	sigs := make(chan os.Signal, 1)
	sigdone := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Signals
	go func() {
		sig := <-sigs
		fmt.Println(sig)
		sigdone <- true
	}()

	<-sigdone
	log.Println("Exiting Server")
	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	s.Shutdown(ctx)
}

func ServerConfig(fs *Flags) *http.Server {

	s := &http.Server{
		Addr:         ":" + fs.port,                          // configure the bind address
		Handler:      http.FileServer(http.Dir(fs.path)),     // set FileServer handler
		ReadTimeout:  time.Duration(fs.readT) * time.Second,  // max time to read request from the client
		WriteTimeout: time.Duration(fs.writeT) * time.Second, // max time to write response to the client
		IdleTimeout:  time.Duration(fs.idleT) * time.Second,  // max time for connections using TCP Keep-Alive
	}
	return s
}
