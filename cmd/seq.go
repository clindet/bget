package cmd

import (
	"strings"

	"github.com/clindet/bget/spider"
	cio "github.com/openbiox/ligo/io"
	cnet "github.com/openbiox/ligo/net"
	"github.com/openbiox/ligo/stringo"
	"github.com/spf13/cobra"
)

var SeqCmd = &cobra.Command{
	Use:   "seq [id1 id2 id3... | manifest1 manifest2 manifest3...]",
	Short: "Can be used to access sequence data via unique id (dbGAP and EGA) or manifest files (TCGA).",
	Long:  `Can be used to access sequence data via unique id or manifest files. More see here https://github.com/clindet/bget.`,
	Run: func(cmd *cobra.Command, args []string) {
		seqCmdRunOptions(cmd, args)
	},
}

func seqCmdRunOptions(cmd *cobra.Command, args []string) {
	initCmd(cmd, args)
	checkArgs(cmd, "seq")
	checkDownloadDir(bgetClis.Seqs != "")
	if bgetClis.Seqs != "" || bgetClis.ListFile != "" {
		downloadSeq()
		bgetClis.HelpFlags = false
	}
	if bgetClis.HelpFlags {
		cmd.Help()
	}
}

func parseSeq() (seqs map[string][]string) {
	seqs = make(map[string][]string)
	seqsTmp := []string{}
	if bgetClis.Seqs != "" && strings.Contains(bgetClis.Seqs, bgetClis.Seperator) {
		seqsTmp = strings.Split(bgetClis.Seqs, bgetClis.Seperator)
	} else if bgetClis.Seqs != "" {
		seqsTmp = []string{bgetClis.Seqs}
	} else if bgetClis.ListFile != "" {
		seqsTmp = cio.ReadLines(bgetClis.ListFile)
	}
	for i := range seqsTmp {
		if stringo.StrDetect(strings.ToUpper(seqsTmp[i]), "^GSE|^GPL|^GDS|^GSM") {
			seqs["geo"] = append(seqs["geo"], seqsTmp[i])
		} else if stringo.StrDetect(strings.ToUpper(seqsTmp[i]), "^SRR|^ERR") {
			seqs["sra"] = append(seqs["sra"], seqsTmp[i])
		} else if stringo.StrDetect(strings.ToLower(seqsTmp[i]), ".krt$") {
			seqs["sraKrt"] = append(seqs["sraKrt"], seqsTmp[i])
		} else if stringo.StrDetect(strings.ToUpper(seqsTmp[i]), "^EGAD") {
			seqs["egad"] = append(seqs["egad"], seqsTmp[i])
		} else if stringo.StrDetect(strings.ToUpper(seqsTmp[i]), "^EGAF") {
			seqs["egaf"] = append(seqs["egaf"], seqsTmp[i])
		} else if stringo.StrDetect(strings.ToLower(seqsTmp[i]), ".txt$") {
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
	sem := make(chan bool, bgetClis.Thread)
	netOpt := setNetParams(&bgetClis)
	for k, v := range seqs {
		for i := range v {
			sem <- true
			done[k] = append(done[k], seqs[k][i])
			go func(k string, i int) {
				defer func() {
					<-sem
				}()
				if k == "sra" {
					cnet.Prefetch(seqs[k][i], "", bgetClis.DownloadDir, netOpt)
				} else if k == "sraKrt" {
					cnet.Prefetch("", seqs[k][i], bgetClis.DownloadDir, netOpt)
				} else if k == "tcgaManifest" {
					cnet.GdcClient("", seqs[k][i], bgetClis.DownloadDir, netOpt)
				} else if k == "tcgaFileID" {
					cnet.GdcClient(seqs[k][i], "", bgetClis.DownloadDir, netOpt)
				} else if k == "egad" {
					cnet.Egafetch(seqs[k][i], "", bgetClis.DownloadDir, netOpt)
				} else if k == "egaf" {
					cnet.Egafetch("", seqs[k][i], bgetClis.DownloadDir, netOpt)
				}
			}(k, i)
		}
	}
	for i := 0; i < cap(sem); i++ {
		sem <- true
	}
	sem = make(chan bool, bgetClis.Thread)
	for k, v := range seqs {
		for i := range v {
			sem <- true
			done[k] = append(done[k], seqs[k][i])
			go func(k string, i int) {
				defer func() {
					<-sem
				}()
				if k == "geo" {
					spider.Geofetch(seqs[k][i], bgetClis.DownloadDir,
						bgetClis.GeoGPL, bgetClis.Uncompress, netOpt)
				}
			}(k, i)
		}
	}
	for i := 0; i < cap(sem); i++ {
		sem <- true
	}
}

func init() {
	SeqCmd.Flags().BoolVarP(&(bgetClis.GeoGPL), "query-gpl", "", false, "Wheather fetch GPL files from GEO database.")
	SeqCmd.Flags().StringVarP(&(bgetClis.GdcToken), "token-gdc", "", "", "Token to access TCGA portal files.")
	SeqCmd.Flags().StringVarP(&(bgetClis.EgaCredFile), "token-file-ega", "", "", `Credential file to access EGA archive files, {"username": "{your_user_name}", 
  "password": "{your_password}","client_secret":"AMenuDLjVdVo4BSwi0QD54LL6NeVDEZRzEQUJ7h
	JOM3g4imDZBHHX0hNfKHPeQIGkskhtCmqAJtt_jm7EKq-rWw"}.`)
	setGlobalFlag(SeqCmd, &bgetClis)
	setKeyListFlag(SeqCmd, &bgetClis, "accession ids")
	SeqCmd.Example = `  bget seq ERR3324530 SRR544879 # download files from SRA databaes
  bget seq GSE23543 GSM1098572 -t 2 # download files from GEO databaes (auto download SRA acc list and run info)
  bget seq dbgap.krt # download files from dbGap database using krt files
  bget seq EGAD00001000951 # download dataset from EGA databaes
  bget seq EGAF00000585895 # download file from EGA databaes
	
  # download TCGA files using file id
  bget seq b7670817-9d6b-494e-9e22-8494e2fd430d

  # download TCGA files using manifest files
  # split for parallel
  split -a 3 --additional-suffix=.txt -l 100 gdc_manifest.2019-08-23-TCGA.txt -d
  for i in x*.txt
  do
    head -n 1 x000.txt > ${i}.tmp && cat ${i} >> ${i}.tmp &&mv ${i}.tmp ${i}
  done
  sed -i '1d' x000.txt
  bget seq *.txt -t 5

  # support auto (if you do not have *.krt, TCGA manifest, please not include it for test)
  bget seq SRR544879 GSE23543 EGAD00001000951 b7670817-9d6b-494e-9e22-8494e2fd430d dbgap.krt *.txt -t 5`
}
