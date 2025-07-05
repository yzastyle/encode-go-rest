package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/yzastyle/encode-go-rest/internal/config"
)

var Logger log.Logger

func InitLogger(config *config.LoggerConfig) error {
	Logger := log.New()

	Logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	Logger.SetOutput(os.Stdout)
	Logger.SetReportCaller(config.ReportCaller)

	Logger.Info("Logger initialized successfully")
	return nil
}
