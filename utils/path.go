package utils

import (
	"fmt"
	"os"
	"path"

	"github.com/JhuangLab/bioget/log"
)

// PathExists check wheather file is existed
func PathExists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// LinkFileToDir create os.Symlink
func LinkFileToDir(oldname, newname string) bool {
	hasTargetFile, _ := PathExists(newname)
	if !hasTargetFile {
		if err := CreateFileParDir(newname); err != nil {
			log.Fatalf("Can not create output directory %s", path.Dir(newname))
			return false
		}
		err := os.Symlink(oldname, newname)
		if err != nil {
			log.Fatalf("Can not create soft link of %s to %s", oldname, newname)
			return false
		}
	}
	return true
}

//CheckFilesOverwrite check wheather to overwrite existing files
func CheckFilesOverwrite(files []string, overwrite bool) int {
	for _, v := range files {
		hasFile, _ := PathExists(v)
		if hasFile && overwrite {
			if err := os.Remove(v); err != nil {
				log.Fatalf("Can not overwrite the output file %s.", v)
				return 1
			}
		} else if hasFile {
			log.Infof("%s existed.", v)
			return 0
		}
	}
	return -1
}

//CheckRenameOverwrite check wheather to rename existing files
func CheckRenameOverwrite(old string, new string, overwrite bool) int {
	hasNew, _ := PathExists(new)
	hasOld, _ := PathExists(old)
	if hasNew && !overwrite {
		return 0
	} else if hasOld && !overwrite {
		log.Infoln(fmt.Sprintf("%s existed.", old))
		err := os.Rename(old, new)
		log.LogBash.Infof("mv %s %s", old, new)
		if err != nil {
			log.Fatalln(err)
			return 1
		}
	}
	return -1
}

// CheckFilesFinal batch to check wheathre final files is existed
func CheckFilesFinal(files []string) bool {
	for _, v := range files {
		if hasFile, _ := PathExists(v); hasFile {
			return true
		}
	}
	return false
}

// CheckRenameFinal check final file wheather renamed
func CheckRenameFinal(old string, new string) bool {
	hasNew, _ := PathExists(new)
	hasOld, _ := PathExists(old)
	if !hasNew && hasOld {
		log.Infoln(fmt.Sprintf("%s existed.", old))
		err := os.Rename(old, new)
		log.LogBash.Infof("mv %s %s", old, new)
		if err != nil {
			log.Fatalln(err)
			return false
		}
	}
	return true
}

// CheckInputFiles check input files
func CheckInputFiles(files []string) error {
	for _, v := range files {
		if v != "" {
			hasFile, _ := PathExists(v)
			var err error
			if !hasFile {
				err = fmt.Errorf(fmt.Sprintf("Input file %s not existed.", v))
				return err
			}
		}
	}
	return nil
}
