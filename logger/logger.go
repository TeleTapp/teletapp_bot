package logger

import (
	"fmt"
	"io"

	zlogsentry "github.com/archdx/zerolog-sentry"
	"github.com/rs/zerolog"
)

type Config struct {
	SentryServiceName string `env:"SENTRY_SERVICE_NAME,required"`
	SentryDSN         string `env:"SENTRY_DSN,required"`
	SentryEnvironment string `env:"SENTRY_ENVIRONMENT,required" envDefault:"release"`
	SentryRelease     string `env:"SENTRY_RELEASE,required"`
}

var cfg Config

// Выполнить в main функции приложения, чтобы добавить конфиг для Sentry
func Init(c Config) {
	cfg = c
}

func NewLogger(name string) *zerolog.Logger {
	var w io.Writer
	cw := zerolog.NewConsoleWriter()
	cw.NoColor = false
	cw.TimeFormat = "[LOG] 2006/01/02 - 15:04:05"
	cw.FormatCaller = func(i interface{}) string {
		return fmt.Sprintf("\x1b[%dm[%v]\x1b[0m", 32, name)
	}
	cw.FieldsExclude = append(cw.FieldsExclude, "logger")
	w = cw

	// Sentry
	if cfg.SentryDSN != "" {
		s, _ := zlogsentry.New(cfg.SentryDSN,
			zlogsentry.WithEnvironment(cfg.SentryEnvironment),
			zlogsentry.WithRelease(cfg.SentryRelease),
			zlogsentry.WithServerName(cfg.SentryServiceName),
		)

		w = zerolog.MultiLevelWriter(cw, s)
	}

	l := zerolog.New(w).
		With().Timestamp().Str("logger", name).Logger()
	return &l
}
