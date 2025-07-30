package appcore_logger

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

// InitLogger returns a new logger instance
func InitLogger() {
	Logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	Logger.Info("Init Logger")
}
