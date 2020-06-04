package urlpool

import (
	"os"
	"path"

	"github.com/openbiox/ligo/archive"
	stringo "github.com/openbiox/ligo/stringo"
)

func PostKeyCmds(key string, dest []string, rawKey string) {
	if key == "reffa/defuse" {
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
			err = archive.UnarchiveLog(srcFn, destDir)
		} else if stringo.StrDetect(fn, ".fa.gz$") && stringo.StrDetect(fn, ".dna.") {
			srcFn = stringo.StrReplaceAll(fn, "Homo_sapiens.*.dna.chromosome.", "defuse.dna.chromosomes.")
			log.Infof("Link %s => %s", fn, srcFn)
			err = os.Symlink(fn, srcFn)
			if err != nil {
				log.Warn(err)
			}
		} else if stringo.StrDetect(fn, ".gz$") {
			srcFn = fn
			err = archive.UnarchiveLog(srcFn, destDir)
		}
		if err != nil {
			log.Warn(err)
		}
		errList = append(errList, err)
	}
	return errList
}
