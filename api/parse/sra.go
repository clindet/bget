package parse

import (
	"bytes"
	"encoding/xml"
	"io"
	"io/ioutil"
	"os"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/openbiox/ligo/slice"
	"github.com/openbiox/ligo/stringo"
	prose "gopkg.in/jdkato/prose.v2"
	xurls "mvdan.cc/xurls/v2"
)

// SraFields defines extracted Sra fields
type SraFields struct {
	Title        *string
	StudyTitle   *string
	Abstract     *string
	Type         *string
	SRX          *string
	SRA          *string
	SRAFile      *SRAFile
	Correlation  *map[string]string
	AbstractURLs *[]string
	Keywords     *[]string
}

// SraXML convert Sra XML to json
func SraXML(xmlPaths *[]string, stdin *[]byte, outfn string, keywords *[]string, thread int, callCor bool) {
	if len(*xmlPaths) == 1 {
		thread = 1
	}
	if len(*stdin) > 0 {
		*xmlPaths = append(*xmlPaths, "ParseSraXMLStdin")
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
			var sra = SraSets{}
			if xmlPath != "ParseSraXMLStdin" {
				xmlData, err := ioutil.ReadFile(xmlPath)
				if err != nil {
					log.Warnln(err)
				}
				err = xml.Unmarshal(xmlData, &sra)
			} else if xmlPath == "ParseSraXMLStdin" && len(*stdin) > 0 {
				err = xml.Unmarshal(*stdin, &sra)
			}
			if err != nil {
				log.Warnf("%v", err)
				return
			}
			var obj []*SraFields
			var done = make(map[string]int)
			for _, v := range sra.EXPERIMENTPACKAGE {
				tmp := getSraFields(keywords, &v, callCor, done)
				obj = append(obj, tmp)
				done[*tmp.Title+*tmp.StudyTitle] = 1
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

func getSraFields(keywords *[]string, sra *ExperimentPkg, callCor bool, done map[string]int) *SraFields {
	titleAbs := sra.EXPERIMENT.TITLE + "\n" + sra.STUDY.DESCRIPTOR.STUDYTITLE +
		"\n" + sra.STUDY.DESCRIPTOR.STUDYABSTRACT
	doc, err := prose.NewDocument(titleAbs)
	if done[sra.EXPERIMENT.TITLE+sra.STUDY.DESCRIPTOR.STUDYTITLE] == 1 {
		return &SraFields{
			Title:      &sra.EXPERIMENT.TITLE,
			StudyTitle: &sra.STUDY.DESCRIPTOR.STUDYTITLE,
			Type:       &sra.STUDY.DESCRIPTOR.STUDYTYPE.ExistingStudyType,
			SRX:        &sra.EXPERIMENT.Accession,
			SRA:        &sra.RUNSET.RUN.Accession,
			SRAFile:    &sra.RUNSET.RUN.SRAFiles.SRAFile,
		}
	}
	corela := make(map[string]string)
	urls := xurls.Relaxed().FindAllString(titleAbs, -1)
	keywordsPat := strings.Join(*keywords, "|")
	key := stringo.StrExtract(titleAbs, keywordsPat, 1000000)
	key = slice.DropSliceDup(key)
	if len(key) >= 2 && callCor {
		getKeywordsCorleations(doc, &keywordsPat, &corela)
	}
	if err != nil {
		log.Warn(err)
	}
	return &SraFields{
		Title:        &sra.EXPERIMENT.TITLE,
		StudyTitle:   &sra.STUDY.DESCRIPTOR.STUDYTITLE,
		Abstract:     &sra.STUDY.DESCRIPTOR.STUDYABSTRACT,
		Type:         &sra.STUDY.DESCRIPTOR.STUDYTYPE.ExistingStudyType,
		SRX:          &sra.EXPERIMENT.Accession,
		SRA:          &sra.RUNSET.RUN.Accession,
		SRAFile:      &sra.RUNSET.RUN.SRAFiles.SRAFile,
		Correlation:  &corela,
		AbstractURLs: &urls,
		Keywords:     &key,
	}
}

// SraSets defines extracted Sra fields
type SraSets struct {
	XMLName           xml.Name        `xml:"EXPERIMENT_PACKAGE_SET"`
	EXPERIMENTPACKAGE []ExperimentPkg `xml:"EXPERIMENT_PACKAGE"`
}

type ExperimentPkg struct {
	EXPERIMENT struct {
		Accession   string `xml:"accession,attr"`
		Alias       string `xml:"alias,attr"`
		IDENTIFIERS struct {
			PRIMARYID string `xml:"PRIMARY_ID"`
		} `xml:"IDENTIFIERS"`
		TITLE    string `xml:"TITLE"`
		STUDYREF struct {
			Accession   string `xml:"accession,attr"`
			IDENTIFIERS struct {
				PRIMARYID  string `xml:"PRIMARY_ID"`
				EXTERNALID struct {
					Namespace string `xml:"namespace,attr"`
				} `xml:"EXTERNAL_ID"`
			} `xml:"IDENTIFIERS"`
		} `xml:"STUDY_REF"`
		DESIGN struct {
			DESIGNDESCRIPTION string `xml:"DESIGN_DESCRIPTION"`
			SAMPLEDESCRIPTOR  struct {
				Accession   string `xml:"accession,attr"`
				IDENTIFIERS struct {
					PRIMARYID  string `xml:"PRIMARY_ID"`
					EXTERNALID struct {
						Namespace string `xml:"namespace,attr"`
					} `xml:"EXTERNAL_ID"`
				} `xml:"IDENTIFIERS"`
			} `xml:"SAMPLE_DESCRIPTOR"`
			LIBRARYDESCRIPTOR struct {
				LIBRARYNAME      string `xml:"LIBRARY_NAME"`
				LIBRARYSTRATEGY  string `xml:"LIBRARY_STRATEGY"`
				LIBRARYSOURCE    string `xml:"LIBRARY_SOURCE"`
				LIBRARYSELECTION string `xml:"LIBRARY_SELECTION"`
				LIBRARYLAYOUT    struct {
					PAIRED string `xml:"PAIRED"`
				} `xml:"LIBRARY_LAYOUT"`
			} `xml:"LIBRARY_DESCRIPTOR"`
		} `xml:"DESIGN"`
		PLATFORM struct {
			ILLUMINA struct {
				INSTRUMENTMODEL string `xml:"INSTRUMENT_MODEL"`
			} `xml:"ILLUMINA"`
		} `xml:"PLATFORM"`
	} `xml:"EXPERIMENT"`
	SUBMISSION struct {
		LabName     string `xml:"lab_name,attr"`
		CenterName  string `xml:"center_name,attr"`
		Accession   string `xml:"accession,attr"`
		Alias       string `xml:"alias,attr"`
		IDENTIFIERS struct {
			PRIMARYID string `xml:"PRIMARY_ID"`
		} `xml:"IDENTIFIERS"`
	} `xml:"SUBMISSION"`
	Organization struct {
		Type    string `xml:"type,attr"`
		Name    string `xml:"Name"`
		Address struct {
			PostalCode  string `xml:"postal_code,attr"`
			Department  string `xml:"Department"`
			Institution string `xml:"Institution"`
			Street      string `xml:"Street"`
			City        string `xml:"City"`
			Country     string `xml:"Country"`
		} `xml:"Address"`
		Contact struct {
			Email    string `xml:"email,attr"`
			SecEmail string `xml:"sec_email,attr"`
			Address  struct {
				PostalCode  string `xml:"postal_code,attr"`
				Department  string `xml:"Department"`
				Institution string `xml:"Institution"`
				Street      string `xml:"Street"`
				City        string `xml:"City"`
				Country     string `xml:"Country"`
			} `xml:"Address"`
			Name struct {
				First string `xml:"First"`
				Last  string `xml:"Last"`
			} `xml:"Name"`
		} `xml:"Contact"`
	} `xml:"Organization"`
	STUDY struct {
		CenterName  string `xml:"center_name,attr"`
		Alias       string `xml:"alias,attr"`
		Accession   string `xml:"accession,attr"`
		IDENTIFIERS struct {
			PRIMARYID  string `xml:"PRIMARY_ID"`
			EXTERNALID struct {
				Namespace string `xml:"namespace,attr"`
				Label     string `xml:"label,attr"`
			} `xml:"EXTERNAL_ID"`
		} `xml:"IDENTIFIERS"`
		DESCRIPTOR struct {
			STUDYTITLE string `xml:"STUDY_TITLE"`
			STUDYTYPE  struct {
				ExistingStudyType string `xml:"existing_study_type,attr"`
			} `xml:"STUDY_TYPE"`
			STUDYABSTRACT     string `xml:"STUDY_ABSTRACT"`
			CENTERPROJECTNAME string `xml:"CENTER_PROJECT_NAME"`
		} `xml:"DESCRIPTOR"`
	} `xml:"STUDY"`
	SAMPLE struct {
		Alias       string `xml:"alias,attr"`
		Accession   string `xml:"accession,attr"`
		IDENTIFIERS struct {
			PRIMARYID  string `xml:"PRIMARY_ID"`
			EXTERNALID struct {
				Namespace string `xml:"namespace,attr"`
			} `xml:"EXTERNAL_ID"`
		} `xml:"IDENTIFIERS"`
		TITLE      string `xml:"TITLE"`
		SAMPLENAME struct {
			TAXONID        string `xml:"TAXON_ID"`
			SCIENTIFICNAME string `xml:"SCIENTIFIC_NAME"`
		} `xml:"SAMPLE_NAME"`
		SAMPLELINKS struct {
			SAMPLELINK struct {
				XREFLINK struct {
					DB    string `xml:"DB"`
					ID    string `xml:"ID"`
					LABEL string `xml:"LABEL"`
				} `xml:"XREF_LINK"`
			} `xml:"SAMPLE_LINK"`
		} `xml:"SAMPLE_LINKS"`
		SAMPLEATTRIBUTES struct {
			SAMPLEATTRIBUTE []struct {
				TAG   string `xml:"TAG"`
				VALUE string `xml:"VALUE"`
			} `xml:"SAMPLE_ATTRIBUTE"`
		} `xml:"SAMPLE_ATTRIBUTES"`
	} `xml:"SAMPLE"`
	Pool struct {
		Member struct {
			MemberName  string `xml:"member_name,attr"`
			Accession   string `xml:"accession,attr"`
			SampleName  string `xml:"sample_name,attr"`
			SampleTitle string `xml:"sample_title,attr"`
			Spots       string `xml:"spots,attr"`
			Bases       string `xml:"bases,attr"`
			TaxID       string `xml:"tax_id,attr"`
			Organism    string `xml:"organism,attr"`
			IDENTIFIERS struct {
				PRIMARYID  string `xml:"PRIMARY_ID"`
				EXTERNALID struct {
					Namespace string `xml:"namespace,attr"`
				} `xml:"EXTERNAL_ID"`
			} `xml:"IDENTIFIERS"`
		} `xml:"Member"`
	} `xml:"Pool"`
	RUNSET struct {
		RUN struct {
			Accession           string `xml:"accession,attr"`
			Alias               string `xml:"alias,attr"`
			TotalSpots          string `xml:"total_spots,attr"`
			TotalBases          string `xml:"total_bases,attr"`
			Size                string `xml:"size,attr"`
			LoadDone            string `xml:"load_done,attr"`
			Published           string `xml:"published,attr"`
			IsPublic            string `xml:"is_public,attr"`
			ClusterName         string `xml:"cluster_name,attr"`
			StaticDataAvailable string `xml:"static_data_available,attr"`
			IDENTIFIERS         struct {
				PRIMARYID string `xml:"PRIMARY_ID"`
			} `xml:"IDENTIFIERS"`
			EXPERIMENTREF struct {
				Accession   string `xml:"accession,attr"`
				IDENTIFIERS string `xml:"IDENTIFIERS"`
			} `xml:"EXPERIMENT_REF"`
			Pool struct {
				Member struct {
					MemberName  string `xml:"member_name,attr"`
					Accession   string `xml:"accession,attr"`
					SampleName  string `xml:"sample_name,attr"`
					SampleTitle string `xml:"sample_title,attr"`
					Spots       string `xml:"spots,attr"`
					Bases       string `xml:"bases,attr"`
					TaxID       string `xml:"tax_id,attr"`
					Organism    string `xml:"organism,attr"`
					IDENTIFIERS struct {
						PRIMARYID  string `xml:"PRIMARY_ID"`
						EXTERNALID struct {
							Namespace string `xml:"namespace,attr"`
						} `xml:"EXTERNAL_ID"`
					} `xml:"IDENTIFIERS"`
				} `xml:"Member"`
			} `xml:"Pool"`
			SRAFiles   SraFns `xml:"SRAFiles"`
			CloudFiles struct {
				CloudFile []struct {
					Filetype string `xml:"filetype,attr"`
					Provider string `xml:"provider,attr"`
					Location string `xml:"location,attr"`
				} `xml:"CloudFile"`
			} `xml:"CloudFiles"`
			Statistics struct {
				Nreads string `xml:"nreads,attr"`
				Nspots string `xml:"nspots,attr"`
				Read   []struct {
					Index   string `xml:"index,attr"`
					Count   string `xml:"count,attr"`
					Average string `xml:"average,attr"`
					Stdev   string `xml:"stdev,attr"`
				} `xml:"Read"`
			} `xml:"Statistics"`
			Databases struct {
				Database struct {
					Table struct {
						Name       string `xml:"name,attr"`
						Statistics struct {
							Source string `xml:"source,attr"`
							Rows   struct {
								Count string `xml:"count,attr"`
							} `xml:"Rows"`
							Elements struct {
								Count string `xml:"count,attr"`
							} `xml:"Elements"`
						} `xml:"Statistics"`
					} `xml:"Table"`
				} `xml:"Database"`
			} `xml:"Databases"`
			Bases struct {
				CsNative string `xml:"cs_native,attr"`
				Count    string `xml:"count,attr"`
				Base     []struct {
					Value string `xml:"value,attr"`
					Count string `xml:"count,attr"`
				} `xml:"Base"`
			} `xml:"Bases"`
		} `xml:"RUN"`
	} `xml:"RUN_SET"`
}

type SraFns struct {
	SRAFile SRAFile `xml:"SRAFile"`
}

type SRAFile struct {
	Cluster      string `xml:"cluster,attr"`
	Filename     string `xml:"filename,attr"`
	URL          string `xml:"url,attr"`
	Size         string `xml:"size,attr"`
	Date         string `xml:"date,attr"`
	Md5          string `xml:"md5,attr"`
	SemanticName string `xml:"semantic_name,attr"`
	Supertype    string `xml:"supertype,attr"`
	Sratoolkit   string `xml:"sratoolkit,attr"`
	Alternatives []struct {
		URL        string `xml:"url,attr"`
		FreeEgress string `xml:"free_egress,attr"`
		AccessType string `xml:"access_type,attr"`
		Org        string `xml:"org,attr"`
	} `xml:"Alternatives"`
}
