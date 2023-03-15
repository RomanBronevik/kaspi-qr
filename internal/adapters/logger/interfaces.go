package logger

type Full interface {
	Debug(args ...any)
	Debugf(tmpl string, args ...any)
	Debugw(msg string, args ...any)

	Info(args ...any)
	Infof(tmpl string, args ...any)
	Infow(msg string, args ...any)

	Warn(args ...any)
	Warnf(tmpl string, args ...any)
	Warnw(msg string, args ...any)

	Error(args ...any)
	Errorf(tmpl string, args ...any)
	Errorw(msg string, err any, args ...any)

	Fatal(args ...any)
	Fatalf(tmpl string, args ...any)
	Fatalw(msg string, err any, args ...any)
}

type Lite interface {
	Infow(msg string, args ...any)
	Warnw(msg string, args ...any)
	Errorw(msg string, err any, args ...any)
}

type WarnAndError interface {
	Warnw(msg string, args ...any)
	Errorw(msg string, err any, args ...any)
}
