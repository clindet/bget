package cmd

import (
	"path"
	"strings"

	"github.com/JhuangLab/bget/utils"
)

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

func formatURLfileName(url string) (fname string) {
	fname = path.Base(url)
	// cell.com
	if utils.StrDetect(url, "/pdfExtended/") {
		fname = path.Base(url) + ".pdf"
	} else if utils.StrDetect(url, "showPdf[?]pii=") {
		fname = path.Base(utils.StrReplaceAll(url, "showPdf[?]pii=", "")) + ".pdf"
	} else if utils.StrDetect(url, "track/pdf") {
		fname = path.Base(url) + ".pdf"
	} else if utils.StrDetect(url, "&type=printable") {
		fname = strings.ReplaceAll(path.Base(url), "&type=printable", "") + ".pdf"
	} else if fname == "pdf" {
		fname = path.Base(strings.ReplaceAll(url, "/pdf", ".pdf"))
	} else if utils.StrDetect(fname, "[?]Expires=") {
		fname = utils.StrReplaceAll(fname, "[?]Expires=.*", "")
	} else if utils.StrDetect(url, "/action/downloadSupplement[?].*") {
		fname = utils.StrReplaceAll(fname, "downloadSupplement.*file=", "")
	} else if utils.StrDetect(url, "(www.embopress.org/doi/pdf)|(ascopubs.org/doi/pdfdirect)") {
		fname = fname + ".pdf"
	}
	return fname
}
