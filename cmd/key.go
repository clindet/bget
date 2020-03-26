package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path"
	"sort"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/openbiox/bget/spider"
	"github.com/openbiox/bget/urlpool"
	vers "github.com/openbiox/bget/versions"
	"github.com/openbiox/ligo/archive"
	bexec "github.com/openbiox/ligo/exec"
	cio "github.com/openbiox/ligo/io"
	cnet "github.com/openbiox/ligo/net"
	"github.com/openbiox/ligo/slice"
	"github.com/spf13/cobra"
)

var keyVs map[string][]string

// KeyCmd is the cobra command object to run bget key
var KeyCmd = &cobra.Command{
	Use:   "key [key1 key2 key3...]",
	Short: "Can be used to access URLs via a key string.",
	Long:  `Can be used to access URLs via a key string. e.g. 'item' or 'item@version #releaseVersion', : bwa, reffa-defuse@GRCh38 #97. More see here https://github.com/openbiox/bget.`,
	Run: func(cmd *cobra.Command, args []string) {
		keyCmdRunOptions(cmd, args)
	},
}

func keyCmdRunOptions(cmd *cobra.Command, args []string) {
	initCmd(cmd, args)
	checkArgs(cmd, "key")
	if bgetClis.ShowVersions {
		log.Infoln("Featching versions from local and remote website...")
		keys := parseKeys()
		if bgetClis.PrintFormat == "" {
			bgetClis.PrintFormat = "table"
		}
		bgetClis.Env["PrintFormat"] = bgetClis.PrintFormat
		keyVs = vers.QueryKeysVersions(keys, &bgetClis.Env)
		return
	}
	if bgetClis.KeysAll {
		getAllKeys()
		bgetClis.HelpFlags = false
		return
	}
	checkDownloadDir(bgetClis.Keys != "")
	if bgetClis.Keys != "" || bgetClis.ListFile != "" {
		downloadKey()
		bgetClis.HelpFlags = false
	}
	if bgetClis.HelpFlags {
		cmd.Help()
	}
}

func parseKeys() (keys []string) {
	if bgetClis.Keys != "" && strings.Contains(bgetClis.Keys, bgetClis.Seperator) {
		keys = strings.Split(bgetClis.Keys, bgetClis.Seperator)
	} else if bgetClis.Keys != "" {
		keys = []string{bgetClis.Keys}
	} else if bgetClis.ListFile != "" {
		keys = cio.ReadLines(bgetClis.ListFile)
	}
	return keys
}

func downloadKey() {
	keys := parseKeys()
	urls, postShellCmd, _ := vers.QueryKeysInfo(keys, &bgetClis.Env)
	done := make(map[string][]string)
	var destDirArray []string
	netOpt := setNetParams(&bgetClis)
	sem := make(chan bool, bgetClis.Thread)
	for key, v := range urls {
		for i := range v {
			v[i] = preURLFilter(v[i])
			u, _ := url.Parse(v[i])
			v[i] = strings.TrimSpace(u.String())
			assetsDir := ""
			if strings.Contains(v[i], "github.com") && strings.Contains(v[i], "/releases/download/") {
				assetsDir = "github-assets"
			}
			if bgetClis.AutoPath {
				destDirArray = append(destDirArray, path.Join(bgetClis.DownloadDir, key, assetsDir))
			} else {
				destDirArray = append(destDirArray, path.Join(bgetClis.DownloadDir, assetsDir))
			}
		}
		sem <- true
		go func(key string, v []string, destDirArray []string) {
			defer func() {
				<-sem
			}()
			done[key] = cnet.HTTPGetURLs(v, destDirArray, netOpt)
		}(key, v, destDirArray)
	}
	for i := 0; i < cap(sem); i++ {
		sem <- true
	}
	dest := ""
	for key := range done {
		for i := range done[key] {
			args := ""
			dest = done[key][i]
			if len(postShellCmd[key]) > i {
				args = postShellCmd[key][i]
				args = postCmdRender(args, dest)
			}

			if bgetClis.Uncompress {
				if err := archive.UnarchiveLog(dest, path.Dir(dest)); err != nil {
					log.Warn(err)
				}
			}
			if args == "" {
				continue
			}
			cmd := exec.Command("sh", "-c", args)
			logPath := ""
			if bgetClis.SaveLog {
				logPath = path.Join(bgetClis.LogDir, bgetClis.TaskID+"_postShellCmd_"+strings.ReplaceAll(key, "/", "_")+".log")
			}
			if args != "" {
				err := bexec.System(cmd, logPath, bgetClis.Verbose == 0)
				if err != nil {
					log.Warn(err)
				}
			}
		}
		urlpool.PostKeyCmds(key, done[key], bgetClis.Keys)
	}
}

func preURLFilter(url string) string {
	if strings.Contains(url, "doaj.org") {
		url, err := spider.RetriveRedirectLink(url, bgetClis.Timeout, bgetClis.Proxy)
		if err != nil {
			log.Warnln(err)
			return url
		}
	}
	return url
}

func postCmdRender(oldCmd string, dest string) (newCmd string) {
	if hasDest, _ := cio.PathExists(dest); !hasDest {
		return ""
	}
	if bgetClis.CmdExtraFromFlag != "" {
		newCmd = oldCmd + " " + bgetClis.CmdExtraFromFlag
	}
	// define your pattern replace
	newCmd = strings.Replace(oldCmd, "{{downloadDir}}", bgetClis.DownloadDir, 100)
	newCmd = strings.Replace(newCmd, "{{dest}}", dest, 100)
	newCmd = strings.Replace(newCmd, "{{pdir}}", path.Dir(dest), 100)
	return newCmd
}

func getAllKeys() (keys []string) {
	for i := range urlpool.BgetToolsPool {
		keys = append(keys, urlpool.BgetToolsPool[i].Name)
	}
	for i := range urlpool.BgetFilesPool {
		keys = append(keys, urlpool.BgetFilesPool[i].Name)
	}
	keys = slice.DropSliceDup(keys)
	sort.Strings(keys)
	if bgetClis.PrintFormat == "" || bgetClis.PrintFormat == "table" {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetRowLine(false)
		table.SetRowSeparator("")
		table.SetAlignment(tablewriter.ALIGN_LEFT)

		tmp := []string{}
		for i := range keys {
			if i%2 == 0 {
				table.Append(tmp)
				tmp = []string{}
			}
			keys[i] = strings.ToLower(keys[i])
			keys[i] = strings.ReplaceAll(keys[i], "_", "-")
			tmp = append(tmp, keys[i])
		}
		if len(tmp) != 0 {
			table.Append(tmp)
		}
		table.Render()
	} else if bgetClis.PrintFormat == "text" {
		fmt.Println(strings.Join(keys, "\n"))
	} else if bgetClis.PrintFormat == "json" {
		var str bytes.Buffer
		js := fmt.Sprintf(`{"keys":[ "%s" ]}`, strings.Join(keys, `", "`))
		_ = json.Indent(&str, []byte(js), "", "  ")
		fmt.Println(str.String())
	}

	//fmt.Printf("%s\n", strings.Join(keys, "\n"))
	return keys
}

func init() {
	KeyCmd.Flags().BoolVar(&(bgetClis.AutoPath), "autopath", false, "Logical indicating that whether to create subdir in download dir: e.g. reffa/{{key}}/")
	KeyCmd.Flags().BoolVarP(&(bgetClis.ShowVersions), "show-versions", "v", false, "Show all available versions of key.")
	KeyCmd.Flags().StringVarP(&(bgetClis.PrintFormat), "format", "", "", "Output format (text, json, table)")
	KeyCmd.Flags().BoolVarP(&(bgetClis.KeysAll), "keys-all", "a", false, "Show all available string key can be download.")
	KeyCmd.Flags().BoolVarP(&(bgetClis.WithAssets), "with-assets", "", false, "Logical indicating that whether to download associated assets files.")
	setGlobalFlag(KeyCmd, &bgetClis)
	setUncompressFlag(KeyCmd, &bgetClis)
	setKeyListFlag(KeyCmd, &bgetClis, "keys")
	KeyCmd.Example = `  # download bwa source (with task env info)
  bget key bwa --verbose 2
  # get all available keys
  bget key -a
  # in JSON format
  bget key -a --format json
  # view all bwa and samtools available tags in table
  bget key bwa samtools -v
  # view all bwa and samtools available tags in json
  bget key bwa samtools -v --format json
	
  # force download defuse reference (with task env info and save log to file)
  bget key "reffa/defuse@GRCh38 #97" -t 10 -f --verbose 2 --save-log
  bget key reffa/defuse@GRCh38 release=97 -t 10 -f
  # download annovar reference
  bget key db/annovar@clinvar_20170501 db/annovar@clinvar_20180603 builder=hg38

  bget key db/annovar -v --formt text
  bget key db/annovar version='clinvar_20131105, clinvar_20140211, clinvar_20140303, clinvar_20140702, clinvar_20140902, clinvar_20140929, clinvar_20150330, clinvar_20150629, clinvar_20151201, clinvar_20160302, clinvar_20161128, clinvar_20170130, clinvar_20170501, clinvar_20170905, clinvar_20180603, avsnp150, avsnp147, avsnp144, avsnp142, avsnp138, cadd, caddgt10, caddgt20, cadd13, cadd13gt10, cadd13gt20, cg69, cg46, cosmic70, cosmic68wgs, cosmic68, cosmic67wgs, cosmic67, cosmic65, cosmic64, dbnsfp35a, dbnsfp33a, dbnsfp31a_interpro, dbnsfp30a, dbscsnv11, eigen, esp6500siv2_ea, esp6500siv2_aa, esp6500siv2_all, exac03nontcga, exac03nonpsych, exac03, fathmm, gerp++gt2, gme, gnomad_exome, gnomad_genome, gwava, hrcr1, icgc21, intervar_20170202, kaviar_20150923, ljb26_all, mcap, mitimpact2, mitimpact24, nci60, popfreq_max_20150413, popfreq_all_20150413, revel, regsnpintron' builder=hg19 -t 10 -f`
}
