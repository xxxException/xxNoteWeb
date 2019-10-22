package logger

import (
	log "github.com/sirupsen/logrus"
	"os"
)

var Logger *log.Logger

func NewLogger() {
	Logger = log.New()
	Logger.Out = os.Stdout
}
