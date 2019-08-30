package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	cio "github.com/openbiox/butils/io"
	"github.com/openbiox/butils/log"
)

func checkQuiet() {
	if bgetClis.quiet {
		log.SetOutput(ioutil.Discard)
	} else {
		log.SetOutput(os.Stderr)
	}
}

func checkDownloadDir(condtions bool) {
	if hasDir, _ := cio.PathExists(bgetClis.downloadDir); !hasDir {
		if condtions {
			if err := cio.CreateDir(bgetClis.downloadDir); err != nil {
				log.FATAL(fmt.Sprintf("Could not to create %s", bgetClis.downloadDir))
			}
		}
	}
}
