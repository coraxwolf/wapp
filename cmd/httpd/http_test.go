package web_test

import (
	"syscall"
	"testing"
	"time"

	web "github.com/coraxwolf/wapp/cmd/httpd"
	"github.com/coraxwolf/wapp/internal/config"
)

/* Should create a Web Server for responding to requests */
func Test_Create_WebServer(t *testing.T) {
	t.Setenv("ADDR", ":8888")
	t.Setenv("STATIC_PATH", "../data/assets")
	t.Setenv("HTTP_TIMEOUT", "10")
	t.Setenv("HTTP_SECRET", "secret")
	t.Setenv("HTTP_LOG", "true")
	t.Setenv("HTTP_LOG_NAME", "http.log")
	t.Setenv("OW_TOKEN", "ABCDEF1234") // Ensure ENV is set to the sample value
	t.Setenv("LOG_LEVEL", "debug")
	t.Setenv("LOG_PATH", "./")
	t.Setenv("LOG_FILENAME", "log")
	t.Setenv("LOG_CONSOLE", "true")
	t.Setenv("LOG_FMT_JSON", "true")
	t.Setenv("DB_HOST", "localhost")
	t.Setenv("DB_PORT", "3306")
	t.Setenv("DB_USER", "user")
	t.Setenv("DB_PASS", "secret")
	t.Setenv("DB_NAME", "db")
	t.Setenv("DATA_PATH", "../data")
	cfg, err := config.New()
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
	if cfg == nil {
		t.Errorf("expected config object, but got nil")
	}
	// Create Web Server
	ws, err := web.NewServer(cfg.HTTP)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
	if ws == nil {
		t.Errorf("expected web server object, but got nil")
	}
}

/* Web Server Should Start and Stop Gracefully */
func Test_Start_Stop_WebServer(t *testing.T) {
	t.Setenv("ADDR", ":8888")
	t.Setenv("STATIC_PATH", "../data/assets")
	t.Setenv("HTTP_TIMEOUT", "10")
	t.Setenv("HTTP_SECRET", "secret")
	t.Setenv("HTTP_LOG", "true")
	t.Setenv("HTTP_LOG_NAME", "http.log")
	t.Setenv("OW_TOKEN", "ABCDEF1234") // Ensure ENV is set to the sample value
	t.Setenv("LOG_LEVEL", "debug")
	t.Setenv("LOG_PATH", "./")
	t.Setenv("LOG_FILENAME", "log")
	t.Setenv("LOG_CONSOLE", "true")
	t.Setenv("LOG_FMT_JSON", "true")
	t.Setenv("DB_HOST", "localhost")
	t.Setenv("DB_PORT", "3306")
	t.Setenv("DB_USER", "user")
	t.Setenv("DB_PASS", "secret")
	t.Setenv("DB_NAME", "db")
	t.Setenv("DATA_PATH", "../data")
	cfg, err := config.New()
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
	if cfg == nil {
		t.Errorf("expected config object, but got nil")
	}
	// Create Web Server
	ws, err := web.NewServer(cfg.HTTP)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
	if ws == nil {
		t.Errorf("expected web server object, but got nil")
	}
	go func() {
		if err := ws.Start(); err != nil {
			t.Errorf("expected no error, but got %v", err)
		}
	}()
	// Wait 2 seconds then send interrupt signal
	time.Sleep(2 * time.Second)
	syscall.Kill(syscall.Getpid(), syscall.SIGINT)
}

/* Web Server Should Server Static Files from Static Directory */
func Test_Serve_Static_Files(t *testing.T) {
	t.Skip("Not implemented")
}

/* Web Server Should Return 404 for Non-Existent Routes */
func Test_404_Route(t *testing.T) {
	t.Skip("Not implemented")
}

/* Web Server Should Return 500 for Internal Server Errors */
func Test_500_Route(t *testing.T) {
	t.Skip("Not implemented")
}

/* Web Server Should Return 200 for Health Check */
func Test_Health_Check(t *testing.T) {
	t.Skip("Not implemented")
}
