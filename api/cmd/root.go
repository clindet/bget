package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/clindet/bget/api/types"
	cio "github.com/openbiox/ligo/io"
	"github.com/openbiox/ligo/stringo"
	"github.com/spf13/cobra"
)

var bapiClis = types.BapiClisT{}

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
	bapiClis.HelpFlags = true
	BapiCmd.AddCommand(BioToolsCmd)
	BapiCmd.AddCommand(MgRastCmd)
	BapiCmd.AddCommand(NcbiCmd)
	BapiCmd.AddCommand(GdcCmd)
	BapiCmd.AddCommand(Dataset2toolsCmd)
	BapiCmd.AddCommand(CligovCmd)
	BapiCmd.AddCommand(CrossRefCmd)
	BapiCmd.AddCommand(Covid19Cmd)
	wd, _ := os.Getwd()

	BapiCmd.PersistentFlags().StringVarP(&(bapiClis.TaskID), "task-id", "k", stringo.RandString(15), "Task ID (random).")
	BapiCmd.PersistentFlags().StringVarP(&(bapiClis.LogDir), "log-dir", "", path.Join(wd, "_log"), "Log dir.")
	BapiCmd.PersistentFlags().IntVarP(&(bapiClis.Verbose), "verbose", "", 1, "verbose level (0:no output, 1: basic level, 2: with env info)")
	BapiCmd.PersistentFlags().BoolVarP(&(bapiClis.SaveLog), "save-log", "", false, "Save log to local file.")
	BapiCmd.PersistentFlags().StringVarP(&bapiClis.Outfn, "outfn", "o", "", "Out specifies destination of the returned data (default to stdout).")

	BapiCmd.Version = bapiClis.Version
}

// BapiCmdRunOptions is the main of bapi
func BapiCmdRunOptions(cmd *cobra.Command) {
	if hasDir, _ := cio.PathExists(bapiClis.Outfn); bapiClis.Outfn != "" && !hasDir {
		if err := cio.CreateDir(path.Dir(bapiClis.Outfn)); err != nil {
			log.Warningf("Could not to create %s", path.Dir(bapiClis.Outfn))
		}
	}
	if bapiClis.HelpFlags {
		cmd.Help()
	}
}
