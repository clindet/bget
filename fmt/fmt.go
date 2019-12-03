package fmt

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"

	jsoniter "github.com/json-iterator/go"
	"github.com/openbiox/butils/log"
	stringo "github.com/openbiox/butils/stringo"
	"github.com/tidwall/pretty"
)

type FmtClisT struct {
	Stdin       *[]byte
	Files       *[]string
	JSON        *map[int]map[string]interface{}
	Table       *map[int][]interface{}
	PrettyJSON  bool
	JSONToSlice bool
	JSONToCSV   bool
	Indent      int
	SortKeys    bool
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func PrettyJSON(fmtClis *FmtClisT, thread int) {
	sem := make(chan bool, thread)
	var m map[string]interface{}
	m = make(map[string]interface{})
	var d []byte
	indent := ""
	for i := 0; i < fmtClis.Indent; i++ {
		indent = indent + " "
	}
	opt := pretty.Options{
		Indent:   indent,
		SortKeys: fmtClis.SortKeys,
	}
	for k, fn := range *fmtClis.Files {
		sem <- true
		go func(fn string, k int) {
			defer func() {
				<-sem
			}()
			outfn := stringo.StrReplaceAll(fn, "json$", "pretty.json")
			(*fmtClis.Files)[k] = outfn
			d, err := ioutil.ReadFile(fn)
			if err != nil {
				log.Fatal(err)
			}
			d = pretty.PrettyOptions(d, &opt)
			json.Unmarshal(d, &m)
			(*fmtClis.JSON)[k] = m
			f, err := os.OpenFile(outfn, os.O_RDWR|os.O_CREATE, 0664)
			io.Copy(f, bytes.NewBuffer(d))
		}(fn, k)
	}
	for i := 0; i < cap(sem); i++ {
		sem <- true
	}
	if fmtClis.Stdin != nil {
		var m2 map[string]interface{}
		m2 = make(map[string]interface{})
		json.Unmarshal(*fmtClis.Stdin, &m2)
		d = pretty.PrettyOptions(*fmtClis.Stdin, &opt)
		(*fmtClis.JSON)[-1] = m2
		io.Copy(os.Stdout, bytes.NewBuffer(d))
	}
}

func JSON2Slice(fmtClis *FmtClisT, thread int) {
	sem := make(chan bool, thread)
	var final []interface{}
	var j string
	indent := ""
	for i := 0; i < fmtClis.Indent; i++ {
		indent = indent + " "
	}
	opt := pretty.Options{
		Indent:   indent,
		SortKeys: fmtClis.SortKeys,
	}
	for k, fn := range *(fmtClis.Files) {
		sem <- true
		go func(fn string, k int) {
			defer func() {
				<-sem
			}()
			var m map[string]interface{}
			m = make(map[string]interface{})
			outfn := stringo.StrReplaceAll(fn, "json$", "slice.json")
			(*fmtClis.Files)[k] = outfn
			var d []byte
			var err error
			if len((*fmtClis.JSON)[k]) > 0 {
				m = (*fmtClis.JSON)[k]
			} else {
				d, err = ioutil.ReadFile(fn)
				if err != nil {
					log.Fatal(err)
				}
			}
			json.Unmarshal(d, &m)

			for j = range m {
				final = append(final, m[j])
			}
			(*fmtClis.Table)[k] = final
			if j != "" {
				d, _ = json.Marshal(final)
				d = pretty.PrettyOptions(d, &opt)
				f, err := os.OpenFile(outfn, os.O_RDWR|os.O_CREATE, 0664)
				if err != nil {
					log.Fatal(err)
				}
				defer f.Close()
				io.Copy(f, bytes.NewBuffer(d))
			}
		}(fn, k)
	}
	for i := 0; i < cap(sem); i++ {
		sem <- true
	}
	if fmtClis.Stdin != nil {
		var m2 map[string]interface{}
		m2 = make(map[string]interface{})
		json.Unmarshal(*fmtClis.Stdin, &m2)
		var final []interface{}
		var j string
		for j = range m2 {
			final = append(final, m2[j])
		}
		(*fmtClis.Table)[-1] = final
		if j != "" {
			d, _ := json.Marshal(final)
			d = pretty.PrettyOptions(d, &opt)
			io.Copy(os.Stdout, bytes.NewBuffer(d))
		}
	}
}
