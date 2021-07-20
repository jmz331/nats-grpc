package logger

import "log"

func Info(msg string, v ...interface{}) {
	log.Printf(msg+"\n", v...)
}

func Error(msg string, v ...interface{}) {
	log.Printf(msg+"\n", v...)
}
