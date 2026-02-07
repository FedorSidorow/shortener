package logger

import (
	"time"

	"go.uber.org/zap"
)

var Log *zap.Logger = zap.NewNop()

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

func DurationField(s string, duration time.Duration) zap.Field {
	return zap.Duration(s, duration)
}

func StringField(s1, s2 string) zap.Field {
	return zap.String(s1, s2)
}

func IntField(s string, val int) zap.Field {
	return zap.Int(s, val)
}

func ErrorField(err error) zap.Field {
	return zap.NamedError("error", err)
}
