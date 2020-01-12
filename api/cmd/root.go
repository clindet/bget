package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/openbiox/bget/api/types"
	cio "github.com/openbiox/ligo/io"
	clog "github.com/openbiox/ligo/log"
	"github.com/openbiox/ligo/stringo"
	"github.com/spf13/cobra"
)

var log = clog.Logger
var wd = ""

// BapiClis is the param to run bapi
var BapiClis = types.BapiClisT{}

// BapiCmd is the "bget api" cobra command object
var BapiCmd = &cobra.Command{
	Use:   "api",
	Short: "Query bioinformatics website APIs.",
	Long:  `Query bioinformatics website APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		BapiCmdRunOptions(cmd)
	},
}

// Execute main interface of bget api
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
	BapiClis.HelpFlags = true
	BapiCmd.AddCommand(bioToolsCmd)
	BapiCmd.AddCommand(ncbiCmd)
	BapiCmd.AddCommand(gdcCmd)
	BapiCmd.AddCommand(dataset2toolsCmd)
	BapiCmd.AddCommand(cligovCmd)
	wd, _ := os.Getwd()

	BapiCmd.PersistentFlags().StringVarP(&(BapiClis.TaskID), "task-id", "", stringo.RandString(15), "Task ID (random).")
	BapiCmd.PersistentFlags().StringVarP(&(BapiClis.LogDir), "log-dir", "", path.Join(wd, "_log"), "Log dir.")
	BapiCmd.PersistentFlags().StringVarP(&(BapiClis.Quiet), "quiet", "", "false", "No log output.")
	BapiCmd.PersistentFlags().StringVarP(&(BapiClis.SaveLog), "save-log", "", "true", "Save download log to local file.")

	BapiCmd.Version = BapiClis.Version
}

// BapiCmdRunOptions is the main of bapi
func BapiCmdRunOptions(cmd *cobra.Command) {
	if hasDir, _ := cio.PathExists(BapiClis.Outfn); BapiClis.Outfn != "" && !hasDir {
		if err := cio.CreateDir(path.Dir(BapiClis.Outfn)); err != nil {
			log.Warningf("Could not to create %s", path.Dir(BapiClis.Outfn))
		}
	}
	if BapiClis.HelpFlags {
		cmd.Help()
	}
}
