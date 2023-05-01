package log

// Interface represents the API of both Logger and Entry.
type Interface interface {
	WithField(string, interface{}) *Entry
	WithError(error) *Entry
	WithoutPadding() *Entry
	Debug(string)
	Info(string)
	Warn(string)
	Error(string)
	Fatal(string)
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	Fatalf(string, ...interface{})
	ResetPadding()
	IncreasePadding()
	DecreasePadding()
}
