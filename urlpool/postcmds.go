package urlpool

import (
	"os"
	"path"
	"strings"

	butils "github.com/openbiox/butils"
	log "github.com/openbiox/butils/log"
)

func PostKeyCmds(key string, dest []string, rawKey string) {
	if key == "reffa" && strings.Contains(rawKey, "defuse") {
		gunzipDefuseReffa(dest)
	}
}

// gunzipDefuseReffa gunzip and rename defuse reference dir
func gunzipDefuseReffa(dest []string) (errList []error) {
	var err error
	for _, fn := range dest {
		destDir := path.Dir(fn)
		srcFn := ""
		if path.Base(fn) == "rmsk.txt.gz" {
			srcFn = path.Join(destDir, "repeats.txt.gz")
			log.Infof("Link %s => %s", fn, srcFn)
			err = os.Symlink(fn, srcFn)
			if err != nil {
				log.Warn(err)
			}
			err = butils.UnarchiveLog(srcFn, destDir)
		} else if butils.StrDetect(fn, ".fa.gz$") && butils.StrDetect(fn, ".dna.") {
			srcFn = butils.StrReplaceAll(fn, "Homo_sapiens.*.dna.chromosome.", "defuse.dna.chromosomes.")
			log.Infof("Link %s => %s", fn, srcFn)
			err = os.Symlink(fn, srcFn)
			if err != nil {
				log.Warn(err)
			}
			err = butils.UnarchiveLog(srcFn, destDir)
		} else if butils.StrDetect(fn, ".gz$") {
			srcFn = fn
			err = butils.UnarchiveLog(srcFn, destDir)
		}
		if err != nil {
			log.Warn(err)
		}
		errList = append(errList, err)
	}
	return errList
}
