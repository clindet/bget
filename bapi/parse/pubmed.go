package parse

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	jsoniter "github.com/json-iterator/go"
	"github.com/openbiox/butils/log"
	"github.com/openbiox/butils/slice"
	"github.com/openbiox/butils/stringo"
	prose "gopkg.in/jdkato/prose.v2"
	xurls "mvdan.cc/xurls/v2"
)

type PubmedFields struct {
	Pmid, Doi, Title, Abs, Journal, Issue, Volume, Date, Issn *string
	Corelations                                               *map[string]string
	URLs                                                      *[]string
	Keywords                                                  *[]string
}

// PubmedXML convert Pubmed XML to json
func PubmedXML(xmlPaths *[]string, stdin *[]byte, outfn string, keywords *[]string, thread int, callCor bool) {
	if len(*xmlPaths) == 1 {
		thread = 1
	}
	if len(*stdin) > 0 {
		*xmlPaths = append(*xmlPaths, "ParsePubmedXMLStdin")
	}
	sem := make(chan bool, thread)

	//|os.O_APPEND
	var of *os.File
	if outfn == "" {
		of = os.Stdout
	} else {
		of, err := os.OpenFile(outfn, os.O_CREATE|os.O_WRONLY, 0664)
		if err != nil {
			log.Fatal(err)
		}
		defer of.Close()
	}

	var err error
	for i, xmlPath := range *xmlPaths {
		sem <- true
		go func(xmlPath string, i int) {
			var htmlDoc *goquery.Document
			defer func() {
				<-sem
			}()
			if xmlPath != "ParsePubmedXMLStdin" {
				xml, err := os.Open(xmlPath)
				if err != nil {
					log.Fatal(err)
				}
				defer xml.Close()
				htmlDoc, err = goquery.NewDocumentFromReader(xml)
				if err != nil {
					log.Fatal(err)
				}
			} else if xmlPath == "ParsePubmedXMLStdin" && len(*stdin) > 0 {
				htmlDoc, err = goquery.NewDocumentFromReader(bytes.NewReader(*stdin))
				if err != nil {
					log.Fatal(err)
				}
			}

			htmlDoc.Find("PubmedArticle").Each(func(i int, s *goquery.Selection) {
				io.Copy(of, bytes.NewBuffer(getPubmedFields(keywords, s, callCor)))
				// fmt.Printf("%s, %s\n%s\n%v\n", pmid, doi, abs, key)
			})
		}(xmlPath, i)
	}
	for i := 0; i < cap(sem); i++ {
		sem <- true
	}
}

func getPubmedFields(keywords *[]string, s *goquery.Selection, callCor bool) (jsonData []byte) {
	year := s.Find("PubmedArticle MedlineCitation Article Journal JournalIssue PubDate > Year").Text()
	month := s.Find("PubmedArticle MedlineCitation Article Journal JournalIssue PubDate > Month").Text()
	day := s.Find("PubmedArticle MedlineCitation Article Journal JournalIssue PubDate > Day").Text()
	date := fmt.Sprintf("%s %s %s", year, month, day)
	issue := s.Find("PubmedArticle MedlineCitation Article Journal JournalIssue Issue").Text()
	volume := s.Find("PubmedArticle MedlineCitation Article Journal JournalIssue Volume").Text()
	journal := s.Find("PubmedArticle MedlineCitation Article Journal ISOAbbreviation").Text()
	issn := s.Find("PubmedArticle MedlineCitation Article Journal ISSN").Text()
	pmid := s.Find("PubmedArticle PubmedData > ArticleIdList > ArticleId[IdType=pubmed]").Text()
	doi := s.Find("PubmedArticle PubmedData > ArticleIdList > ArticleId[IdType=doi]").Text()
	abs := s.Find("PubmedArticle MedlineCitation Article Abstract").Text()
	abs = stringo.StrReplaceAll(abs, "\n  *", "\n")
	abs = stringo.StrReplaceAll(abs, "(<[/]AbstractText.*>)|(^\n)|(\n$)", "")
	title := s.Find("PubmedArticle MedlineCitation Article ArticleTitle").Text()
	titleAbs := title + "\n" + abs
	urls := xurls.Relaxed().FindAllString(titleAbs, -1)
	keywordsPat := strings.Join(*keywords, "|")
	key := stringo.StrExtract(titleAbs, keywordsPat, 1000000)
	key = slice.DropSliceDup(key)

	doc, err := prose.NewDocument(titleAbs)
	corela := make(map[string]string)
	if len(key) > 2 && callCor {
		getKeywordsCorleations(doc, &keywordsPat, &corela)
	}
	if err != nil {
		log.Warn(err)
	} else {

	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	obj := make(map[string]PubmedFields)
	obj[pmid] = PubmedFields{
		Pmid:        &pmid,
		Doi:         &doi,
		Title:       &title,
		Abs:         &abs,
		Journal:     &journal,
		Issn:        &issn,
		Date:        &date,
		Issue:       &issue,
		Volume:      &volume,
		Corelations: &corela,
		URLs:        &urls,
		Keywords:    &key,
	}
	jsonData, _ = json.MarshalIndent(obj, "", "  ")
	return jsonData
}

func getKeywordsCorleations(doc *prose.Document, keywordsPat *string, corela *map[string]string) {
	for _, sent := range doc.Sentences() {
		kStr := stringo.StrExtract(sent.Text, *keywordsPat, 1000000)
		kStr = slice.DropSliceDup(kStr)
		if len(kStr) >= 2 {
			(*corela)[strings.Join(kStr, "+")] = sent.Text
		}
	}
}
