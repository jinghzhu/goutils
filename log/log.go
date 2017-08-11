package log

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	DEBUG   = "DEBUG"
	INFO    = "INFO"
	WARNING = "WARNING"
	ERROR   = "ERROR"
)

var (
	logLevels map[string]int
	logLevel  = 0
	channels  []*log.Logger
)

func init() {
	logLevels = map[string]int{ERROR: 3, WARNING: 2, INFO: 1, DEBUG: 0}
	console := GetConsoleLogger()
	conn := GetConnLogger()
	channels = []*log.Logger{conn, console}
}

func SetLogLevel(level string) {
	if !validateLogLevel(level) {
		fmt.Println("Log level(" + level + ") is not correct.")
		return
	}
	logLevel = logLevels[strings.ToUpper(level)]
}

func Debug(msg string) {
	if logLevels[DEBUG] >= logLevel {
		printLog(DEBUG, msg)
	}
}

func Info(msg string) {
	if logLevels[INFO] >= logLevel {
		printLog(INFO, msg)
	}
}

func Warning(msg string) {
	if logLevels[WARNING] >= logLevel {
		printLog(WARNING, msg)
	}
}

func Error(msg string) {
	if logLevels[ERROR] >= logLevel {
		printLog(ERROR, msg)
	}
}

func printLog(prefix string, msg string) {
	filename, line := Locate()
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Hostname error:" + err.Error())
		hostname = "Error Hostname"
	}
	for _, v := range channels {
		v.Println("[" + prefix + "] " + hostname + " " + filename + " ln" + fmt.Sprintf("%d", line) + " " + msg)
	}
}

func validateLogLevel(level string) bool {
	level = strings.ToUpper(level)
	if logLevels == nil {
		Error("logLevels is nil")
		return false
	}
	for k := range logLevels {
		if k == level {
			return true
		}
	}
	return false
}

func Locate() (filename string, line int) {
	_, path, line, ok := runtime.Caller(3)
	file := ""
	if ok {
		_, file = filepath.Split(path)
	} else {
		fmt.Println("Fail to get method caller")
		line = -1
	}
	return file, line
}
