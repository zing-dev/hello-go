package logging

type LogManager struct {
	loggers []Logger
}

func (l *LogManager) GetPosition() Position {
	return PositionSingle
}

func (l *LogManager) SetPosition(pos Position) {}

func (l *LogManager) Error(v ...interface{}) string {
	var content string
	for _, logger := range l.loggers {
		content = logger.Error(v...)
	}
	return content
}

func (l *LogManager) Errorf(format string, v ...interface{}) string {
	var content string
	for _, logger := range l.loggers {
		content = logger.Errorf(format, v...)
	}
	return content
}

func (l *LogManager) Errorln(v ...interface{}) string {
	var content string
	for _, logger := range l.loggers {
		content = logger.Errorln(v...)
	}
	return content
}

func (l *LogManager) Fatal(v ...interface{}) string {
	var content string
	for _, logger := range l.loggers {
		content = logger.Fatal(v...)
	}
	return content
}

func (l *LogManager) Fatalf(format string, v ...interface{}) string {
	var content string
	for _, logger := range l.loggers {
		content = logger.Fatalf(format, v...)
	}
	return content
}

func (l *LogManager) Fatalln(v ...interface{}) string {
	var content string
	for _, logger := range l.loggers {
		content = logger.Fatalln(v...)
	}
	return content
}

func (l *LogManager) Info(v ...interface{}) string {
	var content string
	for _, logger := range l.loggers {
		content = logger.Info(v...)
	}
	return content
}

func (l *LogManager) Infof(format string, v ...interface{}) string {
	var content string
	for _, logger := range l.loggers {
		content = logger.Infof(format, v...)
	}
	return content
}

func (l *LogManager) Infoln(v ...interface{}) string {
	var content string
	for _, logger := range l.loggers {
		content = logger.Infoln(v...)
	}
	return content
}

func (l *LogManager) Panic(v ...interface{}) string {
	var content string
	for _, logger := range l.loggers {
		content = logger.Panic(v...)
	}
	return content
}

func (l *LogManager) Panicf(format string, v ...interface{}) string {
	var content string
	for _, logger := range l.loggers {
		content = logger.Panicf(format, v...)
	}
	return content
}

func (l *LogManager) Panicln(v ...interface{}) string {
	var content string
	for _, logger := range l.loggers {
		content = logger.Panicln(v...)
	}
	return content
}

func (l *LogManager) Warn(v ...interface{}) string {
	var content string
	for _, logger := range l.loggers {
		content = logger.Warn(v...)
	}
	return content
}

func (l *LogManager) Warnf(format string, v ...interface{}) string {
	var content string
	for _, logger := range l.loggers {
		content = logger.Warnf(format, v...)
	}
	return content
}

func (l *LogManager) Warnln(v ...interface{}) string {
	var content string
	for _, logger := range l.loggers {
		content = logger.Warnln(v...)
	}
	return content
}
