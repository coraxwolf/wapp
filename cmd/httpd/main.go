package web

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/coraxwolf/wapp/internal/config"
)

type WebServer struct {
	Addr    string
	Secret  string
	wait    sync.WaitGroup
	runChan chan RunChan
	Mux     *http.ServeMux
	Srv     *http.Server
	Timeout time.Duration
}

type RunChan struct {
	Error  error
	Status string
}

// Errors
var (
	ErrInvalidConfig     = errors.New("invalid configuration")
	ErrWebServerStopped  = errors.New("web server stopped")
	ErrWebServerShutdown = errors.New("error shutting down web server")
)

func NewServer(cfg config.HTTPConfig) (*WebServer, error) {
	if cfg.ADDR == "" || cfg.Secret == "" {
		return nil, ErrInvalidConfig
	}
	m := http.NewServeMux()
	s := &http.Server{
		Addr:              cfg.ADDR,
		Handler:           m,
		ReadHeaderTimeout: cfg.Timeout,
	}

	// Set Static File Route
	fmt.Printf("Static Path: %s\n", cfg.StaticPath)
	m.Handle("GET /assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(cfg.StaticPath))))

	ws := &WebServer{
		Addr:    cfg.ADDR,
		Secret:  cfg.Secret,
		wait:    sync.WaitGroup{},
		runChan: make(chan RunChan),
		Mux:     m,
		Srv:     s,
		Timeout: cfg.Timeout,
	}

	return ws, nil
}

/* Run Web Server in Go Routine */
func (ws *WebServer) run() error {
	ws.wait.Add(1)
	defer ws.wait.Done()
	err := ws.Srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		ws.runChan <- RunChan{Error: err, Status: "error"}
		return ErrWebServerStopped
	}
	ws.runChan <- RunChan{Error: nil, Status: "closing"}
	return nil
}

/* Start Web Server */
func (ws *WebServer) Start() error {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	go ws.run()
	switch <-sig {
	case os.Interrupt:
		ws.Stop()
	case syscall.SIGTERM:
		ws.Stop()
	default:
		slog.Warn("unexpected signal received")
	}
	return nil
}

/* Stop Web Server Gracefully */
func (ws *WebServer) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := ws.Srv.Shutdown(ctx); err != nil {
		ws.runChan <- RunChan{Error: err, Status: "error"}
		return ErrWebServerShutdown
	}
	ws.runChan <- RunChan{Error: nil, Status: "stopped"}
	return nil
}
