package main

import (
	"context"
	logs "firstProject/internal/log"
	"github.com/go-chi/chi"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type MyServer struct {
	server *http.Server
}

func NewServer(mux *chi.Mux) *MyServer {
	s := &http.Server{
		Addr:           ":9000",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &MyServer{s}
}

// Run: runs ListenAndServe on the http.Server with graceful shutdown
func (s *MyServer) Run() {
	logs.Log().Info("starting server...")

	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logs.Sugar().Fatalf("could not listen on %s due to %s", s.server.Addr, err.Error())
		}
	}()
	logs.Sugar().Infof("server is ready to handle requests %s", s.server.Addr)
	s.gracefulShutdown()
}

// wait for an interrupt signal
func (s *MyServer) gracefulShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	logs.Sugar().Infof("server is shutting down %s", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s.server.SetKeepAlivesEnabled(false)
	if err := s.server.Shutdown(ctx); err != nil {
		logs.Sugar().Fatalf("could not gracefully shutdown the server %s", err.Error())
	}
	logs.Log().Info("server stopped")
}
