package cmd

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"

	"github.com/clindet/bget/api/fetch"
	"github.com/clindet/bget/api/types"
	"github.com/openbiox/ligo/flag"
	cio "github.com/openbiox/ligo/io"
	"github.com/spf13/cobra"
)

var ncbiClis types.NcbiClisT

var NcbiCmd = &cobra.Command{
	Use:   "ncbi",
	Short: "Query ncbi website APIs.",
	Long:  `Query ncbi website APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		NcbiCmdRunOptions(cmd, args)
	},
}

func NcbiCmdRunOptions(cmd *cobra.Command, args []string) {
	var stdin []byte
	var err error
	if _, hasStdin := flag.CheckStdInFlag(cmd); hasStdin {
		reader := bufio.NewReader(os.Stdin)
		stdin, err = ioutil.ReadAll(reader)
		if err != nil {
			log.Fatal(err)
		}
		bapiClis.Query = bapiClis.Query + " " + string(stdin)
	}
	if len(args) > 0 {
		bapiClis.Query = bapiClis.Query + " " + strings.Join(args, " OR ")
	}
	if bapiClis.Format == "" {
		bapiClis.Format = "XML"
	}
	if bapiClis.Email != "" && bapiClis.Query != "" {
		initCmd(cmd, args)
		of := cio.NewOutStream(bapiClis.Outfn, "")
		fetch.Ncbi(&bapiClis, &ncbiClis, of)
		bapiClis.HelpFlags = false
	}
	if bapiClis.HelpFlags {
		cmd.Help()
	}
}

func init() {
	setGlobalFlag(NcbiCmd, &bapiClis)
	NcbiCmd.Flags().StringVarP(&ncbiClis.NcbiDB, "db", "d", "pubmed", "Db specifies the database to search")
	NcbiCmd.Flags().IntVarP(&ncbiClis.NcbiRetmax, "per-size", "m", 100, "Retmax specifies the number of records to be retrieved per request.")
	NcbiCmd.Flags().IntVarP(&bapiClis.Thread, "thread", "t", 2, "Thread to process.")
	NcbiCmd.Flags().StringVarP(&bapiClis.Email, "email", "e", "your_email@domain.com", "Email specifies the email address to be sent to the server (NCBI website is required).")
	NcbiCmd.Flags().IntVarP(&bapiClis.From, "from", "", -1, "Parameters of API control the start item of retrived data.")
	NcbiCmd.Flags().IntVarP(&bapiClis.Size, "size", "", -1, "Parameters of API control the lenth of retrived data. Default is auto determined.")
	NcbiCmd.Flags().StringVarP(&bapiClis.Query, "query", "q", "", "Query specifies the search query for record retrieval (required).")

	NcbiCmd.Example = `  # query pubmed with 'B-ALL'
  bget api ncbi -d pubmed -q B-ALL --format XML -e your_email@domain.com

  # query pubmed and convert it to json format
  bget api ncbi -q "Galectins control MTOR and AMPK in response to lysosomal damage to induce autophagy OR MTOR-independent autophagy induced by interrupted endoplasmic reticulum-mitochondrial Ca2+ communication: a dead end in cancer cells. OR The PARK10 gene USP24 is a negative regulator of autophagy and ULK1 protein stability OR Coordinate regulation of autophagy and the ubiquitin proteasome system by MTOR." | bioctl cvrt --xml2json pubmed - > final.json

  # query larger items
  bget api ncbi -q "dataset and RNA-seq and bioinformatics[journal]" -e "your_email@domain.com" -m 20 | awk '/<[?]xml version="1.0" [?]>/{close(f); f="abstract.http.XML.tmp" ++c;next} {print>f;}' && bioctl cvrt --xml2json pubmed abstract.http.XML.tmp* -t 11 > final.json`
}
