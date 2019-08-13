package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	log "github.com/JhuangLab/bioget/log"
	mpb "github.com/vbauerster/mpb/v4"
	"github.com/vbauerster/mpb/v4/decor"
)

var pg *mpb.Progress

// Wget use wget to download files
func Wget(url string, destFn string, taskID string, console bool) {
	args := []string{"-c", url, "-O", destFn}
	cmd := exec.Command("wget", args...)
	wd, _ := os.Getwd()
	logPath := path.Join(wd, "_log", fmt.Sprintf("%s_%s_wget.log", taskID, path.Base(destFn)))
	CreateFileParDir(logPath)
	RunExecCmdConsole(logPath, cmd, console)
}

// Curl use curl to download files
func Curl(url string, destFn string, taskID string, console bool) {
	args := []string{url, "-o", destFn}
	cmd := exec.Command("curl", args...)
	wd, _ := os.Getwd()
	logPath := path.Join(wd, "_log", fmt.Sprintf("%s_%s_curl.log", taskID, path.Base(destFn)))
	CreateFileParDir(logPath)
	RunExecCmdConsole(logPath, cmd, console)
}

// Axel use axel to download files
func Axel(url string, destFn string, thread int, taskID string, console bool) {
	args := []string{url, "-N", "-o", destFn, "-n", strconv.Itoa(thread)}
	cmd := exec.Command("axel", args...)
	wd, _ := os.Getwd()
	logPath := path.Join(wd, "_log", fmt.Sprintf("%s_%s_axel.log", taskID, path.Base(destFn)))
	CreateFileParDir(logPath)
	RunExecCmdConsole(logPath, cmd, console)
}

// Git use git to download files
func Git(url string, destFn string, taskID string, console bool) {
	args := []string{"clone", "--recursive", url, destFn}
	cmd := exec.Command("git", args...)
	wd, _ := os.Getwd()
	logPath := path.Join(wd, "_log", fmt.Sprintf("%s_%s_git.log", taskID, path.Base(destFn)))
	CreateFileParDir(logPath)
	RunExecCmdConsole(logPath, cmd, console)
}

// Rsync use rsync to download files
func Rsync(url string, destFn string, taskID string, console bool) {
	args := []string{url, destFn}
	cmd := exec.Command("rsync", args...)
	wd, _ := os.Getwd()
	logPath := path.Join(wd, "_log", fmt.Sprintf("%s_%s_rsync.log", taskID, path.Base(destFn)))
	CreateFileParDir(logPath)
	RunExecCmdConsole(logPath, cmd, console)
}

// httpGetURL can use golang http.Get to query URL with progress bar
func httpGetURL(url string, destFn string, p *mpb.Progress, index int) {
	resp, err := http.Get(url)
	if err != nil {
		log.Warn(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Warnf("Access failed: %s", url)
		fmt.Println("")
		return
	}
	size := resp.ContentLength

	if hasParDir, _ := PathExists(filepath.Dir(destFn)); !hasParDir {
		err := CreateFileParDir(destFn)
		if err != nil {
			log.Fatal(err)
		}
	}
	// create dest
	destName := filepath.Base(url)
	dest, err := os.Create(destFn)
	if err != nil {
		log.Lg.Warnf("Can't create %s: %v\n", destName, err)
		return
	}
	prefixStr := filepath.Base(destFn)
	prefixStrLen := utf8.RuneCountInString(prefixStr)
	if prefixStrLen > 35 {
		prefixStr = prefixStr[0:31] + "..."
	}
	bar := pg.AddBar(size,
		mpb.BarStyle("[=>-|"),
		mpb.PrependDecorators(
			decor.Name(fmt.Sprintf("%-35s\t", prefixStr), decor.WC{W: index, C: decor.DidentRight}),
			decor.CountersKibiByte("% -.1f / % -.1f\t"),
			decor.OnComplete(decor.Percentage(decor.WC{W: 5}), " "+"√"),
		),
		mpb.AppendDecorators(
			decor.EwmaETA(decor.ET_STYLE_MMSS, float64(size)/2048),
			decor.Name(" ] "),
			decor.AverageSpeed(decor.UnitKiB, "% .1f"),
		),
	)
	// create proxy reader
	reader := bar.ProxyReader(resp.Body)

	// and copy from reader, ignoring errors
	io.Copy(dest, reader)

	//pg.Wait()
	defer dest.Close()
}

// AsyncURL can access URL via using external commandline tools including wget, curl, axel, git and rsync
func AsyncURL(url string, destFn string, engine string, taskID string, mirror string, axelThread int, console bool) {
	if mirror != "" {
		if !strings.HasSuffix(mirror, "/") {
			mirror = mirror + "/"
		}
		url = mirror + filepath.Base(url)
	}
	if engine == "wget" {
		Wget(url, destFn, taskID, console)
	} else if engine == "curl" {
		Curl(url, destFn, taskID, console)
	} else if engine == "axel" {
		Axel(url, destFn, axelThread, taskID, console)
	} else if engine == "git" {
		Git(url, destFn, taskID, console)
	} else if engine == "rsync" {
		Rsync(url, destFn, taskID, console)
	}
}

// AsyncURL2 can access URL via using golang http library (with mbp progress bar) and
// external commandline tools including wget, curl, axel, git and rsync
func AsyncURL2(url string, destFn string, engine string, taskID string, mirror string,
	p *mpb.Progress, index int, axelThread int, console bool) {
	if engine == "go-http" {
		if mirror != "" {
			if !strings.HasSuffix(mirror, "/") {
				mirror = mirror + "/"
			}
			url = mirror + filepath.Base(url)
		}
		httpGetURL(url, destFn, p, index)
	} else {
		AsyncURL(url, destFn, engine, taskID, mirror, axelThread, console)
	}
}

// AsyncURL3 can access URL via using golang http library (with mbp progress bar) and
// external commandline tools including wget, curl, axel, git and rsync
func AsyncURL3(url string, destFn string, engine string, taskID string, mirror string,
	axelThread int, console bool) {
	index := 1
	if engine == "go-http" {
		if mirror != "" {
			if !strings.HasSuffix(mirror, "/") {
				mirror = mirror + "/"
			}
			url = mirror + filepath.Base(url)
		}
		httpGetURL(url, destFn, pg, index)
	} else {
		AsyncURL(url, destFn, engine, taskID, mirror, axelThread, console)
	}
}

// HTTPGetURLs can use golang http.Get and external commandline tools including wget, curl, axel, git and rsync
// to query URL with progress bar
func HTTPGetURLs(urls []string, destDir []string, enging string, taskID string, mirror string, concurrency int, axelThread int, overwrite bool, ignore bool) {
	sem := make(chan bool, concurrency)
	for j := range urls {
		url := urls[j]
		destFn := path.Join(destDir[j], path.Base(urls[j]))
		log.Infof("Trying %s => %s", url, destFn)
	}
	for j := range urls {
		CreateDir(destDir[j])
		destFn := path.Join(destDir[j], path.Base(urls[j]))
		if overwrite {
			os.RemoveAll(destFn)
		}
		if hasDestFn, _ := PathExists(destFn); !hasDestFn || ignore {
			url := urls[j]
			sem <- true
			go func(url string, destFn string) {
				time.Sleep(time.Second)
				defer func() {
					<-sem
				}()
				if len(urls) > 1 {
					AsyncURL2(url, destFn, enging, taskID, mirror, pg, j, axelThread, false)
				} else {
					AsyncURL2(url, destFn, enging, taskID, mirror, pg, j, axelThread, true)
				}
			}(url, destFn)
		} else {
			log.Infof("%s existed.", destFn)
		}
	}
	for i := 0; i < cap(sem); i++ {
		sem <- true
	}
}

func init() {
	pg = mpb.New(
		mpb.WithWidth(45),
		mpb.WithRefreshRate(180*time.Millisecond),
	)
}
