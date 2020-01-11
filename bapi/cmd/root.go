package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/openbiox/bget/bapi/types"
	cio "github.com/openbiox/butils/io"
	"github.com/openbiox/butils/log"
	"github.com/spf13/cobra"
)

var BapiClis = types.BapiClisT{}

var BapiCmd = &cobra.Command{
	Use:   "api",
	Short: "Query bioinformatics website APIs.",
	Long:  `Query bioinformatics website APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		BapiCmdRunOptions(cmd)
	},
}

// Execute main interface of butils
func Execute() {
	if err := BapiCmd.Execute(); err != nil {
		if !BapiCmd.HasFlags() && !BapiCmd.HasSubCommands() {
			BapiCmd.Help()
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func init() {
	BapiClis.Quiet = false
	BapiClis.HelpFlags = true
	BapiCmd.AddCommand(ncbiCmd)
	BapiCmd.AddCommand(gdcCmd)
	BapiCmd.AddCommand(dataset2toolsCmd)
	BapiCmd.AddCommand(cligovCmd)

	BapiCmd.Version = BapiClis.Version
}

func BapiCmdRunOptions(cmd *cobra.Command) {
	if BapiClis.Quiet {
		log.SetOutput(ioutil.Discard)
	} else {
		log.SetOutput(os.Stderr)
	}
	if hasDir, _ := cio.PathExists(BapiClis.Outfn); BapiClis.Outfn != "" && !hasDir {
		if err := cio.CreateDir(path.Dir(BapiClis.Outfn)); err != nil {
			log.FATAL(fmt.Sprintf("Could not to create %s", path.Dir(BapiClis.Outfn)))
		}
	}
	if BapiClis.HelpFlags {
		cmd.Help()
	}
}
