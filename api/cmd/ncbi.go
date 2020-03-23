package cmd

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"

	"github.com/openbiox/bget/api/fetch"
	"github.com/openbiox/bget/api/types"
	"github.com/openbiox/ligo/flag"
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
	var stdin []byte
	var err error
	if _, hasStdin := flag.CheckStdInFlag(cmd); hasStdin {
		reader := bufio.NewReader(os.Stdin)
		stdin, err = ioutil.ReadAll(reader)
		if err != nil {
			log.Fatal(err)
		}
		BapiClis.Query = BapiClis.Query + " " + string(stdin)
	}
	if len(args) > 0 {
		BapiClis.Query = BapiClis.Query + " " + strings.Join(args, " OR ")
	}
	if BapiClis.Format == "" {
		BapiClis.Format = "XML"
	}
	if BapiClis.Email != "" && BapiClis.Query != "" {
		initCmd(cmd, args)
		fetch.Ncbi(&BapiClis, &ncbiClis)
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
	ncbiCmd.Flags().IntVarP(&BapiClis.Thread, "thread", "t", 2, "Thread to process.")
	ncbiCmd.Flags().StringVarP(&BapiClis.Email, "email", "e", "your_email@domain.com", "Email specifies the email address to be sent to the server (NCBI website is required).")
	ncbiCmd.Flags().IntVarP(&BapiClis.From, "from", "", -1, "Parameters of API control the start item of retrived data.")
	ncbiCmd.Flags().IntVarP(&BapiClis.Size, "size", "", -1, "Parameters of API control the lenth of retrived data. Default is auto determined.")

	ncbiCmd.Example = `  # query pubmed with 'B-ALL'
  bget api ncbi -d pubmed -q B-ALL --format XML -e your_email@domain.com

  # query pubmed and convert it to json format
  bget api ncbi -q "Galectins control MTOR and AMPK in response to lysosomal damage to induce autophagy OR MTOR-independent autophagy induced by interrupted endoplasmic reticulum-mitochondrial Ca2+ communication: a dead end in cancer cells. OR The PARK10 gene USP24 is a negative regulator of autophagy and ULK1 protein stability OR Coordinate regulation of autophagy and the ubiquitin proteasome system by MTOR." | bioctl cvrt --xml2json pubmed - > final.json

  # query larger items
  bget api ncbi -q "dataset and RNA-seq and bioinformatics[journal]" -e "your_email@domain.com" -m 20 | awk '/<[?]xml version="1.0" [?]>/{close(f); f="abstract.http.XML.tmp" ++c;next} {print>f;}' && bioctl cvrt --xml2json pubmed abstract.http.XML.tmp* -t 11 > final.json`
}
