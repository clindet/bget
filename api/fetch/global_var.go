package fetch

import (
	"io"
	"io/ioutil"
	"os"

	cio "github.com/openbiox/ligo/io"
	clog "github.com/openbiox/ligo/log"
)

var log = clog.Logger

// SetLogStream set log output stream
func SetLogStream(quiet bool, saveLog bool, logPath string) {
	cio.CreateFileParDir(logPath)
	if quiet && !saveLog {
		log.SetOutput(ioutil.Discard)
	} else if quiet && saveLog {
		logCon, _ := cio.Open(logPath)
		log.SetOutput(logCon)
	} else if !quiet && saveLog {
		logCon, _ := cio.Open(logPath)
		log.SetOutput(io.MultiWriter(os.Stderr, logCon))
	} else {
		log.SetOutput(os.Stderr)
	}
}
