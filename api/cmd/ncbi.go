package cmd

import (
	"bufio"
	"io/ioutil"
	"os"

	"github.com/openbiox/bget/api/fetch"
	"github.com/openbiox/bget/api/parse"
	"github.com/openbiox/bget/api/types"
	"github.com/openbiox/ligo/flag"
	"github.com/openbiox/ligo/stringo"
	"github.com/spf13/cobra"
)

var ncbiClis types.NcbiClisT

var ncbiCmd = &cobra.Command{
	Use:   "ncbi",
	Short: "Query ncbi website APIs.",
	Long:  `Query ncbi website APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		ncbiCmdRunOptions(cmd, args)
	},
}

func ncbiCmdRunOptions(cmd *cobra.Command, args []string) {
	cleanArgs := []string{}
	var stdin []byte
	var err error
	hasStdin := false
	if cleanArgs, hasStdin = flag.CheckStdInFlag(cmd); hasStdin {
		reader := bufio.NewReader(os.Stdin)
		stdin, err = ioutil.ReadAll(reader)
		if err != nil {
			log.Fatal(err)
		}
	}
	if BapiClis.Format == "" {
		BapiClis.Format = "XML"
	}
	if BapiClis.Email != "" && BapiClis.Query != "" {
		initCmd(cmd, args)
		fetch.Ncbi(&BapiClis, &ncbiClis)
		BapiClis.HelpFlags = false
	}
	if ncbiClis.NcbiXMLToJSON == "pubmed" {
		if len(cleanArgs) >= 1 || len(stdin) > 0 {
			ncbiClis.NcbiXMLPaths = append(ncbiClis.NcbiXMLPaths, cleanArgs...)
			keywordsList := stringo.StrSplit(ncbiClis.NcbiKeywords, ", |,", 10000)
			parse.PubmedXML(&ncbiClis.NcbiXMLPaths, &stdin, BapiClis.Outfn, &keywordsList, BapiClis.Thread, BapiClis.CallCor)
		}
		BapiClis.HelpFlags = false
	} else if ncbiClis.NcbiXMLToJSON == "sra" {
		if len(cleanArgs) >= 1 || len(stdin) > 0 {
			ncbiClis.NcbiXMLPaths = append(ncbiClis.NcbiXMLPaths, cleanArgs...)
			keywordsList := stringo.StrSplit(ncbiClis.NcbiKeywords, ", |,", 10000)
			parse.SraXML(&ncbiClis.NcbiXMLPaths, &stdin, BapiClis.Outfn, &keywordsList, BapiClis.Thread, BapiClis.CallCor)
		}
		BapiClis.HelpFlags = false
	}
	if BapiClis.HelpFlags {
		cmd.Help()
	}
}

func init() {
	setGlobalFlag(ncbiCmd, &BapiClis)
	ncbiCmd.Flags().StringVarP(&ncbiClis.NcbiDB, "db", "d", "pubmed", "Db specifies the database to search")
	ncbiCmd.Flags().IntVarP(&ncbiClis.NcbiRetmax, "per-size", "m", 100, "Retmax specifies the number of records to be retrieved per request.")
	ncbiCmd.Flags().StringVarP(&ncbiClis.NcbiXMLToJSON, "xml2json", "", "", "Convert XML files to json [e.g. pubmed].")
	ncbiCmd.Flags().StringVarP(&ncbiClis.NcbiKeywords, "keywords", "w", "algorithm, tool, model, pipleline, method, database, workflow, dataset, bioinformatics, sequencing, http, github.com, gitlab.com, bitbucket.org", "Keywords to extracted from abstract.")
	ncbiCmd.Flags().IntVarP(&BapiClis.Thread, "thread", "t", 2, "Thread to process.")
	ncbiCmd.Flags().BoolVarP(&BapiClis.CallCor, "call-cor", "", false, "Wheather to calculate the corelated keywords, and return the sentence contains >=2 keywords.")
	ncbiCmd.Flags().StringVarP(&BapiClis.Email, "email", "e", "your_email@domain.com", "Email specifies the email address to be sent to the server (NCBI website is required).")

	ncbiCmd.Example = `  # query pubmed with 'B-ALL'
  bget api ncbi -d pubmed -q B-ALL --format XML -e your_email@domain.com

  # query pubmed and convert it to json format that also extract all URLs and calculate the words connections
  bget api ncbi -q "Galectins control MTOR and AMPK in response to lysosomal damage to induce autophagy OR MTOR-independent autophagy induced by interrupted endoplasmic reticulum-mitochondrial Ca2+ communication: a dead end in cancer cells. OR The PARK10 gene USP24 is a negative regulator of autophagy and ULK1 protein stability OR Coordinate regulation of autophagy and the ubiquitin proteasome system by MTOR." | bget api ncbi --xml2json pubmed -k "MAPK, MTOR, autophagy" --call-cor - | > final.json

  # query larger items
  k="algorithm, tool, model, pipleline, method, database, workflow, dataset, bioinformatics, sequencing, http, github.com, gitlab.com, bitbucket.org, RNA-Seq, DNA, profile, landscape"
  bget api ncbi -q "dataset and RNA-seq and bioinformatics[journal]" -e "your_email@domain.com" -m 20 | awk '/<[?]xml version="1.0" [?]>/{close(f); f="abstract.http.XML.tmp" ++c;next} {print>f;}' && bget api ncbi --xml2json pubmed abstract.http.XML.tmp* -k "${k}" --call-cor -t 11 > final.json`
}
