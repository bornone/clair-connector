package common

import (
	"fmt"
	"runtime"
	"strings"

	log "github.com/Sirupsen/logrus"
)

// LOG Function Usage:
// common.LOG(level logrus.LogLevel, message string, additional arguments)
// message can receive {0}, {1}, {2} to be replaced with the additional arguments
// e.g.
// common.LOG(log.ErrorLevel, "Here is a message with paraments {0}, {1}, {2}", "A", "B", 1000)
//
// output:
// ERRO[2016-08-31T14:29:56-04:00] main.main (main.go:22) | [Here is a message with paraments A, B, 100]
func LOG(l log.Level, msg string, args ...interface{}) {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	pc, file, line, _ := runtime.Caller(1)

	_funcName := runtime.FuncForPC(pc).Name()
	splitted := strings.Split(_funcName, "/")
	funcName := splitted[len(splitted)-1]

	splitted = strings.Split(file, "/")
	fileName := splitted[len(splitted)-1]

	fmtMsg := msg
	for n, arg := range args {
		_toBeReplaced := fmt.Sprintf("{%v}", n)
		_whatToReplace := fmt.Sprintf("%v", arg)
		fmtMsg = strings.Replace(fmtMsg, _toBeReplaced, _whatToReplace, 1)
	}

	fullLogMessage := fmt.Sprintf("%s (%s:%v) %s ", funcName, fileName, line, fmtMsg)

	// TODO: implement more levels
	// [ panic fatal error warning info debug ]

	if l == log.PanicLevel {
		log.Panic(fullLogMessage)
	} else if l == log.FatalLevel {
		log.Fatal(fullLogMessage)
	} else if l == log.ErrorLevel {
		log.Error(fullLogMessage)
	} else if l == log.WarnLevel {
		log.Warning(fullLogMessage)
	} else if l == log.InfoLevel {
		log.Info(fullLogMessage)
	} else if l == log.DebugLevel {
		log.Debug(fullLogMessage)
	}
}
