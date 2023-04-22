package zap

import (
	"log"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const callerSkip = 1

type St struct {
	l  *zap.Logger
	sl *zap.SugaredLogger
}

func New(level string, dev bool) *St {
	var cfg zap.Config

	if dev {
		cfg = zap.NewDevelopmentConfig()
		cfg.Encoding = "console"
		//cfg.EncoderConfig.ConsoleSeparator = "\n\t\t"
	} else {
		cfg = zap.NewProductionConfig()
		cfg.Level.SetLevel(getZapLevel(level))
	}

	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	l, err := cfg.Build(zap.AddCallerSkip(callerSkip))
	if err != nil {
		log.Fatal(err)
	}

	return &St{
		l:  l,
		sl: l.Sugar(),
	}
}

func getZapLevel(v string) zapcore.Level {
	switch strings.ToLower(v) {
	case "error":
		return zap.ErrorLevel
	case "warn":
		return zap.WarnLevel
	case "info":
		return zap.InfoLevel
	case "debug":
		return zap.DebugLevel
	default:
		return zap.InfoLevel
	}
}

func (o *St) Fatal(args ...any) {
	o.sl.Fatal(args...)
}

func (o *St) Fatalf(tmpl string, args ...any) {
	o.sl.Fatalf(tmpl, args...)
}

func (o *St) Fatalw(msg string, err any, args ...any) {
	args = append(args, "error", err)
	o.sl.Fatalw(msg, args...)
}

func (o *St) Error(args ...any) {
	o.sl.Error(args...)
}

func (o *St) Errorf(tmpl string, args ...any) {
	o.sl.Errorf(tmpl, args...)
}

func (o *St) Errorw(msg string, err any, args ...any) {
	args = append(args, "error", err)
	o.sl.Errorw(msg, args...)
}

func (o *St) Warn(args ...any) {
	o.sl.Warn(args...)
}

func (o *St) Warnf(tmpl string, args ...any) {
	o.sl.Warnf(tmpl, args...)
}

func (o *St) Warnw(msg string, args ...any) {
	o.sl.Warnw(msg, args...)
}

func (o *St) Info(args ...any) {
	o.sl.Info(args...)
}

func (o *St) Infof(tmpl string, args ...any) {
	o.sl.Infof(tmpl, args...)
}

func (o *St) Infow(msg string, args ...any) {
	o.sl.Infow(msg, args...)
}

func (o *St) Debug(args ...any) {
	o.sl.Debug(args...)
}

func (o *St) Debugf(tmpl string, args ...any) {
	o.sl.Debugf(tmpl, args...)
}

func (o *St) Debugw(msg string, args ...any) {
	o.sl.Debugw(msg, args...)
}

func (o *St) Sync() {
	if err := o.sl.Sync(); err != nil {
		log.Println("Fail to sync zap-logger", err)
	}
}
