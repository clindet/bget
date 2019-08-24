package cmd

import (
	"fmt"
	"net/url"
	"os/exec"
	"path"
	"strings"

	"github.com/JhuangLab/bget/urlpool"
	vers "github.com/JhuangLab/bget/versions"
	butils "github.com/JhuangLab/butils"
	log "github.com/JhuangLab/butils/log"
	"github.com/spf13/cobra"
)

var keyVs map[string][]string

var keyCmd = &cobra.Command{
	Use:   "url-key [key1 key2 key3...]",
	Short: "Can be used to access URLs via a key string.",
	Long:  `Can be used to access URLs via a key string. e.g. 'item' or 'item@version %site #releaseVersion', : bwa, GRCh38 %defuse #97. More see here https://github.com/JhuangLab/bget.`,
	Run: func(cmd *cobra.Command, args []string) {
		keyCmdRunOptions(cmd)
	},
}

func parseKeys() (keys []string) {
	if bgetClis.keys != "" && strings.Contains(bgetClis.keys, bgetClis.separator) {
		keys = strings.Split(bgetClis.keys, bgetClis.separator)
	} else if bgetClis.keys != "" {
		keys = []string{bgetClis.keys}
	} else if bgetClis.listFile != "" {
		keys = butils.ReadLines(bgetClis.listFile)
	}
	return keys
}

func downloadKey() {
	keys := parseKeys()
	urls, postShellCmd := vers.QueryKeysInfo(keys, osType)
	done := make(map[string][]string)
	var destDirArray []string
	sem := make(chan bool, bgetClis.concurrency)
	for key, v := range urls {
		for i := range v {
			u, _ := url.Parse(v[i])
			v[i] = strings.TrimSpace(u.String())
			if bgetClis.autoPath {
				destDirArray = append(destDirArray, path.Join(bgetClis.downloadDir, key))
			} else {
				destDirArray = append(destDirArray, bgetClis.downloadDir)
			}
		}
		sem <- true
		go func(key string, v []string, destDirArray []string) {
			defer func() {
				<-sem
			}()
			done[key] = HTTPGetURLs(v, destDirArray, bgetClis.engine, cmdExtraFromFlag, taskID, bgetClis.mirror,
				bgetClis.concurrency, bgetClis.axelThread, overwrite, ignore, quiet, saveLog)
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
			if bgetClis.uncompress {
				if err := butils.UnarchiveLog(dest, path.Dir(dest)); err != nil {
					log.Warn(err)
				}
			}
			if args == "" {
				continue
			}
			cmd := exec.Command("sh", "-c", args)
			logFn := path.Join(logDir, taskID+"_postShellCmd_"+key+".log")
			if args != "" {
				butils.RunExecCmdConsole(logFn, cmd, quiet, saveLog)
			}
		}
		urlpool.PostKeyCmds(key, done[key], bgetClis.keys)
	}
}

func postCmdRender(oldCmd string, dest string) (newCmd string) {
	if hasDest, _ := butils.PathExists(dest); !hasDest {
		return ""
	}
	if cmdExtraFromFlag != "" {
		newCmd = oldCmd + " " + cmdExtraFromFlag
	}
	// define your pattern replace
	newCmd = strings.Replace(oldCmd, "{{downloadDir}}", bgetClis.downloadDir, 100)
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
	keys = butils.RemoveRepeatEle(keys)
	fmt.Printf("%s\n", strings.Join(keys, "\n"))
	return keys
}

func keyCmdRunOptions(cmd *cobra.Command) {
	checkQuiet()
	items := []string{}
	if len(cmd.Flags().Args()) >= 1 {
		items = append(items, cmd.Flags().Args()...)
		bgetClis.keys = strings.Join(items, bgetClis.separator)
	}
	if bgetClis.getKeyVersions != "" {
		log.Infoln("Featching versions from local and remote website...")
		keys := parseKeys()
		keyVs = vers.QueryKeysVersions(keys, osType, bgetClis.getKeyVersions)
		return
	}
	if bgetClis.keysAll {
		getAllKeys()
		bgetClis.helpFlags = false
		return
	}
	checkDownloadDir(bgetClis.keys != "")
	if bgetClis.keys != "" || bgetClis.listFile != "" {
		downloadKey()
		bgetClis.helpFlags = false
	}
	if bgetClis.helpFlags {
		cmd.Help()
	}
}

func init() {
	keyCmd.PersistentFlags().BoolVar(&(bgetClis.autoPath), "autopath", false, "Logical indicating that whether to create subdir in download dir: e.g. reffa/{{key}}/")
	keyCmd.Flags().StringVarP(&(bgetClis.engine), "engine", "g", "go-http", "Point the download engine: go-http, wget, curl, axel, git, and rsync.")
	keyCmd.Flags().IntVarP(&(bgetClis.axelThread), "thread-axel", "", 5, "Set the thread of axel.")
	keyCmd.Flags().StringVarP(&(bgetClis.mirror), "mirror", "m", "", "Set the mirror of resources.")
	keyCmd.Flags().StringVarP(&(bgetClis.getKeyVersions), "show-versions", "v", "", "Show all available versions of key. Optional (txt, json, table)")
	keyCmd.Flags().StringVarP(&(bgetClis.listFile), "list-file", "l", "", "A file contains keys for download.")
	keyCmd.Flags().BoolVarP(&(bgetClis.keysAll), "keys-all", "a", false, "Show all available string key can be download.")
	keyCmd.Flags().BoolVarP(&(bgetClis.uncompress), "uncompress", "u", false, "Uncompress download files for .zip, .tar.gz, and .gz suffix files.")
	keyCmd.Example = `  bget url-key bwa
  bget url-key "reffa@GRCh38 %defuse #97" -t 10 -f`
}
