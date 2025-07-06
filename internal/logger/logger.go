package logger

import (
	"os"

	logg "github.com/sirupsen/logrus"
)

var Logger *logg.Logger

func InitLogger(config *LoggerConfig) error {
	Logger = logg.New()
	Logger.SetFormatter(&logg.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	Logger.SetOutput(os.Stdout)
	Logger.SetReportCaller(config.ReportCaller)
	Logger.Info("Logger initialized successfully")
	return nil
}

func NewContextLogger(service string) *logg.Entry {
	return Logger.WithFields(logg.Fields{
		"service": service,
	})
}

func NewRequestLogger(method, path, requestId string) *logg.Entry {
	return Logger.WithFields(logg.Fields{
		"method":     method,
		"path":       path,
		"request_id": requestId,
		"component":  "http",
	})
}

func NewRepositoryLogger(table string) *logg.Entry {
	return Logger.WithFields(logg.Fields{
		"table":     table,
		"component": "repository",
	})
}

func NewLogicLogger(operation string) *logg.Entry {
	return Logger.WithFields(logg.Fields{
		"operation": operation,
		"component": "logic",
	})
}
