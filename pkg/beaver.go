package pkg

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// TODO ALL OF THIS

// Beaver struct for handling logs
type Beaver struct {
	level    string
	filePath string
	file     *os.File
	logger   *slog.Logger
}

// Config struct to load settings from a file
type Config struct {
	Level    string `json:"log_level" yaml:"log_level"`
	FilePath string `json:"log_file" yaml:"log_file"`
}

func NewBeaver(filepath string) (*Beaver, error) {
	// if filepath is empty, use default path
	if filepath == "" {
		filepath = "./test.json"
	}

	// Open the log file for writing
	logFile, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	// Create a logger that writes to the file
	logger := slog.New(slog.NewJSONHandler(logFile, nil))

	return &Beaver{
		level:    `info`,
		filePath: filepath,
		file:     logFile,
		logger:   logger,
	}, nil
}

// NewBeaverFromFile loads Beaver configuration from a file and initializes the Beaver
func NewBeaverFromFile(filename string) (*Beaver, error) {
	// Read and parse the configuration file
	configFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	var config Config

	//check if ends in .yaml or .json
	switch filename[len(filename)-4:] {

	case "yaml": // unable to open yaml file
		// if filename is .yaml, dont assume its yaml
		if err := yaml.NewDecoder(configFile).Decode(&config); err != nil {
			return nil, err
		}

	case "json":
		// if filename is .json, assume it's JSON if not yaml
		if err := json.NewDecoder(configFile).Decode(&config); err != nil {
			return nil, err
		}
	}

	// Open the log file for writing
	logFile, err := os.OpenFile(config.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	// Create a logger that writes to the file
	logger := slog.New(slog.NewJSONHandler(logFile, nil))

	return &Beaver{
		level:    config.Level,
		filePath: config.FilePath,
		file:     logFile,
		logger:   logger,
	}, nil
}

// Log method writes messages to the log.json file
func (l *Beaver) Log(message string) {
	switch l.level {
	case "debug":
		l.logger.Debug(message)
	case "info":
		l.logger.Info(message)
	case "warn":
		l.logger.Warn(message)
	case "error":
		l.logger.Error(message)
	}
}

func (b *Beaver) Debug(message string) {
	b.logger.Debug(message)
}

func (b *Beaver) Info(message string) {
	b.logger.Info(message)
}

func (b *Beaver) Error(message string) {
	b.logger.Error(message)
}

func (b *Beaver) Warn(message string) {
	b.logger.Warn(message)
}

func (b *Beaver) Close() {
	b.file.Close()
}

func (b *Beaver) GetLevel() string {
	return b.level
}

func (b *Beaver) GetFilePath() string {
	return b.filePath
}

func LoggingMiddleware(beaver *Beaver, next http.Handler) http.Handler {
	defer beaver.Close()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logMessage := "Method: " + r.Method + " | Path: " + r.URL.Path + " | Duration: " + time.Since(start).String()
		beaver.Log(logMessage)
	})
}
