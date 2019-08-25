package cmd

import (
	"strings"

	butils "github.com/JhuangLab/butils"
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
		seqsTmp = butils.ReadLines(bgetClis.listFile)
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
					Geofetch(seqs[k][i], bgetClis.downloadDir, bgetClis.engine, bgetClis.concurrency,
						bgetClis.axelThread, cmdExtraFromFlag, taskID, overwrite, ignore, quiet, saveLog)
				} else if k == "sra" {
					Prefetch(seqs[k][i], "", bgetClis.downloadDir, cmdExtraFromFlag, taskID, quiet, saveLog)
				} else if k == "sraKrt" {
					Prefetch("", seqs[k][i], bgetClis.downloadDir, cmdExtraFromFlag, taskID, quiet, saveLog)
				} else if k == "tcgaManifest" {
					GdcClient("", seqs[k][i], bgetClis.downloadDir, bgetClis.gdcToken, cmdExtraFromFlag, taskID, quiet, saveLog)
				} else if k == "tcgaFileID" {
					GdcClient(seqs[k][i], "", bgetClis.downloadDir, bgetClis.gdcToken, cmdExtraFromFlag, taskID, quiet, saveLog)
				}
			}(k, i)
		}
	}
	for i := 0; i < cap(sem); i++ {
		sem <- true
	}
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
	seqCmd.Flags().StringVarP(&(bgetClis.engine), "engine", "g", "go-http", "Point the download engine: go-http, wget, curl, and axel.")
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
