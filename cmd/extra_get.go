package cmd

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	butils "github.com/JhuangLab/butils"
	log "github.com/JhuangLab/butils/log"
	mpb "github.com/vbauerster/mpb/v4"
	"github.com/vbauerster/mpb/v4/decor"
)

var pg *mpb.Progress
var gCurCookies []*http.Cookie
var gCurCookieJar *cookiejar.Jar

// Wget use wget to download files
func Wget(url string, destFn string, taskID string, quiet bool, saveLog bool) {
	args := []string{"-c", url, "-O", destFn}
	cmd := exec.Command("wget", args...)
	logPath := path.Join(logDir, fmt.Sprintf("%s_%s_wget.log", taskID, path.Base(destFn)))
	butils.CreateFileParDir(logPath)
	butils.RunExecCmdConsole(logPath, cmd, quiet, saveLog)
}

// Curl use curl to download files
func Curl(url string, destFn string, taskID string, quiet bool, saveLog bool) {
	args := []string{url, "-o", destFn}
	cmd := exec.Command("curl", args...)
	logPath := path.Join(logDir, fmt.Sprintf("%s_%s_curl.log", taskID, path.Base(destFn)))
	butils.CreateFileParDir(logPath)
	butils.RunExecCmdConsole(logPath, cmd, quiet, saveLog)
}

// Axel use axel to download files
func Axel(url string, destFn string, thread int, taskID string, quiet bool, saveLog bool) {
	args := []string{url, "-N", "-o", destFn, "-n", strconv.Itoa(thread)}
	cmd := exec.Command("axel", args...)
	logPath := path.Join(logDir, fmt.Sprintf("%s_%s_axel.log", taskID, path.Base(destFn)))
	butils.CreateFileParDir(logPath)
	butils.RunExecCmdConsole(logPath, cmd, quiet, saveLog)
}

// Git use git to download files
func Git(url string, destFn string, taskID string, quiet bool, saveLog bool) {
	args := []string{"clone", "--recursive", url, destFn}
	cmd := exec.Command("git", args...)
	logPath := path.Join(logDir, fmt.Sprintf("%s_%s_git.log", taskID, path.Base(destFn)))
	butils.CreateFileParDir(logPath)
	butils.RunExecCmdConsole(logPath, cmd, quiet, saveLog)
}

// Rsync use rsync to download files
func Rsync(url string, destFn string, taskID string, quiet bool, saveLog bool) {
	args := []string{url, destFn}
	cmd := exec.Command("rsync", args...)
	logPath := path.Join(logDir, fmt.Sprintf("%s_%s_rsync.log", taskID, path.Base(destFn)))
	butils.CreateFileParDir(logPath)
	butils.RunExecCmdConsole(logPath, cmd, quiet, saveLog)
}

func checkHTTPGetURLRdirect(resp *http.Response, url string, destFn string, pg *mpb.Progress, index int, quiet bool, saveLog bool) (status bool) {
	if strings.Contains(url, "https://www.sciencedirect.com") {
		v, _ := ioutil.ReadAll(resp.Body)
		url = butils.StrExtract(string(v), `https://pdf.sciencedirectassets.com/.*&type=client`, 1)
		httpGetURL(url, destFn, pg, index, quiet, saveLog)
		return true
	}
	return false
}

func defaultCheckRedirect(req *http.Request, via []*http.Request) error {
	if len(via) >= 20 {
		return errors.New("stopped after 20 redirects")
	}
	return nil
}

// httpGetURL can use golang http.Get to query URL with progress bar
func httpGetURL(url string, destFn string, pg *mpb.Progress, index int, quiet bool, saveLog bool) {
	client := &http.Client{
		CheckRedirect: defaultCheckRedirect,
		Jar:           gCurCookieJar,
	}

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36")
	if err != nil {
		// handle error
		log.Warn(err)
		return
	}
	gCurCookies = gCurCookieJar.Cookies(req.URL)
	resp, err := client.Do(req)
	if err != nil {
		log.Warn(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if !quiet {
			log.Warnf("Access failed: %s", url)
			fmt.Println("")
		}
		return
	}
	if checkHTTPGetURLRdirect(resp, url, destFn, pg, index, quiet, saveLog) {
		return
	}
	size := resp.ContentLength

	if hasParDir, _ := butils.PathExists(filepath.Dir(destFn)); !hasParDir {
		err := butils.CreateFileParDir(destFn)
		if err != nil {
			log.Fatal(err)
		}
	}
	// create dest
	destName := filepath.Base(url)
	dest, err := os.Create(destFn)
	if err != nil {
		log.Warnf("Can't create %s: %v\n", destName, err)
		return
	}
	prefixStr := filepath.Base(destFn)
	prefixStrLen := utf8.RuneCountInString(prefixStr)
	if prefixStrLen > 35 {
		prefixStr = prefixStr[0:31] + "..."
	}
	if !quiet {
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
	} else {
		io.Copy(dest, io.Reader(resp.Body))
	}
	defer dest.Close()
}

// AsyncURL can access URL via using external commandline tools including wget, curl, axel, git and rsync
func AsyncURL(url string, destFn string, engine string, taskID string, mirror string, axelThread int, quiet bool, saveLog bool) {
	if mirror != "" {
		if !strings.HasSuffix(mirror, "/") {
			mirror = mirror + "/"
		}
		url = mirror + filepath.Base(url)
	}
	if checkGitEngine(url) == "git" {
		engine = "git"
	}
	if engine == "wget" {
		Wget(url, destFn, taskID, quiet, saveLog)
	} else if engine == "curl" {
		Curl(url, destFn, taskID, quiet, saveLog)
	} else if engine == "axel" {
		Axel(url, destFn, axelThread, taskID, quiet, saveLog)
	} else if engine == "git" {
		Git(url, destFn, taskID, quiet, saveLog)
	} else if engine == "rsync" {
		Rsync(url, destFn, taskID, quiet, saveLog)
	}
}

// AsyncURL2 can access URL via using golang http library (with mbp progress bar) and
// external commandline tools including wget, curl, axel, git and rsync
func AsyncURL2(url string, destFn string, engine string, taskID string, mirror string,
	p *mpb.Progress, index int, axelThread int, quiet bool, saveLog bool) {
	if checkGitEngine(url) == "git" {
		engine = "git"
	}
	if engine == "go-http" {
		if mirror != "" {
			if !strings.HasSuffix(mirror, "/") {
				mirror = mirror + "/"
			}
			url = mirror + filepath.Base(url)
		}
		httpGetURL(url, destFn, p, index, quiet, saveLog)
	} else {
		AsyncURL(url, destFn, engine, taskID, mirror, axelThread, quiet, saveLog)
	}
}

// AsyncURL3 can access URL via using golang http library (with mbp progress bar) and
// external commandline tools including wget, curl, axel, git and rsync
func AsyncURL3(url string, destFn string, engine string, taskID string, mirror string,
	axelThread int, quiet bool, saveLog bool) {
	if checkGitEngine(url) == "git" {
		engine = "git"
	}
	index := 1
	if engine == "go-http" {
		if mirror != "" {
			if !strings.HasSuffix(mirror, "/") {
				mirror = mirror + "/"
			}
			url = mirror + filepath.Base(url)
		}
		httpGetURL(url, destFn, pg, index, quiet, saveLog)
		pg.Wait()
	} else {
		AsyncURL(url, destFn, engine, taskID, mirror, axelThread, quiet, saveLog)
	}
}

// HTTPGetURLs can use golang http.Get and external commandline tools including wget, curl, axel, git and rsync
// to query URL with progress bar
func HTTPGetURLs(urls []string, destDir []string, engine string, taskID string, mirror string, concurrency int, axelThread int, overwrite bool, ignore bool, quiet bool, saveLog bool) {
	sem := make(chan bool, concurrency)
	for j := range urls {
		url := urls[j]
		destFn := path.Join(destDir[j], formatURLfileName(urls[j]))
		log.Infof("Trying %s => %s", url, destFn)
	}
	if len(urls) > 1 && concurrency > 1 && engine != "go-http" {
		quiet = true
	}
	for j := range urls {
		butils.CreateDir(destDir[j])
		destFn := path.Join(destDir[j], formatURLfileName(urls[j]))
		if overwrite {
			err := os.RemoveAll(destFn)
			if err != nil {
				log.Warnf("Can not remove %s.", destFn)
			}
		}
		if hasDestFn, _ := butils.PathExists(destFn); !hasDestFn || ignore {
			url := urls[j]
			sem <- true
			go func(url string, destFn string) {
				defer func() {
					<-sem
				}()
				AsyncURL2(url, destFn, engine, taskID, mirror, pg, j, axelThread, quiet, saveLog)
			}(url, destFn)
		} else {
			log.Infof("%s existed.", destFn)
		}
	}
	for i := 0; i < cap(sem); i++ {
		sem <- true
	}
	pg.Wait()
}

func init() {
	pg = mpb.New(
		mpb.WithWidth(45),
		mpb.WithRefreshRate(180*time.Millisecond),
	)

	gCurCookies = nil
	//var err error;
	gCurCookieJar, _ = cookiejar.New(nil)
}
