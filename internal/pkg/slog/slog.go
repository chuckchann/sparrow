package slog

func Debug(args ...interface{}) {
	logger.Debugln(args)
}

func Info(args ...interface{}) {
	logger.Infoln(args)
}

func Warn(args ...interface{}) {
	logger.Warnln(args)
}

func Error(args ...interface{}) {
	logger.Errorln(args)
}

func Panic(args ...interface{}) {
	logger.Panicln(args)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args)
}

func Infof(format string, args ...interface{}) {
	logger.Debugf(format, args)
}

func Warnf(format string, args ...interface{}) {
	logger.Debugf(format, args)
}

func Errorf(format string, args ...interface{}) {
	logger.Debugf(format, args)
}

func Panicf(format string, args ...interface{}) {
	logger.Debugf(format, args)
}

func recordHttpRequest() {

}
