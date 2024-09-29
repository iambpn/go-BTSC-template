package logger

import (
	"io"
	"log/slog"

	slogmulti "github.com/samber/slog-multi"
)

var Log *slog.Logger

func SetupLogger(writer io.Writer) {
	loggerHandler := []slog.Handler{
		slog.Default().Handler(),
	}

	if writer != nil {
		loggerHandler = append(loggerHandler, slog.NewJSONHandler(writer, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelInfo,
		}))
	}

	Log = slog.New(
		slogmulti.Fanout(
			loggerHandler...,
		),
	)
}
