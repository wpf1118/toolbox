// Package logging is to define our logger
package logging

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

var (
	log zerolog.Logger
)

func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log = zerolog.New(defaultZerologConsoleWriter()).With().Timestamp().CallerWithSkipFrameCount(3).Logger()
}

func defaultZerologConsoleWriter() zerolog.ConsoleWriter {
	return zerolog.ConsoleWriter{
		Out:        os.Stdout,
		NoColor:    true,
		TimeFormat: time.RFC3339,
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("%s:", i))
		},
	}
}

// SetVerbose is to enable the zerolog to debug level.
func SetVerbose() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
}

// Disable is to disable the loggin altogether.
func Disable() {
	zerolog.SetGlobalLevel(zerolog.Disabled)
}

func DebugF(format string, v ...interface{}) {
	log.Debug().Msgf(format, v...)
}

func InfoF(format string, v ...interface{}) {
	log.Info().Msgf(format, v...)
}

func ErrorF(format string, v ...interface{}) error {
	log.Error().Msgf(format, v...)
	return fmt.Errorf(format, v...)
}
