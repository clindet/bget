package utils

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	log "github.com/JhuangLab/bget/log"
)

var stdout, stderr []byte
var errStdout, errStderr error
var logPub = log.New()

// RunExecCmd exec shell cmd using exec.Cmd object
func RunExecCmd(logPath string, cmd *exec.Cmd) bool {
	return RunExecCmdConsole(logPath, cmd, true, true)
}

// RunExecCmdConsole exec shell cmd using exec.Cmd object (can control wheature output to console)
func RunExecCmdConsole(logPath string, cmd *exec.Cmd, quiet bool, saveLog bool) bool {
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	logFn, err := ConnectFile(logPath)
	if err != nil {
		return false
	}

	logBash := log.LogBash
	var log = log.Lg

	cmdStr := ""
	if logPath != "/dev/null" {
		cmdStr = strings.Join(cmd.Args, " ") + " &>> " + logPath
	} else {
		cmdStr = strings.Join(cmd.Args, " ")
	}
	if quiet && saveLog {
		log.SetOutput(io.Writer(logFn))
	} else if quiet && !saveLog {
		log.SetOutput(ioutil.Discard)
	} else if !quiet && saveLog {
		log.SetOutput(io.MultiWriter(os.Stderr, logFn))
	}
	logBash.Info(cmdStr)
	var con *os.File
	if logPath != "" {
		con, err = ConnectFile(logPath)
		if err != nil {
			log.Fatalf("Can not open log file %s", logPath)
		}
	}
	err = cmd.Start()
	go func() {
		stdout, errStdout = copyAndCapture(os.Stdout, stdoutIn, quiet, con)
	}()
	go func() {
		stderr, errStderr = copyAndCapture(os.Stderr, stderrIn, quiet, con)
	}()
	cmd.Wait()
	con.Close()
	if err != nil {
		log.Warningf("cmd.Run() failed with %s\n", err)
		log.SetOutput(os.Stderr)
		return false
	}
	log.SetOutput(os.Stderr)
	return true
}

func copyAndCapture(w io.Writer, r io.Reader, quiet bool, con *os.File) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			if !quiet {
				os.Stderr.Write(d)
			}
			if con != nil {
				con.Write(d)
			}
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
}
