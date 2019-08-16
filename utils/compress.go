package utils

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"

	log "github.com/JhuangLab/bget/log"
)

// TarGz compress files and get .tar.gz file
func TarGz(files []*os.File, dest string) error {
	d, _ := os.Create(dest)
	defer d.Close()
	gw := gzip.NewWriter(d)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	for _, file := range files {
		err := compress(file, "", tw)
		if err != nil {
			return err
		}
	}
	return nil
}

func compress(file *os.File, prefix string, tw *tar.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, tw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := tar.FileInfoHeader(info, "")
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		err = tw.WriteHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(tw, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

// UnTarGz uncompress .tar.gz files
func UnTarGz(tarFile, dest string) error {
	srcFile, err := os.Open(tarFile)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		filename := dest + hdr.Name
		file, err := createFile(filename)
		if err != nil {
			return err
		}
		io.Copy(file, tr)
	}
	return nil
}

// Gunzip use compress/gzip to gunzip files
func Gunzip(inputFn, outFn string) (int64, error) {
	gzFile, err := os.Open(inputFn)
	if err != nil {
		return 0, fmt.Errorf("Failed to open file %s for unpack: %s", inputFn, err)
	}
	dstFile, err := os.OpenFile(outFn, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		return 0, fmt.Errorf("Failed to create destination file %s for unpack: %s", outFn, err)
	}

	ioReader, ioWriter := io.Pipe()

	go func() { // goroutine leak is possible here
		gzReader, _ := gzip.NewReader(gzFile)
		// it is important to close the writer or reading from the other end of the
		// pipe or io.copy() will never finish
		defer func() {
			gzFile.Close()
			gzReader.Close()
			ioWriter.Close()
		}()

		io.Copy(ioWriter, gzReader)
	}()

	written, err := io.Copy(dstFile, ioReader)
	if err != nil {
		return 0, err // goroutine leak is possible here
	}
	ioReader.Close()
	dstFile.Close()

	return written, nil
}

func createFile(name string) (*os.File, error) {
	err := os.MkdirAll(string([]rune(name)[0:strings.LastIndex(name, "/")]), 0755)
	if err != nil {
		return nil, err
	}
	return os.Create(name)
}

// GunzipDefuseReffa gunzip and rename defuse reference dir
func GunzipDefuseReffa(destDir string) {
	dirList, e := ioutil.ReadDir(destDir)
	if e != nil {
		fmt.Println("read dir error")
		return
	}
	for _, v := range dirList {
		destFn := ""
		oldPath := path.Join(destDir, v.Name())
		if path.Base(v.Name()) == "rmsk.txt.gz" {
			destFn = path.Join(destDir, "repeats.txt.gz")
			log.Infof("Rename %s => %s", oldPath, destFn)
			err := os.Rename(oldPath, destFn)
			if err != nil {
				log.Warn(err)
			}
			destFn2 := StrReplaceAll(destFn, ".gz$", "")
			GunzipLog(destFn, destFn2)
		} else if StrDetect(v.Name(), ".fa.gz$") && StrDetect(v.Name(), ".dna.") {
			destFn = StrReplaceAll(v.Name(), ".fa.gz$", ".fa")
			destFn = StrReplaceAll(destFn, "Homo_sapiens.*.dna.chromosome.", "defuse.dna.chromosomes.")
			destFn = path.Join(destDir, destFn)
			GunzipLog(oldPath, destFn)
		} else if StrDetect(v.Name(), ".gz$") {
			destFn = StrReplaceAll(v.Name(), ".gz$", "")
			destFn = path.Join(destDir, destFn)
			GunzipLog(oldPath, destFn)
		}
	}
}

// GunzipLog gunzip files with "Uncompressing =>" log
func GunzipLog(oldPath string, destFn string) {
	if hasFn, _ := PathExists(destFn); !hasFn {
		log.Infof("Uncompressing %s => %s", oldPath, destFn)
		_, err := Gunzip(oldPath, destFn)
		if err != nil {
			log.Warn(err)
		}
	} else {
		log.Infof("%s existed.", destFn)
	}
}
