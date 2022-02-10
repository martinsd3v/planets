package logger

//ILoggerProvider for sign token struct
type ILoggerProvider interface {
	Error(v ...interface{})
	Info(v ...interface{})
}
