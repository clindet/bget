package cmd

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/cookiejar"
	neturl "net/url"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode/utf8"

	"github.com/Miachol/bget/spider"
	butils "github.com/openbiox/butils"
	log "github.com/openbiox/butils/log"
	mpb "github.com/vbauerster/mpb/v4"
	"github.com/vbauerster/mpb/v4/decor"
)

var pg *mpb.Progress
var gCurCookies []*http.Cookie
var gCurCookieJar *cookiejar.Jar

func setCmdProxyEnv(cmd *exec.Cmd, proxy, url string) {
	if proxy == "" {
		return
	}
	proxyFlag := ""
	if strings.Contains(url, "https://") {
		proxyFlag = "https_proxy=" + proxy
	} else if strings.Contains(url, "ftp://") {
		proxyFlag = "ftp_proxy=" + proxy
	} else {
		proxyFlag = "http_proxy=" + proxy
	}
	_, proxy = spider.RandProxy(proxy)
	cmd.Env = append(cmd.Env, proxyFlag)
	cmd.Env = append(cmd.Env, "use_proxy=on")
}

// Wget use wget to download files
func Wget(url string, destFn string, extraArgs string, taskID string, quiet bool, saveLog bool, retries int, timeout int, retSleepTime int) (err error) {
	if url == "" {
		return errors.New("at least one of URL is required")
	}
	args := []string{"-c", url, "-O", destFn, "--timeout=" +
		strconv.Itoa(timeout), "--show-progress"}
	if extraArgs != "" {
		extraArgsList := strings.Split(extraArgs, " ")
		args = append(args, extraArgsList...)
	}
	cmd := exec.Command("wget", args...)
	setCmdProxyEnv(cmd, bgetClis.proxy, url)
	logPath := path.Join(logDir, fmt.Sprintf("%s_%s_wget.log", taskID, path.Base(destFn)))
	if saveLog {
		log.Infof("Download Log of %s saved to => %s", url, logPath)
		butils.CreateFileParDir(logPath)
	}
	time.Sleep(1 * time.Second)
	err = retryURL(url, retries, retSleepTime, logPath, cmd, quiet, saveLog)
	if err != nil {
		return err
	}
	return nil
}

// Curl use curl to download files
func Curl(url string, destFn string, extraArgs string, taskID string, quiet bool, saveLog bool, retries int, timeout int, retSleepTime int) (err error) {
	if url == "" {
		return errors.New("at least one of URL is required")
	}
	args := []string{url, "-o", destFn, "--connect-timeout", strconv.Itoa(timeout)}
	if extraArgs != "" {
		extraArgsList := strings.Split(extraArgs, " ")
		args = append(args, extraArgsList...)
	}
	cmd := exec.Command("curl", args...)
	setCmdProxyEnv(cmd, bgetClis.proxy, url)
	logPath := path.Join(logDir, fmt.Sprintf("%s_%s_curl.log", taskID, path.Base(destFn)))
	if saveLog {
		log.Infof("Download Log of %s saved to => %s", url, logPath)
		butils.CreateFileParDir(logPath)
	}
	time.Sleep(1 * time.Second)
	err = retryURL(url, retries, retSleepTime, logPath, cmd, quiet, saveLog)
	if err != nil {
		return err
	}
	return nil
}

// Axel use axel to download files
func Axel(url string, destFn string, thread int, extraArgs string, taskID string, quiet bool, saveLog bool, retries int, timeout int, retSleepTime int) (err error) {
	if url == "" {
		return errors.New("at least one of URL is required")
	}
	args := []string{url, "-N", "-o", destFn, "-n", strconv.Itoa(thread),
		"--timeout=" + strconv.Itoa(timeout)}
	if extraArgs != "" {
		extraArgsList := strings.Split(extraArgs, " ")
		args = append(args, extraArgsList...)
	}
	cmd := exec.Command("axel", args...)
	setCmdProxyEnv(cmd, bgetClis.proxy, url)
	logPath := path.Join(logDir, fmt.Sprintf("%s_%s_axel.log", taskID, path.Base(destFn)))
	if saveLog {
		log.Infof("Download Log of %s saved to => %s", url, logPath)
		butils.CreateFileParDir(logPath)
	}
	time.Sleep(1 * time.Second)
	err = retryURL(url, retries, retSleepTime, logPath, cmd, quiet, saveLog)
	if err != nil {
		return err
	}
	return nil
}

// Git use git to download files
func Git(url string, destFn string, extraArgs string, taskID string, quiet bool, saveLog bool, retries int, timeout int, retSleepTime int) (err error) {
	if url == "" {
		return errors.New("at least one of URL is required")
	}
	args := []string{"clone", "--recursive", "--progress"}
	if extraArgs != "" {
		extraArgsList := strings.Split(extraArgs, " ")
		args = append(args, extraArgsList...)
	}
	args = append(args, url, destFn)
	cmd := exec.Command("git", args...)
	setCmdProxyEnv(cmd, bgetClis.proxy, url)
	logPath := path.Join(logDir, fmt.Sprintf("%s_%s_git.log", taskID, path.Base(destFn)))
	if saveLog {
		log.Infof("Download Log of %s saved to => %s", url, logPath)
		butils.CreateFileParDir(logPath)
	}
	time.Sleep(1 * time.Second)
	err = retryURL(url, retries, retSleepTime, logPath, cmd, quiet, saveLog)
	if err != nil {
		return err
	}
	return nil
}

// Rsync use rsync to download files
func Rsync(url string, destFn string, extraArgs string, taskID string, quiet bool, saveLog bool, retries int, timeout int, retSleepTime int) (err error) {
	if url == "" {
		return errors.New("at least one of URL is required")
	}
	args := []string{url, destFn}
	if extraArgs != "" {
		extraArgsList := strings.Split(extraArgs, " ")
		args = append(args, extraArgsList...)
	}
	cmd := exec.Command("rsync", args...)
	setCmdProxyEnv(cmd, bgetClis.proxy, url)
	logPath := path.Join(logDir, fmt.Sprintf("%s_%s_rsync.log", taskID, path.Base(destFn)))
	if saveLog {
		log.Infof("Download Log of %s saved to => %s", url, logPath)
		butils.CreateFileParDir(logPath)
	}
	time.Sleep(1 * time.Second)
	err = retryURL(url, retries, retSleepTime, logPath, cmd, quiet, saveLog)
	if err != nil {
		return err
	}
	return nil
}

// GdcClient use gdc-client to download files
func GdcClient(fileID string, manifest string, outDir string, token string, extraArgs string, taskID string, quiet bool, saveLog bool, retries int, timeout int, retSleepTime int) (err error) {
	if fileID == "" && manifest == "" {
		return errors.New("at least one of fileID or manifest is required")
	}
	args := []string{}
	if manifest == "" {
		args = []string{"download", fileID, "-d", outDir}
	} else {
		args = []string{"download", "-m", manifest, "-d", outDir}
	}
	if extraArgs != "" {
		extraArgsList := strings.Split(extraArgs, " ")
		args = append(args, extraArgsList...)
	}
	if token != "" {
		args = append(args, "-t", token)
	}
	cmd := exec.Command("gdc-client", args...)
	setCmdProxyEnv(cmd, bgetClis.proxy, "")
	logPath := path.Join(logDir, fmt.Sprintf("%s_gdc-client.log", taskID))
	taskName := ""
	if fileID != "" {
		taskName = fileID
	} else {
		taskName = manifest
	}
	if saveLog {
		log.Infof("Download Log of %s saved to => %s", taskName, logPath)
		butils.CreateFileParDir(logPath)
	}
	time.Sleep(1 * time.Second)
	err = retriesTask(taskName, retries, retSleepTime, logPath, cmd, quiet, saveLog)
	if err != nil {
		return err
	}
	return nil
}

// Prefetch use sra-tools prefetch to download files
func Prefetch(srr string, krt string, outDir string, extraArgs string, taskID string, quiet bool, saveLog bool, retries int, timeout int, retSleepTime int) (err error) {
	if srr == "" && krt == "" {
		return errors.New("at least one of srr or krt is required")
	}
	args := []string{"-O", outDir, "-X", "500GB"}
	if extraArgs != "" {
		extraArgsList := strings.Split(extraArgs, " ")
		args = append(args, extraArgsList...)
	}
	if krt == "" {
		args = append(args, srr)
	} else {
		args = append(args, krt)
	}

	cmd := exec.Command("prefetch", args...)
	setCmdProxyEnv(cmd, bgetClis.proxy, "")
	logPath := path.Join(logDir, fmt.Sprintf("%s_prefetch.log", taskID))
	taskName := ""
	if srr != "" {
		taskName = srr
	} else {
		taskName = krt
	}
	if saveLog {
		log.Infof("Download Log of %s saved to => %s", taskName, logPath)
		butils.CreateFileParDir(logPath)
	}
	time.Sleep(1 * time.Second)
	err = retriesTask(taskName, retries, retSleepTime, logPath, cmd, quiet, saveLog)
	if err != nil {
		return err
	}
	return nil
}

// Geofetch get GEO files
func Geofetch(geo string, outDir string, engine string, concurrency int, axelThread int, extraArgs string, taskID string, overwrite bool, ignore bool, quiet bool, saveLog bool, retries int, timeout int, retSleepTime int, remoteName bool) (err error) {
	if geo == "" {
		return errors.New("at least one of geo is required")
	}
	gseURLs, gplURLs, sraLink := spider.GeoSpider(geo, bgetClis.proxy, timeout)
	u, _ := neturl.Parse(sraLink)
	uQ := u.Query()
	accAll := fmt.Sprintf(`https://www.ncbi.nlm.nih.gov/Traces/study/backends/solr_proxy/solr_proxy.cgi?core=run_sel_index&action=acc_all&fl=acc_s&rs=(primary_search_ids:"%s")`, uQ["acc"][0])
	rtAll := `https://www.ncbi.nlm.nih.gov/Traces/study/backends/solr_proxy/solr_proxy.cgi?core=run_sel_index&action=rt_all&fl=acc_s%2Cantibody_sam_ss%2Cassay_type_s%2Cavgspotlen_l%2Cbioproject_s%2Cbiosample_s%2Ccell_line_sam_ss_dpl110_ss%2Ccenter_name_s%2Cconsent_s%2Cdatastore_filetype_ss%2Cdatastore_provider_ss%2Cdatastore_region_ss%2Cexperiment_s%2Cgeo_accession_exp_ss%2Cinstrument_s%2Clibrary_name_s%2Clibrarylayout_s%2Clibraryselection_s%2Clibrarysource_s%2Cmbases_l%2Cmbytes_l%2Corganism_s%2Cplatform_s%2Creleasedate_dt%2Csample_acc_s%2Csample_name_s%2Csource_name_sam_ss%2Csra_study_s%2Ctreatment_sam_ss%2Cquality_book_char_run_ss&ft=Run%2CAntibody%2CAssay+Type%2CAvgSpotLen%2CBioProject%2CBioSample%2Ccell_line%2CCenter+Name%2CConsent%2CDATASTORE+filetype%2CDATASTORE+provider%2CDATASTORE+region%2CExperiment%2CGEO_Accession%2CInstrument%2CLibrary+Name%2CLibraryLayout%2CLibrarySelection%2CLibrarySource%2CMBases%2CMBytes%2COrganism%2CPlatform%2CReleaseDate%2Csample_acc%2CSample+Name%2Csource_name%2CSRA+Study%2Ctreatment%2Cquality_book_char&rs=%28primary_search_ids%3A%22` + uQ["acc"][0] + `%22%29`
	sraLinks := []string{accAll, rtAll}
	urls := append(gseURLs, gplURLs...)
	urls = append(urls, sraLinks...)
	destDirArray := []string{}
	for range urls {
		destDirArray = append(destDirArray, outDir)
	}
	done := HTTPGetURLs(urls, destDirArray, engine, extraArgs, taskID, "",
		concurrency, axelThread, overwrite, ignore, quiet, saveLog,
		retries, timeout, retSleepTime, remoteName)
	for _, dest := range done {
		if bgetClis.uncompress {
			if err := butils.UnarchiveLog(dest, path.Dir(dest)); err != nil {
				log.Warn(err)
			}
		}
	}
	return nil
}
func checkHTTPGetURLRdirect(resp *http.Response, url string, destFn string, pg *mpb.Progress, quiet bool, saveLog bool, retries int, timeout int, retSleepTime int) (status bool) {
	if strings.Contains(url, "https://www.sciencedirect.com") {
		v, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			if butils.StrDetect(string(v), "https://pdf.sciencedirectassets.com") {
				url = butils.StrExtract(string(v), `https://pdf.sciencedirectassets.com/.*&type=client`, 1)[0]
				httpGetURL(url, destFn, pg, quiet, saveLog, retries, timeout, retSleepTime)
				return true
			}
		}
	}
	return false
}

func defaultCheckRedirect(req *http.Request, via []*http.Request) error {
	if len(via) >= 20 {
		return errors.New("stopped after 20 redirects")
	}
	return nil
}

func newHTTPClient(timeout int) *http.Client {
	transPort := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: time.Duration(timeout) * time.Second,
		}).Dial,
	}
	if bgetClis.proxy != "" {
		urlproxy, _ := spider.RandProxy(bgetClis.proxy)
		transPort.Proxy = http.ProxyURL(urlproxy)
	}
	return &http.Client{
		CheckRedirect: defaultCheckRedirect,
		Jar:           gCurCookieJar,
		Transport:     transPort,
	}
}

// httpGetURL can use golang http.Get to query URL with progress bar
func httpGetURL(url string, destFn string, pg *mpb.Progress, quiet bool,
	saveLog bool, retries int, timeout int, retSleepTime int) error {
	client := newHTTPClient(timeout)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36")
	if err != nil {
		// handle error
		log.Warn(err)
		return err
	}
	gCurCookies = gCurCookieJar.Cookies(req.URL)

	var t int
	var success = false
	for t = 0; t < retries; t++ {
		err = downloadWorker(client, req, url, destFn, pg, quiet, saveLog, retries, timeout, retSleepTime)
		if err == nil {
			success = true
			break
		} else {
			log.Warnf("Failed to retrive on attempt %d... error: %v ... retrying after %d seconds.", t+1, err, retSleepTime)
			time.Sleep(time.Duration(retSleepTime) * time.Second)
		}
	}
	if !success {
		return err
	}
	return nil
}

// AsyncURL can access URL via using external commandline tools including wget, curl, axel, git and rsync
func AsyncURL(url string, destFn string, engine string, extraArgs string, taskID string, mirror string, axelThread int, quiet bool, saveLog bool, retries int, timeout int, retSleepTime int) error {
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
		return Wget(url, destFn, extraArgs, taskID, quiet, saveLog, retries, timeout, retSleepTime)
	} else if engine == "curl" {
		return Curl(url, destFn, extraArgs, taskID, quiet, saveLog, retries, timeout, retSleepTime)
	} else if engine == "axel" {
		return Axel(url, destFn, axelThread, extraArgs, taskID, quiet, saveLog, retries, timeout, retSleepTime)
	} else if engine == "git" {
		return Git(url, destFn, extraArgs, taskID, quiet, saveLog, retries, timeout, retSleepTime)
	} else if engine == "rsync" {
		return Rsync(url, destFn, extraArgs, taskID, quiet, saveLog, retries, timeout, retSleepTime)
	}
	return nil
}

// AsyncURL2 can access URL via using golang http library (with mbp progress bar) and
// external commandline tools including wget, curl, axel, git and rsync
func AsyncURL2(url string, destFn string, engine string, extraArgs string, taskID string,
	mirror string, p *mpb.Progress, axelThread int, quiet bool, saveLog bool, retries int, timeout int, retSleepTime int) error {
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
		return httpGetURL(url, destFn, p, quiet, saveLog, retries, timeout, retSleepTime)
	} else {
		return AsyncURL(url, destFn, engine, extraArgs, taskID, mirror, axelThread, quiet, saveLog, retries, timeout, retSleepTime)
	}
}

// AsyncURL3 can access URL via using golang http library (with mbp progress bar) and
// external commandline tools including wget, curl, axel, git and rsync
func AsyncURL3(url string, destFn string, engine string, extraArgs string, taskID string,
	mirror string, axelThread int, quiet bool, saveLog bool, retries int, timeout int, retSleepTime int) (err error) {
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
		err = httpGetURL(url, destFn, pg, quiet, saveLog, retries, timeout, retSleepTime)
		pg.Wait()
	} else {
		err = AsyncURL(url, destFn, engine, extraArgs, taskID, mirror, axelThread, quiet, saveLog, retries, timeout, retSleepTime)
	}
	return nil
}

// HTTPGetURLs can use golang http.Get and external commandline tools including wget, curl, axel, git and rsync
// to query URL with progress bar
func HTTPGetURLs(urls []string, destDir []string, engine string, extraArgs string, taskID string, mirror string, concurrency int, axelThread int, overwrite bool, ignore bool, quiet bool, saveLog bool, retries int, timeout int, retSleepTime int, remoteName bool) (destFns []string) {
	sem := make(chan bool, concurrency)
	if len(urls) > 1 && concurrency > 1 && engine != "go-http" {
		quiet = true
	}
	wg := sync.WaitGroup{}
	destMap := make(map[string]string)
	wg.Add(len(urls))
	for j := range urls {
		url := urls[j]
		dest := destDir[j]
		go func(url, dest string) {
			destMap[url] = path.Join(dest,
				formatURLfileName(url, remoteName))
			wg.Done()
		}(url, dest)
	}
	wg.Wait()
	for k, v := range destMap {
		log.Infof("Trying %s => %s", k, v)
	}
	for j := range urls {
		butils.CreateDir(destDir[j])
		destFn := destMap[urls[j]]
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
				err := AsyncURL2(url, destFn, engine, extraArgs, taskID, mirror, pg, axelThread, quiet, saveLog, retries, timeout, retSleepTime)
				if err == nil {
					destFns = append(destFns, destFn)
				}
			}(url, destFn)
		} else {
			destFns = append(destFns, destFn)
			log.Infof("%s existed.", destFn)
		}
	}
	for i := 0; i < cap(sem); i++ {
		sem <- true
	}
	pg.Wait()
	return destFns
}

func retryURL(url string, retries int, retSleepTime int,
	logPath string, cmd *exec.Cmd, quiet bool, saveLog bool) (err error) {
	var t int
	success := false
	var cmdRun exec.Cmd
	for t = 0; t < retries; t++ {
		cmdRun = *cmd
		err := butils.RunExecCmdConsole(logPath, &cmdRun, quiet, saveLog)
		if err == nil {
			success = true
			break
		} else {
			log.Warnf("Failed to retrive on attempt %d... error: %v ... retrying after %d seconds.", t+1, err, retSleepTime)
			time.Sleep(time.Duration(retSleepTime) * time.Second)
		}
	}
	if !success {
		return errors.New("Faild to access: " + url)
	}
	return nil
}

func retriesTask(taskName string, retries int, retSleepTime int,
	logPath string, cmd *exec.Cmd, quiet bool, saveLog bool) (err error) {
	var t int
	success := false
	var cmdRun exec.Cmd
	for t = 0; t < retries; t++ {
		cmdRun = *cmd
		err := butils.RunExecCmdConsole(logPath, &cmdRun, quiet, saveLog)
		if err == nil {
			success = true
			break
		} else {
			log.Warnf("Failed to retrive on attempt %d... error: %v ... retrying after %d seconds.", t+1, err, retSleepTime)
			time.Sleep(time.Duration(retSleepTime) * time.Second)
		}
	}
	if !success {
		return errors.New("Faild to access: " + taskName)
	}
	return nil
}

func downloadWorker(client *http.Client, req *http.Request, url string, destFn string, pg *mpb.Progress, quiet bool, saveLog bool, retries int, timeout int, retSleepTime int) error {
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if !quiet {
			log.Warnf("Access failed: %s", url)
			fmt.Println("")
		}
		return err
	}
	if checkHTTPGetURLRdirect(resp, url, destFn, pg, quiet, saveLog,
		retries, timeout, retSleepTime) {
		//return errors.New("Redirect fail")
		return nil
	}
	size := resp.ContentLength

	if hasParDir, _ := butils.PathExists(filepath.Dir(destFn)); !hasParDir {
		err = butils.CreateFileParDir(destFn)
		if err != nil {
			return err
		}
	}
	// create dest
	destName := filepath.Base(url)
	dest, err := os.Create(destFn)
	if err != nil {
		log.Warnf("Can't create %s: %v\n", destName, err)
		return err
	}
	defer dest.Close()
	prefixStr := filepath.Base(destFn)
	prefixStrLen := utf8.RuneCountInString(prefixStr)
	if prefixStrLen > 35 {
		prefixStr = prefixStr[0:31] + "..."
	}
	prefixStr = fmt.Sprintf("%-35s\t", prefixStr)
	if !quiet {
		bar := pg.AddBar(size,
			mpb.BarStyle("[=>-|"),
			mpb.PrependDecorators(
				decor.Name(prefixStr, decor.WC{W: len(prefixStr) + 1, C: decor.DidentRight}),
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
		_, err = io.Copy(dest, reader)
		if err != nil {
			reader.Close()
			bar.Abort(true)
			return err
		}
	} else {
		_, err = io.Copy(dest, io.Reader(resp.Body))
		if err != nil {
			return err
		}
	}
	return nil
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
