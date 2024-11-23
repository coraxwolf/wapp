package config_test

import (
	"fmt"
	"testing"

	"github.com/coraxwolf/wapp/internal/config"
)

/*
Should Give Error if ANY ENV Variable is not set
*/
func TestConfig_MissingENV(t *testing.T) {

	// Blank all ENV Variables
	t.Setenv("ADDR", "")
	t.Setenv("STATIC_PATH", "")
	t.Setenv("HTTP_TIMEOUT", "")
	t.Setenv("HTTP_SECRET", "")
	t.Setenv("HTTP_LOG", "")
	t.Setenv("HTTP_LOG_NAME", "")
	t.Setenv("OW_TOKEN", "")
	t.Setenv("LOG_LEVEL", "")
	t.Setenv("LOG_PATH", "")
	t.Setenv("LOG_FILENAME", "")
	t.Setenv("LOG_CONSOLE", "")
	t.Setenv("LOG_FMT_JSON", "")
	t.Setenv("DB_HOST", "")
	t.Setenv("DB_PORT", "")
	t.Setenv("DB_USER", "")
	t.Setenv("DB_PASS", "")
	t.Setenv("DB_NAME", "")
	t.Setenv("DATA_PATH", "")

	_, err := config.New() // Generate new Config Object
	if err == nil || err.Error() != config.ErrMissingENVVars.Error() {
		t.Errorf("expected error of ErrMissingENV, but got %v", err)
	}
}

/*
Should set the WeatherAPIToken value to API Key string (sample: ABCDEF1234)
*/
func TestConfig_SetAPIToken(t *testing.T) {
	t.Setenv("ADDR", ":8888")
	t.Setenv("STATIC_PATH", "./")
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
	t.Setenv("DATA_PATH", "./")

	fmt.Println("Updated ENV Variables")

	config, err := config.New()
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
	if config.WeatherAPIToken != "ABCDEF1234" {
		t.Errorf("expected WeatherAPIToken to be ABCDEF1234, but got %v", config.WeatherAPIToken)
	}
}
