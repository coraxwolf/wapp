package config

import (
	"errors"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTP            HTTPConfig
	DB              DBConfig
	Logging         LogConfig
	WeatherAPIToken string
}

type DBConfig struct {
	Dsn string
}

type LogConfig struct {
	Level   string
	Path    string
	File    string
	Console bool
	JSON    bool
}

type HTTPConfig struct {
	ADDR       string
	Mux        *http.ServeMux
	Timeout    time.Duration
	Secret     string
	Log        bool
	LogFile    string
	StaticPath string
}

// Errors
var (
	ErrMissingENVVars = errors.New("missing required ENV variables")
	ErrLoadingENV     = errors.New("error loading .env file")
)

func New() (*Config, error) {
	// Check for testing environment
	if os.Getenv("ENV") == "testing" {
		slog.Debug("Running in testing environment")
	} else {
		err := godotenv.Load()
		if err != nil {
			slog.Error("error loading .env file", "error", err)
			return nil, ErrLoadingENV
		}
	}
	addr := os.Getenv("ADDR")
	spath := os.Getenv("STATIC_PATH")
	httplog := os.Getenv("HTTP_LOG") == "true"
	httplname := os.Getenv("HTTP_LOG_NAME")
	httptm, err := strconv.Atoi(os.Getenv("HTTP_TIMEOUT"))
	if err != nil {
		httptm = 10
	}
	secret := os.Getenv("HTTP_SECRET")
	owToken := os.Getenv("OW_TOKEN")
	logLevel := os.Getenv("LOG_LEVEL")
	logPath := os.Getenv("LOG_PATH")
	logFile := os.Getenv("LOG_FILENAME")
	logConsole := os.Getenv("LOG_CONSOLE") == "true"
	logJSON := os.Getenv("LOG_FMT_JSON") == "true"
	dataPath := os.Getenv("DATA_PATH")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	// Verify all vars are set
	if addr == "" || spath == "" ||
		secret == "" || httptm == 0 ||
		owToken == "" || logLevel == "" ||
		logPath == "" || logFile == "" ||
		dataPath == "" || dbHost == "" ||
		dbPort == "" || dbUser == "" ||
		dbPass == "" || dbName == "" {
		return nil, ErrMissingENVVars
	}

	// Setup Configs
	ht := &HTTPConfig{
		ADDR:       addr,
		Timeout:    time.Duration(httptm) * time.Second,
		Secret:     secret,
		Log:        httplog,
		LogFile:    dataPath + httplname,
		StaticPath: spath,
	}

	db := &DBConfig{
		Dsn: dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName,
	}

	lc := &LogConfig{
		Level:   logLevel,
		Path:    dataPath + logPath,
		File:    logFile,
		Console: logConsole,
		JSON:    logJSON,
	}

	c := &Config{
		WeatherAPIToken: owToken,
		HTTP:            *ht,
		DB:              *db,
		Logging:         *lc,
	}
	return c, nil
}
