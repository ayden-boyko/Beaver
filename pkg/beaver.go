package pkg

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

// TODO ALL OF THIS

// Logger struct for handling logs
type Logger struct {
	level    string
	filePath string
	file     *os.File
	logger   *log.Logger
}

// Config struct to load settings from a file
type Config struct {
	Level    string `json:"log_level"`
	FilePath string `json:"log_file"`
}

// Log method writes messages to the log file
func (l *Logger) Log(message string) {
	l.logger.Println(message)
}

// NewLoggerFromFile loads logger configuration from a file and initializes the logger
func NewLoggerFromFile(filename string) (*Logger, error) {
	// Read and parse the configuration file
	configFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	var config Config
	if err := json.NewDecoder(configFile).Decode(&config); err != nil {
		return nil, err
	}

	// Open the log file for writing
	logFile, err := os.OpenFile(config.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	// Create a logger that writes to the file
	logger := log.New(logFile, "", log.LstdFlags)

	return &Logger{
		level:    config.Level,
		filePath: config.FilePath,
		file:     logFile,
		logger:   logger,
	}, nil
}

func LoggingMiddleware(logger *Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logMessage := "Method: " + r.Method + " | Path: " + r.URL.Path + " | Duration: " + time.Since(start).String()
		logger.Log(logMessage)
	})
}
