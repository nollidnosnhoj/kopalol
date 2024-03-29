package log

import (
	"log/slog"
	"os"
)

func LogErrorAndPanic(logger *slog.Logger, err error) {
	logger.Error(err.Error())
	os.Exit(1)
}
