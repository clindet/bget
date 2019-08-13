package log

import (
	"fmt"
	"io"
	"path/filepath"
	"runtime"
	"strings"

	nested "github.com/antonfisher/nested-logrus-formatter"
	logrus "github.com/sirupsen/logrus"
)

// Lg is the tai log object
var Lg = logrus.New()

// LogBash show [BASH] prefix in logrus message
var LogBash = Lg.WithFields(logrus.Fields{
	"prefix": "BASH"})

// DEBUG show log message (debug level) with filename, line, and funcname
func DEBUG(formating string, args ...interface{}) {
	printTrace(logrus.DebugLevel, "debug", formating, args...)
}

// INFO show log message (info level) with filename, line, and funcname
func INFO(formating string, args ...interface{}) {
	printTrace(logrus.InfoLevel, "info", formating, args...)
}

// WARN show log message (info level) with filename, line, and funcname
func WARN(formating string, args ...interface{}) {
	printTrace(logrus.WarnLevel, "warn", formating, args...)
}

// ERROR show log message (error level) with filename, line, and funcname
func ERROR(formating string, args ...interface{}) {
	printTrace(logrus.ErrorLevel, "error", formating, args...)
}

// FATAL show log message (fatal level) with filename, line, and funcname
func FATAL(formating string, args ...interface{}) {
	printTrace(logrus.FatalLevel, "fatal", formating, args...)
}

func printTrace(level logrus.Level, mode string, formating string, args ...interface{}) {
	filename, line, funcname := "???", 0, "???"
	pc, filename, line, ok := runtime.Caller(2)
	// fmt.Println(reflect.TypeOf(pc), reflect.ValueOf(pc))
	if ok {
		funcname = runtime.FuncForPC(pc).Name()
		funcname = filepath.Ext(funcname)            // .foo
		funcname = strings.TrimPrefix(funcname, ".") // foo
	}
	if mode == "info" {
		Lg.Infof("%s:%d:%s %s", filename, line, funcname, fmt.Sprintf(formating, args...))
	} else if mode == "warn" {
		Lg.Warnf("%s:%d:%s %s", filename, line, funcname, fmt.Sprintf(formating, args...))
	} else if mode == "error" {
		Lg.Errorf("%s:%d:%s %s", filename, line, funcname, fmt.Sprintf(formating, args...))
	} else if mode == "debug" {
		Lg.Debugf("%s:%d:%s %s", filename, line, funcname, fmt.Sprintf(formating, args...))
	} else {
		Lg.Fatalf("%s:%d:%s %s", filename, line, funcname, fmt.Sprintf(formating, args...))
	}
}

// Infoln print info
func Infoln(args ...interface{}) {
	Lg.Infoln(args...)
}

// Infof print info
func Infof(format string, args ...interface{}) {
	Lg.Infof(format, args...)
}

// SetOutput set log output
func SetOutput(output io.Writer) {
	Lg.SetOutput(output)
}

// Fatal print fatal info
func Fatal(args ...interface{}) {
	Lg.Fatal(args...)
}

// Fatalf print fatal info
func Fatalf(format string, args ...interface{}) {
	Lg.Fatalf(format, args...)
}

// Fatalln print fatal info
func Fatalln(args ...interface{}) {
	Lg.Fatalln(args...)
}

// Warn print Warn info
func Warn(args ...interface{}) {
	Lg.Warn(args...)
}

// Warnf print Warn info
func Warnf(format string, args ...interface{}) {
	Lg.Warnf(format, args...)
}

// New is wrapper of logrus.New()
func New() *logrus.Logger {
	return logrus.New()
}

func init() {
	Lg.SetLevel(logrus.InfoLevel)
	Lg.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		TimestampFormat: "[2006-01-02 15:04:05]",
	})
}
