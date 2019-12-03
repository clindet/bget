package cmd

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"os"

	bflag "github.com/openbiox/bget/flag"
	bfmt "github.com/openbiox/bget/fmt"
	"github.com/openbiox/butils/log"
	"github.com/spf13/cobra"
)

var FmtClis = bfmt.FmtClisT{}

var FmtCmd = &cobra.Command{
	Use:   "fmt [input1 input2]",
	Short: "A set of file format (fmt) command.",
	Long:  `A set of file format (fmt) command. More see here https://github.com/openbiox/bget/api.`,
	Run: func(cmd *cobra.Command, args []string) {
		FmtCmdRunOptions(cmd)
	},
}

func FmtCmdRunOptions(cmd *cobra.Command) {
	cleanArgs := []string{}
	hasStdin := false
	if cleanArgs, hasStdin = bflag.CheckStdInFlag(cmd); hasStdin {
		reader := bufio.NewReader(os.Stdin)
		result, err := ioutil.ReadAll(reader)
		if err != nil {
			log.Fatal(err)
		} else if len(result) > 0 {
			FmtClis.Stdin = &result
		}
	} else {
		FmtClis.Stdin = nil
	}

	if len(cleanArgs) >= 1 || hasStdin {
		FmtClis.Files = &cleanArgs
		runFlag := false
		if FmtClis.PrettyJSON {
			bfmt.PrettyJSON(&FmtClis, bgetClis.thread)
			runFlag = true
		} else if FmtClis.JSONToSlice {
			bfmt.JSON2Slice(&FmtClis, bgetClis.thread)
			runFlag = true
		}
		if !runFlag && hasStdin {
			io.Copy(os.Stdout, bytes.NewBuffer(*FmtClis.Stdin))
		}
		bgetClis.helpFlags = false
	}
	if bgetClis.helpFlags {
		cmd.Help()
	}
}

func init() {
	FmtCmd.Flags().IntVarP(&bgetClis.thread, "thread", "t", 1, "Thread to process.")
	FmtCmd.Flags().BoolVarP(&FmtClis.JSONToSlice, "json-to-slice", "", false, "Convert key-value JSON  to []key-value and easy to export to readable table.")
	FmtCmd.Flags().BoolVarP(&FmtClis.PrettyJSON, "json-pretty", "", false, "Pretty json files.")
	FmtCmd.Flags().IntVarP(&FmtClis.Indent, "indent", "", 4, "Control the indent of output json files.")
	FmtCmd.Flags().BoolVarP(&FmtClis.SortKeys, "sort-keys", "", false, "Control wheather to sort JSON key.")
	FmtCmd.Example = `  bget api ncbi -q "Galectins control MTOR and AMPK in response to lysosomal damage to induce autophagy OR MTOR-independent autophagy induced by interrupted endoplasmic reticulum-mitochondrial Ca2+ communication: a dead end in cancer cells. OR The PARK10 gene USP24 is a negative regulator of autophagy and ULK1 protein stability OR Coordinate regulation of autophagy and the ubiquitin proteasome system by MTOR." | bget ncbi --xml2json pubmed - | sed 's;}{;,;g' | bget fmt --json-to-slice --indent 4 -| json2csv -o final.csv`
	JSON := make(map[int]map[string]interface{})
	FmtClis.JSON = &JSON
	Table := make(map[int][]interface{})
	FmtClis.Table = &Table
}
