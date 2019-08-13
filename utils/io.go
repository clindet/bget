package utils

import (
	"bufio"
	"io"
	"os"
	"path"
	"strings"

	log "github.com/JhuangLab/bioget/log"
)

// ConnectFile open connect to file using os.OpenFile (auto create if not exist the file)
func ConnectFile(name string) (*os.File, error) {
	var fn *os.File
	var err error
	if hasFile, _ := PathExists(name); !hasFile {
		CreateFile(name)
	}
	fn, err = os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0664)
	return fn, err
}

// CreateFile create file
func CreateFile(name string) error {
	fo, err := os.Create(name)
	if err != nil {
		return err
	}
	defer func() {
		fo.Close()
	}()
	return nil
}

// CreateFileParDir create parant dir of file
func CreateFileParDir(name string) error {
	var err error
	if isExists, _ := PathExists(path.Dir(name)); !isExists {
		err = os.MkdirAll(path.Dir(name), os.ModePerm)
	}
	return err
}

// CreateDir create dir with os.MkdirAll
func CreateDir(name string) error {
	var err error
	if isExists, _ := PathExists(name); !isExists {
		err = os.MkdirAll(name, os.ModePerm)
	}
	return err
}

// ReadLines import a file as []string
func ReadLines(fn string) []string {
	var final []string
	if hasFile, _ := PathExists(fn); !hasFile {
		log.Fatalf("%s not existed.", fn)
	}
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	buf := bufio.NewReader(f)
	for {
		b, errR := buf.ReadBytes('\n')
		if errR != nil {
			if errR == io.EOF {
				break
			}
			log.Fatal(errR.Error())
		}
		final = append(final, strings.ReplaceAll(string(b), "\n", ""))
	}
	return final
}
