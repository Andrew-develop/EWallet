package logger

import (
	"log/slog"
	"os"

	"EWallet/internal/adapter/config"
	"github.com/natefinch/lumberjack"
	slogmulti "github.com/samber/slog-multi"
)

var logger *slog.Logger

func Set(config *config.App) {
	logger = slog.New(
		slog.NewTextHandler(os.Stderr, nil),
	)

	if config.Env == "production" {
		logRotate := &lumberjack.Logger{
			Filename:   "log/app.log",
			MaxSize:    100, // megabytes
			MaxBackups: 3,
			MaxAge:     28, // days
			Compress:   true,
		}

		logger = slog.New(
			slogmulti.Fanout(
				slog.NewJSONHandler(logRotate, nil),
				slog.NewTextHandler(os.Stderr, nil),
			),
		)
	}

	slog.SetDefault(logger)
}
