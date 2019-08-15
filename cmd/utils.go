package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	log "github.com/JhuangLab/bget/log"
	"github.com/JhuangLab/bget/utils"
	mpb "github.com/vbauerster/mpb/v4"
	"github.com/vbauerster/mpb/v4/decor"
)

var pg *mpb.Progress

// Wget use wget to download files
func Wget(url string, destFn string, taskID string, quiet bool, saveLog bool) {
	args := []string{"-c", url, "-O", destFn}
	cmd := exec.Command("wget", args...)
	logPath := path.Join(logDir, fmt.Sprintf("%s_%s_wget.log", taskID, path.Base(destFn)))
	utils.CreateFileParDir(logPath)
	utils.RunExecCmdConsole(logPath, cmd, quiet, saveLog)
}

// Curl use curl to download files
func Curl(url string, destFn string, taskID string, quiet bool, saveLog bool) {
	args := []string{url, "-o", destFn}
	cmd := exec.Command("curl", args...)
	logPath := path.Join(logDir, fmt.Sprintf("%s_%s_curl.log", taskID, path.Base(destFn)))
	utils.CreateFileParDir(logPath)
	utils.RunExecCmdConsole(logPath, cmd, quiet, saveLog)
}

// Axel use axel to download files
func Axel(url string, destFn string, thread int, taskID string, quiet bool, saveLog bool) {
	args := []string{url, "-N", "-o", destFn, "-n", strconv.Itoa(thread)}
	cmd := exec.Command("axel", args...)
	logPath := path.Join(logDir, fmt.Sprintf("%s_%s_axel.log", taskID, path.Base(destFn)))
	utils.CreateFileParDir(logPath)
	utils.RunExecCmdConsole(logPath, cmd, quiet, saveLog)
}

// Git use git to download files
func Git(url string, destFn string, taskID string, quiet bool, saveLog bool) {
	args := []string{"clone", "--recursive", url, destFn}
	cmd := exec.Command("git", args...)
	logPath := path.Join(logDir, fmt.Sprintf("%s_%s_git.log", taskID, path.Base(destFn)))
	utils.CreateFileParDir(logPath)
	utils.RunExecCmdConsole(logPath, cmd, quiet, saveLog)
}

// Rsync use rsync to download files
func Rsync(url string, destFn string, taskID string, quiet bool, saveLog bool) {
	args := []string{url, destFn}
	cmd := exec.Command("rsync", args...)
	logPath := path.Join(logDir, fmt.Sprintf("%s_%s_rsync.log", taskID, path.Base(destFn)))
	utils.CreateFileParDir(logPath)
	utils.RunExecCmdConsole(logPath, cmd, quiet, saveLog)
}

// httpGetURL can use golang http.Get to query URL with progress bar
func httpGetURL(url string, destFn string, pg *mpb.Progress, index int, quiet bool, saveLog bool) {
	resp, err := http.Get(url)
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
	size := resp.ContentLength

	if hasParDir, _ := utils.PathExists(filepath.Dir(destFn)); !hasParDir {
		err := utils.CreateFileParDir(destFn)
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
		destFn := path.Join(destDir[j], path.Base(urls[j]))
		log.Infof("Trying %s => %s", url, destFn)
	}
	if len(urls) > 1 && concurrency > 1 && engine != "go-http" {
		quiet = true
	}
	for j := range urls {
		utils.CreateDir(destDir[j])
		destFn := path.Join(destDir[j], path.Base(urls[j]))
		if overwrite {
			err := os.RemoveAll(destFn)
			if err != nil {
				log.Warnf("Can not remove %s.", destFn)
			}
		}
		if hasDestFn, _ := utils.PathExists(destFn); !hasDestFn || ignore {
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

func gunzipDefuseReffa(destDir string) {
	dirList, e := ioutil.ReadDir(destDir)
	if e != nil {
		fmt.Println("read dir error")
		return
	}
	for _, v := range dirList {
		destFn := ""
		oldPath := path.Join(destDir, v.Name())
		if path.Base(v.Name()) == "rmsk.txt.gz" {
			destFn = path.Join(destDir, "repeats.txt.gz")
			log.Infof("Rename %s => %s", oldPath, destFn)
			err := os.Rename(oldPath, destFn)
			if err != nil {
				log.Warn(err)
			}
			destFn2 := utils.StrReplaceAll(destFn, ".gz$", "")
			gunzipLog(destFn, destFn2)
		} else if utils.StrDetect(v.Name(), ".fa.gz$") && utils.StrDetect(v.Name(), ".dna.") {
			destFn = utils.StrReplaceAll(v.Name(), ".fa.gz$", ".fa")
			destFn = utils.StrReplaceAll(destFn, "Homo_sapiens.*.dna.chromosome.", "defuse.dna.chromosomes.")
			destFn = path.Join(destDir, destFn)
			gunzipLog(oldPath, destFn)
		} else if utils.StrDetect(v.Name(), ".gz$") {
			destFn = utils.StrReplaceAll(v.Name(), ".gz$", "")
			destFn = path.Join(destDir, destFn)
			gunzipLog(oldPath, destFn)
		}
	}
}

func gunzipLog(oldPath string, destFn string) {
	if hasFn, _ := utils.PathExists(destFn); !hasFn {
		log.Infof("Uncompressing %s => %s", oldPath, destFn)
		_, err := utils.Gunzip(oldPath, destFn)
		if err != nil {
			log.Warn(err)
		}
	} else {
		log.Infof("%s existed.", destFn)
	}
}

func checkGitEngine(url string) string {
	if utils.StrDetect(url, "^git@") {
		return "git"
	}
	sites := []string{"https://github.com", "http://github.com",
		"https://gitlab.com", "https://gitlab.com", "https://bitbucket.org", "http://bitbucket.org"}
	for _, v := range sites {
		if utils.StrDetect(url, v) && strings.Count(url, "/") == 4 {
			return "git"
		}
	}
	return ""
}

func init() {
	pg = mpb.New(
		mpb.WithWidth(45),
		mpb.WithRefreshRate(180*time.Millisecond),
	)
}
