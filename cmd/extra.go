package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	neturl "net/url"

	butils "github.com/JhuangLab/butils"
	"github.com/JhuangLab/butils/log"
)

func checkGitEngine(url string) string {
	if butils.StrDetect(url, "^git@") {
		return "git"
	}
	sites := []string{"https://github.com", "http://github.com",
		"https://gitlab.com", "https://gitlab.com", "https://bitbucket.org", "http://bitbucket.org"}
	for _, v := range sites {
		if butils.StrDetect(url, v) && strings.Count(url, "/") == 4 {
			return "git"
		}
	}
	return ""
}

func formatURLfileName(url string) (fname string) {
	u, _ := neturl.Parse(url)
	uQ := u.Query()
	fname = path.Base(url)
	// cell.com
	if butils.StrDetect(url, "/pdfExtended/") {
		fname = path.Base(url) + ".pdf"
	} else if butils.StrDetect(url, "showPdf[?]pii=") {
		fname = path.Base(butils.StrReplaceAll(url, "showPdf[?]pii=", "")) + ".pdf"
	} else if butils.StrDetect(url, "track/pdf") {
		fname = path.Base(url) + ".pdf"
	} else if butils.StrDetect(url, "&type=printable") {
		fname = strings.ReplaceAll(path.Base(url), "&type=printable", "") + ".pdf"
	} else if fname == "pdf" {
		fname = path.Base(strings.ReplaceAll(url, "/pdf", ".pdf"))
	} else if butils.StrDetect(fname, "[?]Expires=") {
		fname = butils.StrReplaceAll(fname, "[?]Expires=.*", "")
	} else if butils.StrDetect(url, "/action/downloadSupplement[?].*") {
		fname = butils.StrReplaceAll(fname, "downloadSupplement.*file=", "")
	} else if butils.StrDetect(url, "(.com/doi/pdf/)|(.org/doi/pdf/)|(.org/doi/pdfdirect/)") {
		if butils.StrDetect(url, "[?]articleTools=true") {
			fname = butils.StrReplaceAll(fname, "[?]articleTools=true", "")
		}
		fname = fname + ".pdf"
	} else if butils.StrDetect(url, "[?]md5=.*&pid=.*") {
		fname = butils.StrReplaceAll(fname, "[?]md5=.*&pid=", "")
	} else if butils.StrDetect(fname, "[?]download=true$") {
		fname = butils.StrReplaceAll(fname, "[?]download=true$", "")
	} else if butils.StrDetect(fname, "[?]_hash=.*") {
		fname = butils.StrReplaceAll(fname, "[?]_hash=.*", "")
	} else if butils.StrDetect(url, "sd/pdf/render") {
		fname = "supp." + fname + ".pdf"
	} else if strings.Contains(url, "https://www.ncbi.nlm.nih.gov/geo/download/?acc=") {
		if strings.Contains(url, "file&file=") {
			fname = uQ["file"][0]
		} else {
			fname = uQ["acc"][0] + ".tar"
		}
		fname, _ = neturl.QueryUnescape(fname)
	} else if strings.Contains(url, "www.ncbi.nlm.nih.gov/geo/query/acc") {
		fname = uQ["acc"][0] + ".txt"
	} else if strings.Contains(url, "https://www.ncbi.nlm.nih.gov/Traces/study/backends") &&
		strings.Contains(url, "rt_all") &&
		strings.Contains(url, "rs=") {
		fname = butils.StrReplaceAll(uQ["rs"][0], `[(]primary_search_ids:|[)]|"`, "") + "_SraRunTable.txt"
	} else if strings.Contains(url, "https://www.ncbi.nlm.nih.gov/Traces/study/backends") &&
		strings.Contains(url, "acc_all") &&
		strings.Contains(url, "rs=") {
		fname = butils.StrReplaceAll(uQ["rs"][0], `[(]primary_search_ids:|[)]|"`, "") + "_SraAccList.txt"
	}
	return fname
}

func checkQuiet() {
	if quiet {
		log.SetOutput(ioutil.Discard)
	} else {
		log.SetOutput(os.Stderr)
	}
}

func checkDownloadDir(condtions bool) {
	if hasDir, _ := butils.PathExists(bgetClis.downloadDir); !hasDir {
		if condtions {
			if err := butils.CreateDir(bgetClis.downloadDir); err != nil {
				log.FATAL(fmt.Sprintf("Could not to create %s", bgetClis.downloadDir))
			}
		}
	}
}
