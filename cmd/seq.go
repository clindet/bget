package cmd

import (
	"path"
	"strings"

	"github.com/JhuangLab/bget/spider"
	butils "github.com/JhuangLab/butils"
	log "github.com/JhuangLab/butils/log"
	"github.com/spf13/cobra"
)

var seqCmd = &cobra.Command{
	Use:   "seq [id1 id2 id3... | manifest1 manifest2 manifest3...]",
	Short: "Can be used to access sequence data via unique id (dbGAP and EGA) or manifest files (TCGA).",
	Long:  `Can be used to access sequence data via unique id or manifest files. More see here https://github.com/JhuangLab/bget.`,
	Run: func(cmd *cobra.Command, args []string) {
		seqCmdRunOptions(cmd)
	},
}

func parseSeq() (seqs map[string][]string) {
	seqs = make(map[string][]string)
	seqsTmp := []string{}
	if bgetClis.seqs != "" && strings.Contains(bgetClis.seqs, bgetClis.separator) {
		seqsTmp = strings.Split(bgetClis.seqs, bgetClis.separator)
	} else if bgetClis.seqs != "" {
		seqsTmp = []string{bgetClis.seqs}
	} else if bgetClis.listFile != "" {
		seqsTmp = butils.ReadLines(bgetClis.seqs)
	}
	for i := range seqsTmp {
		if butils.StrDetect(strings.ToUpper(seqsTmp[i]), "^GSE|^GPL|^GDS") {
			seqs["geo"] = append(seqs["geo"], seqsTmp[i])
		} else if butils.StrDetect(strings.ToUpper(seqsTmp[i]), "^SRR|^ERR") {
			seqs["sra"] = append(seqs["sra"], seqsTmp[i])
		} else if butils.StrDetect(strings.ToLower(seqsTmp[i]), ".krt$") {
			seqs["sraKrt"] = append(seqs["sraKrt"], seqsTmp[i])
		} else if butils.StrDetect(strings.ToUpper(seqsTmp[i]), "^EGA") {
			seqs["ega"] = append(seqs["ega"], seqsTmp[i])
		} else if butils.StrDetect(strings.ToLower(seqsTmp[i]), ".txt$") {
			seqs["tcgaManifest"] = append(seqs["tcgaManifest"], seqsTmp[i])
		} else {
			seqs["tcgaFileID"] = append(seqs["tcgaFileID"], seqsTmp[i])
		}
	}
	return seqs
}

func downloadSeq() {
	seqs := parseSeq()
	done := make(map[string][]string)
	sem := make(chan bool, bgetClis.concurrency)
	for k, v := range seqs {
		for i := range v {
			sem <- true
			done[k] = append(done[k], seqs[k][i])
			go func(k string, i int) {
				defer func() {
					<-sem
				}()
				if k == "geo" {
					clientGetSeq(seqs[k][i], "", "", "", "", "", bgetClis.downloadDir, "", cmdExtraFromFlag,
						taskID, overwrite, ignore, quiet, saveLog)
				} else if k == "sra" {
					clientGetSeq("", seqs[k][i], "", "", "", "", bgetClis.downloadDir, "", cmdExtraFromFlag,
						taskID, overwrite, ignore, quiet, saveLog)
				} else if k == "sraKrt" {
					clientGetSeq("", "", seqs[k][i], "", "", "", bgetClis.downloadDir, "", cmdExtraFromFlag,
						taskID, overwrite, ignore, quiet, saveLog)
				} else if k == "tcgaManifest" {
					clientGetSeq("", "", "", "", seqs[k][i], "", bgetClis.downloadDir, bgetClis.gdcToken, cmdExtraFromFlag,
						taskID, overwrite, ignore, quiet, saveLog)
				} else if k == "tcgaFileID" {
					clientGetSeq("", "", "", "", "", seqs[k][i], bgetClis.downloadDir, bgetClis.gdcToken, cmdExtraFromFlag,
						taskID, overwrite, ignore, quiet, saveLog)
				}
			}(k, i)
		}
	}
	for i := 0; i < cap(sem); i++ {
		sem <- true
	}
}

func clientGetSeq(geo string, sra string, sraKrt string, ega string, tcgaManifest string, tcgaFileID string, destDir string, token string, extraArgs string, taskID string, overwrite bool, ignore bool, quiet bool, saveLog bool) (done []string) {
	if sra != "" || sraKrt != "" {
		Prefetch(sra, sraKrt, destDir, extraArgs, taskID, quiet, saveLog)
	} else if ega != "" {

	} else if tcgaFileID != "" || tcgaManifest != "" {
		GdcClient(tcgaFileID, tcgaManifest, destDir, token, extraArgs, taskID, quiet, saveLog)
	} else if geo != "" {
		gseURLs, gplURLs := spider.GeoSpider(geo)
		urls := append(gseURLs, gplURLs...)
		destDirArray := []string{}
		for range urls {
			destDirArray = append(destDirArray, bgetClis.downloadDir)
		}
		done := HTTPGetURLs(urls, destDirArray, bgetClis.engine, cmdExtraFromFlag, taskID, bgetClis.mirror,
			bgetClis.concurrency, bgetClis.axelThread, overwrite, ignore, quiet, saveLog)
		for _, dest := range done {
			if bgetClis.uncompress {
				if err := butils.UnarchiveLog(dest, path.Dir(dest)); err != nil {
					log.Warn(err)
				}
			}
		}
	}
	return done
}

func seqCmdRunOptions(cmd *cobra.Command) {
	checkQuiet()
	items := []string{}
	if len(cmd.Flags().Args()) >= 1 {
		items = append(items, cmd.Flags().Args()...)
		bgetClis.seqs = strings.Join(items, bgetClis.separator)
	}

	checkDownloadDir(bgetClis.seqs != "")
	if bgetClis.seqs != "" || bgetClis.listFile != "" {
		downloadSeq()
		bgetClis.helpFlags = false
	}
	if bgetClis.helpFlags {
		cmd.Help()
	}
}

func init() {
	seqCmd.Flags().BoolVarP(&(bgetClis.uncompress), "uncompress", "u", false, "Uncompress download files for .zip, .tar.gz, and .gz suffix files (now support GEO database).")
	seqCmd.Flags().StringVarP(&(bgetClis.gdcToken), "gdc-token", "", "", "Token to access TCGA portal files.")
	seqCmd.Flags().StringVarP(&(bgetClis.listFile), "list-file", "l", "", "A file contains seq id (e.g. SRR) or manifest files for download.")
	seqCmd.Example = `  bget seq ERR3324530
  split -a 3 --additional-suffix=.txt -l 100 gdc_manifest.2019-08-23-TCGA.txt -d
  for i in x*.txt
  do
    head -n 1 x000.txt > ${i}.tmp
    cat ${i} >> ${i}.tmp
    mv ${i}.tmp ${i}
  done
  sed -i '1d' x000.txt
  bget seq *.txt -t 5`
}
