package parse

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	jsoniter "github.com/json-iterator/go"
	clog "github.com/openbiox/ligo/log"
	"github.com/openbiox/ligo/slice"
	"github.com/openbiox/ligo/stringo"
	prose "gopkg.in/jdkato/prose.v2"
	xurls "mvdan.cc/xurls/v2"
)

var log = clog.Logger

// PubmedFields defines extracted Pubmed fields
type PubmedFields struct {
	Pmid, Doi, Title, Abs, Journal, Issue, Volume, Date, Issn *string
	Author                                                    *[]string
	Affiliation                                               *[]string
	Correlation                                               *map[string]string
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
			defer func() {
				<-sem
			}()
			var pubmed = PubmedArticleSet{}
			if xmlPath != "ParsePubmedXMLStdin" {
				xmlData, err := ioutil.ReadFile(xmlPath)
				if err != nil {
					log.Warnln(err)
				}
				err = xml.Unmarshal(xmlData, &pubmed)
			} else if xmlPath == "ParsePubmedXMLStdin" && len(*stdin) > 0 {
				err = xml.Unmarshal(*stdin, &pubmed)
			}
			if err != nil {
				log.Warnf("%v", err)
				return
			}
			var obj []*PubmedFields
			for _, v := range pubmed.PubmedArticle {
				tmp := getPubmedFields(keywords, &v, callCor)
				obj = append(obj, tmp)
			}
			var json = jsoniter.ConfigCompatibleWithStandardLibrary
			jsonData, _ := json.MarshalIndent(obj, "", "  ")
			io.Copy(of, bytes.NewBuffer(jsonData))
		}(xmlPath, i)
	}
	for i := 0; i < cap(sem); i++ {
		sem <- true
	}
}

func getPubmedFields(keywords *[]string, article *PubmedArticle, callCor bool) *PubmedFields {
	year := article.MedlineCitation.Article.ArticleDate.Year
	month := article.MedlineCitation.Article.ArticleDate.Month
	day := article.MedlineCitation.Article.ArticleDate.Day
	date := fmt.Sprintf("%s/%s/%s", year, month, day)
	var pmid, doi, abs string
	for _, v := range article.PubmedData.ArticleIdList.ArticleId {
		if v.IdType == "pubmed" {
			pmid = v.Text
		} else if v.IdType == "doi" {
			doi = v.Text
		}
	}
	abs = stringo.StrReplaceAll(article.MedlineCitation.Article.Abstract.AbstractText.Text, "\n  *", "\n")
	abs = stringo.StrReplaceAll(abs, "(<[/]AbstractText.*>)|(^\n)|(\n$)", "")
	titleAbs := article.MedlineCitation.Article.ArticleTitle.Text + "\n" + abs
	urls := xurls.Relaxed().FindAllString(titleAbs, -1)
	keywordsPat := strings.Join(*keywords, "|")
	key := stringo.StrExtract(titleAbs, keywordsPat, 1000000)
	key = slice.DropSliceDup(key)

	doc, err := prose.NewDocument(titleAbs)
	corela := make(map[string]string)
	if len(key) >= 2 && callCor {
		getKeywordsCorleations(doc, &keywordsPat, &corela)
	}
	if err != nil {
		log.Warn(err)
	}
	var author, affiliation []string
	for _, v := range article.MedlineCitation.Article.AuthorList.Author {
		author = append(author, v.ForeName+" "+v.LastName)
		affiliationTmp := ""
		for _, j := range v.AffiliationInfo {
			if affiliationTmp == "" {
				affiliationTmp = j.Affiliation
			} else {
				affiliationTmp = affiliationTmp + "; " + j.Affiliation
			}
		}
		affiliation = append(affiliation, affiliationTmp)
	}
	return &PubmedFields{
		Pmid:        &pmid,
		Doi:         &doi,
		Title:       &article.MedlineCitation.Article.ArticleTitle.Text,
		Abs:         &abs,
		Journal:     &article.MedlineCitation.Article.Journal.ISOAbbreviation,
		Issn:        &article.MedlineCitation.Article.Journal.ISSN.Text,
		Date:        &date,
		Issue:       &article.MedlineCitation.Article.Journal.JournalIssue.Issue,
		Volume:      &article.MedlineCitation.Article.Journal.JournalIssue.Volume,
		Author:      &author,
		Affiliation: &affiliation,
		Correlation: &corela,
		URLs:        &urls,
		Keywords:    &key,
	}
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

type PubmedArticleSet struct {
	XMLName       xml.Name        `xml:"PubmedArticleSet"`
	Text          string          `xml:",chardata"`
	PubmedArticle []PubmedArticle `xml:"PubmedArticle"`
}

type PubmedArticle struct {
	Text            string `xml:",chardata"`
	MedlineCitation struct {
		Text   string `xml:",chardata"`
		Status string `xml:"Status,attr"`
		Owner  string `xml:"Owner,attr"`
		PMID   struct {
			Text    string `xml:",chardata"`
			Version string `xml:"Version,attr"`
		} `xml:"PMID"`
		DateRevised struct {
			Text  string `xml:",chardata"`
			Year  string `xml:"Year"`
			Month string `xml:"Month"`
			Day   string `xml:"Day"`
		} `xml:"DateRevised"`
		Article struct {
			Text     string `xml:",chardata"`
			PubModel string `xml:"PubModel,attr"`
			Journal  struct {
				Text string `xml:",chardata"`
				ISSN struct {
					Text     string `xml:",chardata"`
					IssnType string `xml:"IssnType,attr"`
				} `xml:"ISSN"`
				JournalIssue struct {
					Text        string `xml:",chardata"`
					CitedMedium string `xml:"CitedMedium,attr"`
					Volume      string `xml:"Volume"`
					Issue       string `xml:"Issue"`
					PubDate     struct {
						Text  string `xml:",chardata"`
						Year  string `xml:"Year"`
						Month string `xml:"Month"`
					} `xml:"PubDate"`
				} `xml:"JournalIssue"`
				Title           string `xml:"Title"`
				ISOAbbreviation string `xml:"ISOAbbreviation"`
			} `xml:"Journal"`
			ArticleTitle struct {
				Text string   `xml:",chardata"`
				I    []string `xml:"i"`
				Sup  string   `xml:"sup"`
			} `xml:"ArticleTitle"`
			Pagination struct {
				Text       string `xml:",chardata"`
				MedlinePgn string `xml:"MedlinePgn"`
			} `xml:"Pagination"`
			ELocationID struct {
				Text    string `xml:",chardata"`
				EIdType string `xml:"EIdType,attr"`
				ValidYN string `xml:"ValidYN,attr"`
			} `xml:"ELocationID"`
			Abstract struct {
				Text         string `xml:",chardata"`
				AbstractText struct {
					Text string   `xml:",chardata"`
					I    []string `xml:"i"`
					B    string   `xml:"b"`
					Sub  string   `xml:"sub"`
					Sup  []string `xml:"sup"`
				} `xml:"AbstractText"`
			} `xml:"Abstract"`
			AuthorList struct {
				Text       string `xml:",chardata"`
				CompleteYN string `xml:"CompleteYN,attr"`
				Author     []struct {
					Text            string `xml:",chardata"`
					ValidYN         string `xml:"ValidYN,attr"`
					LastName        string `xml:"LastName"`
					ForeName        string `xml:"ForeName"`
					Initials        string `xml:"Initials"`
					AffiliationInfo []struct {
						Text        string `xml:",chardata"`
						Affiliation string `xml:"Affiliation"`
					} `xml:"AffiliationInfo"`
					Identifier struct {
						Text   string `xml:",chardata"`
						Source string `xml:"Source,attr"`
					} `xml:"Identifier"`
				} `xml:"Author"`
			} `xml:"AuthorList"`
			Language  string `xml:"Language"`
			GrantList struct {
				Text       string `xml:",chardata"`
				CompleteYN string `xml:"CompleteYN,attr"`
				Grant      []struct {
					Text    string `xml:",chardata"`
					GrantID string `xml:"GrantID"`
					Acronym string `xml:"Acronym"`
					Agency  string `xml:"Agency"`
					Country string `xml:"Country"`
				} `xml:"Grant"`
			} `xml:"GrantList"`
			PublicationTypeList struct {
				Text            string `xml:",chardata"`
				PublicationType []struct {
					Text string `xml:",chardata"`
					UI   string `xml:"UI,attr"`
				} `xml:"PublicationType"`
			} `xml:"PublicationTypeList"`
			ArticleDate struct {
				Text     string `xml:",chardata"`
				DateType string `xml:"DateType,attr"`
				Year     string `xml:"Year"`
				Month    string `xml:"Month"`
				Day      string `xml:"Day"`
			} `xml:"ArticleDate"`
		} `xml:"Article"`
		MedlineJournalInfo struct {
			Text        string `xml:",chardata"`
			Country     string `xml:"Country"`
			MedlineTA   string `xml:"MedlineTA"`
			NlmUniqueID string `xml:"NlmUniqueID"`
			ISSNLinking string `xml:"ISSNLinking"`
		} `xml:"MedlineJournalInfo"`
		CitationSubset string `xml:"CitationSubset"`
		KeywordList    struct {
			Text    string `xml:",chardata"`
			Owner   string `xml:"Owner,attr"`
			Keyword []struct {
				Text         string `xml:",chardata"`
				MajorTopicYN string `xml:"MajorTopicYN,attr"`
			} `xml:"Keyword"`
		} `xml:"KeywordList"`
	} `xml:"MedlineCitation"`
	PubmedData struct {
		Text    string `xml:",chardata"`
		History struct {
			Text          string `xml:",chardata"`
			PubMedPubDate []struct {
				Text      string `xml:",chardata"`
				PubStatus string `xml:"PubStatus,attr"`
				Year      string `xml:"Year"`
				Month     string `xml:"Month"`
				Day       string `xml:"Day"`
				Hour      string `xml:"Hour"`
				Minute    string `xml:"Minute"`
			} `xml:"PubMedPubDate"`
		} `xml:"History"`
		PublicationStatus string `xml:"PublicationStatus"`
		ArticleIdList     struct {
			Text      string      `xml:",chardata"`
			ArticleId []ArticleId `xml:"ArticleId"`
		} `xml:"ArticleIdList"`
		ReferenceList struct {
			Text      string `xml:",chardata"`
			Reference struct {
				Text          string `xml:",chardata"`
				Citation      string `xml:"Citation"`
				ArticleIdList struct {
					Text      string `xml:",chardata"`
					ArticleId struct {
						Text   string `xml:",chardata"`
						IdType string `xml:"IdType,attr"`
					} `xml:"ArticleId"`
				} `xml:"ArticleIdList"`
			} `xml:"Reference"`
		} `xml:"ReferenceList"`
	} `xml:"PubmedData"`
}

type ArticleId struct {
	Text   string `xml:",chardata"`
	IdType string `xml:"IdType,attr"`
}
