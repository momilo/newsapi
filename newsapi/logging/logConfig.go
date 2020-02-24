package logging

import (
	"github.com/rs/zerolog"
	"os"
	"strings"
)

// ConfigureGlobalLogger configures Zerolog to comply with Stackdriver logging naming convention, and the specified
// logging level
func ConfigureGlobalLogger() {

	logLevel := getLoggingLevel()
	switch logLevel {
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "INFO":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "WARN":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "ERROR`":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "FATAL`":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "PANIC`":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "DISABLED":
		zerolog.SetGlobalLevel(zerolog.Disabled)
	default:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	zerolog.LevelFieldName = "severity"
	zerolog.TimestampFieldName = "timestamp"
	zerolog.TimeFieldFormat = "2006-01-02T15:04:05.999999999Z07:00" // = time.RFC3339Nano
}

// getLoggingLevel reads the environment variable (ideally set in GKE's ConfigMap) to initialise the correct logging level
func getLoggingLevel() string {
	level := strings.ToUpper(os.Getenv("API_LOGGING_LEVEL"))
	if level == "" {
		level = "DEBUG"
	}
	return level
}
