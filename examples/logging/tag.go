package logging

const (
	ErrorLogKey = "ERROR"
	FatalLogKey = "FATAL"
	InfoLogKey  = "INFO"
	PanicLogKey = "PANIC"
	WarnLogKey  = "WARN"
)

type LogTag struct {
	name   string
	prefix string
}

func (l *LogTag) Name() string {
	return l.name
}

func (l *LogTag) Prefix() string {
	return l.prefix
}

var logTagMap = map[string]LogTag{
	ErrorLogKey: {name: ErrorLogKey, prefix: "[" + ErrorLogKey + "]"},
	FatalLogKey: {name: FatalLogKey, prefix: "[" + FatalLogKey + "]"},
	InfoLogKey:  {name: InfoLogKey, prefix: "[" + InfoLogKey + "]"},
	PanicLogKey: {name: PanicLogKey, prefix: "[" + PanicLogKey + "]"},
	WarnLogKey:  {name: WarnLogKey, prefix: "[" + WarnLogKey + "]"},
}

func getErrorLogTag() LogTag {
	return logTagMap[ErrorLogKey]
}

func getFatalLogTag() LogTag {
	return logTagMap[FatalLogKey]
}

func getInfoLogTag() LogTag {
	return logTagMap[InfoLogKey]
}

func getPanicLogTag() LogTag {
	return logTagMap[PanicLogKey]
}

func getWarnLogTag() LogTag {
	return logTagMap[WarnLogKey]
}
