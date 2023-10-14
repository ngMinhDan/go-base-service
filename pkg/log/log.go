/*
Package log: Provides methods for working with logs.
Package Functionality: Enables the creation of global logs with customizable log formats and output options.
*/

package log

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"base/pkg/config"
	"base/pkg/constant"
)

// Log Variable
var logger *logrus.Logger

// Log Level Data Type
type logLevel string

// Log Level Data Type Constant
const (
	LogLevelPanic logLevel = "panic"
	LogLevelFatal logLevel = "fatal"
	LogLevelError logLevel = "error"
	LogLevelWarn  logLevel = "warn"
	LogLevelDebug logLevel = "debug"
	LogLevelTrace logLevel = "trace"
	LogLevelInfo  logLevel = "info"
)

// Initialize Function in Helper Logging
func init() {
	// Initialize Log as New Logrus Logger
	logger = logrus.New()

	if strings.ToLower(config.Config.GetString("LOG_FORMAT")) == constant.JsonFormat {
		// Set Log Format to JSON Format
		logger.SetFormatter(&logrus.JSONFormatter{
			DisableTimestamp: false,
			TimestampFormat:  time.RFC3339Nano,
		})
	} else {
		// Set Log Format to TEXT Format
		logger.SetFormatter(&logrus.TextFormatter{
			DisableTimestamp: false,
			TimestampFormat:  time.RFC3339Nano,
			DisableColors:    false, // will see color in console
		})
	}

	filePath := strings.ToLower(config.Config.GetString("LOG_OUTPUT"))

	if filePath != constant.ConsoleOuput {
		// Set Log Output to File
		file, error := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if error != nil {
			log.Fatalln(error)
		}
		logger.SetOutput(file)
	} else {
		// Dedault: Set Log Output to STDOUT
		logger.SetOutput(os.Stdout)
	}

	// Set Log Level
	switch strings.ToLower(config.Config.GetString("SERVER_LOG_LEVEL")) {
	case "panic":
		logger.SetLevel(logrus.PanicLevel)
	case "fatal":
		logger.SetLevel(logrus.FatalLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "trace":
		logger.SetLevel(logrus.TraceLevel)
	default:
		logger.SetLevel(logrus.InfoLevel)
	}
}

// Println Function for Log
func Println(level logLevel, label string, message interface{}) {
	// Make Sure Log Is Not Empty Variable
	if logger != nil {
		// Set Service Name Log Information
		service := strings.ToLower(config.Config.GetString("SERVER_NAME"))

		// Print Log Based On Log Level Type
		switch level {
		case "panic":
			logger.WithFields(logrus.Fields{
				"service": service,
				"label":   label,
			}).Panicln(message)
		case "fatal":
			logger.WithFields(logrus.Fields{
				"service": service,
				"label":   label,
			}).Fatalln(message)
		case "error":
			logger.WithFields(logrus.Fields{
				"service": service,
				"label":   label,
			}).Errorln(message)
		case "warn":
			logger.WithFields(logrus.Fields{
				"service": service,
				"label":   label,
			}).Warnln(message)
		case "debug":
			logger.WithFields(logrus.Fields{
				"service": service,
				"label":   label,
			}).Debug(message)
		case "trace":
			logger.WithFields(logrus.Fields{
				"service": service,
				"label":   label,
			}).Traceln(message)
		default:
			logger.WithFields(logrus.Fields{
				"service": service,
				"label":   label,
			}).Infoln(message)
		}
	}
}
