package logger

import (
	"time"

	"go.uber.org/zap"
)

// По умолчанию установлен no-op-логер, который не выводит никаких сообщений.
var Log *zap.Logger = zap.NewNop()

// Initialize инициализирует синглтон логера с необходимым уровнем логирования.
func Initialize(level string) error {

	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return err
	}

	cfg := zap.NewProductionConfig()
	cfg.Level = lvl

	zl, err := cfg.Build()
	if err != nil {
		return err
	}

	defer zl.Sync()
	Log = zl

	return nil
}

// DurationField создает поле с ключем и длительностью.
func DurationField(s string, duration time.Duration) zap.Field {
	return zap.Duration(s, duration)
}

// StringField создает поле с клюючем и полем типа строка.
func StringField(s1, s2 string) zap.Field {
	return zap.String(s1, s2)
}

// IntField создает поле с клюючем и полем типа число.
func IntField(s string, val int) zap.Field {
	return zap.Int(s, val)
}

// IntField создает поле с клюючем и полем ошибки.
func ErrorField(err error) zap.Field {
	return zap.NamedError("error", err)
}
