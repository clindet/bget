package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/JhuangLab/bget/log"
	"github.com/JhuangLab/bget/utils"
)

func installSpack() {
	if overwrite {
		os.RemoveAll(path.Join(downloadClis.downloadDir, "spack"))
	}
	AsyncURL(getEnvToolsURL("spack", "latest", ""), path.Join(downloadClis.downloadDir, "spack"),
		"git", taskID, downloadClis.mirror, 1, quiet, saveLog)
}

func installMiniconda() {
	url := getEnvToolsURL("miniconda", "latest", downloadClis.installConda)
	destFn := path.Join(os.TempDir(), path.Base(url))
	destDir := ""
	destDir = path.Join(downloadClis.downloadDir, "miniconda"+downloadClis.installConda)

	if overwrite {
		os.RemoveAll(destDir)
		os.RemoveAll(destFn)
	}
	logPath := path.Join(logDir, fmt.Sprintf("%s_installMiniconda.log", taskID))
	if saveLog {
		utils.CreateFileParDir(logPath)
	}
	if status, _ := utils.PathExists(destFn); !status {
		log.Infof("Trying %s => %s", url, destFn)
		AsyncURL3(url, destFn, downloadClis.engine, taskID, downloadClis.mirror, downloadClis.axelThread, quiet, saveLog)
	} else {
		log.Infof("%s existed.", destFn)
	}

	if osType != "windows" {
		args := []string{destFn, "-b", "-p", destDir, cmdExtraFromFlag}
		cmd := exec.Command("sh", args...)
		utils.RunExecCmdConsole(logPath, cmd, quiet, saveLog)
	}
}
