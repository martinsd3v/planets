package logger

import log "github.com/amoghe/distillog"

//Logger struct for assign interface
type Logger struct{}

//New create a new instance
func New() *Logger {
	return &Logger{}
}

//Token implements the TokenInterface
var _ ILoggerProvider = &Logger{}

//Error ...
func (l *Logger) Error(v ...interface{}) {
	log.Errorln(v...)
}

//Info ...
func (l *Logger) Info(v ...interface{}) {
	log.Infoln(v...)
}
