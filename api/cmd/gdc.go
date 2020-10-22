package cmd

import (
	"github.com/clindet/bget/api/fetch"
	"github.com/clindet/bget/api/types"
	"github.com/spf13/cobra"
)

var endp types.GdcEndpoints

var GdcCmd = &cobra.Command{
	Use:   "gdc",
	Short: "Query GDC portal website APIs.",
	Long:  `Query GDC portal website APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		GdcCmdRunOptions(cmd, args)
	},
}

func GdcCmdRunOptions(cmd *cobra.Command, args []string) {
	endp.ExtraParams.From = bapiClis.From
	endp.ExtraParams.Size = bapiClis.Size
	endp.ExtraParams.Format = bapiClis.Format
	endp.ExtraParams.Query = bapiClis.Query
	endp.ExtraParams.Pretty = bapiClis.PrettyJSON
	if endp.ExtraParams.JSON {
		endp.ExtraParams.Format = "json"
	}
	if endp.Status || endp.Projects || endp.Cases || endp.Files || endp.Annotations || endp.Data || endp.Manifest || endp.Slicing {
		initCmd(cmd, args)
		fetch.Gdc(&endp, &bapiClis, nil)
		bapiClis.HelpFlags = false
	}
	if bapiClis.HelpFlags {
		cmd.Help()
	}
}

func init() {
	setGlobalFlag(GdcCmd, &bapiClis)
	GdcCmd.Flags().BoolVarP(&endp.ExtraParams.RemoteName, "remote-name", "n", false, "Use remote defined filename.")
	GdcCmd.Flags().BoolVarP(&endp.Status, "status", "s", false, "Check GDC portal status (https://portal.gdc.cancer.gov/).")
	GdcCmd.Flags().BoolVarP(&endp.Cases, "cases", "c", false, "Retrive cases info from GDC portal.")
	GdcCmd.Flags().BoolVarP(&endp.Files, "files", "f", false, "Retrive files info from GDC portal.")
	GdcCmd.Flags().BoolVarP(&endp.Projects, "projects", "p", false, "Retrive projects meta info from GDC portal.")
	GdcCmd.Flags().BoolVarP(&endp.Annotations, "annotations", "a", false, "Retrive annotations info from GDC portal.")
	GdcCmd.Flags().BoolVarP(&endp.Data, "data", "d", false, "Retrive /data from GDC portal.")
	GdcCmd.Flags().BoolVarP(&endp.Manifest, "manifest", "m", false, "Retrive /manifest data from GDC portal.")
	GdcCmd.Flags().BoolVarP(&endp.Slicing, "slicing", "", false, "Retrive BAM slicing from GDC portal.")
	GdcCmd.Flags().StringVarP(&endp.ExtraParams.Filter, "filter", "", "", "Retrive data with GDC filter.")
	GdcCmd.Flags().BoolVarP(&endp.Legacy, "legacy", "l", false, "Use legacy API of GDC portal.")
	GdcCmd.Flags().StringVarP(&endp.ExtraParams.Token, "token", "", "", "Token to access GDC.")
	GdcCmd.Flags().StringVarP(&endp.ExtraParams.Sort, "sort", "", "", "Sort parameters.")
	GdcCmd.Flags().StringVarP(&endp.ExtraParams.Fields, "fields", "", "", "Fields parameters.")
	GdcCmd.Flags().IntVarP(&bapiClis.From, "from", "", -1, "Parameters of API control the start item of retrived data.")
	GdcCmd.Flags().IntVarP(&bapiClis.Size, "size", "", -1, "Parameters of API control the lenth of retrived data. Default is auto determined.")
	GdcCmd.Flags().StringVarP(&bapiClis.Query, "query", "q", "", "Query specifies the search query for record retrieval (required).")

	GdcCmd.Example = `  # retrive projects meta info from GDC portal
  bget api gdc -p
  bget api gdc -p --json-pretty
  bget api gdc -p -q TARGET-NBL --json-pretty
  bget api gdc -p --format tsv > tcga_projects.tsv
  bget api gdc -p --format csv > tcga_projects.csv
  bget api gdc -p --from 1 --size 2
  # check GDC portal status (https://portal.gdc.cancer.gov/)
  bget api gdc -s
  # retrive cases info from GDC portal
  bget api gdc -c
  # retrive files info from GDC portal
  bget api gdc -f
  # retrive annotations info from GDC portal
  bget api gdc -a
  # query manifest for gdc-client
  bget api gdc -m -q "5b2974ad-f932-499b-90a3-93577a9f0573,556e5e3f-0ab9-4b6c-aa62-c42f6a6cf20c" -o my_manifest.txt
  bget api gdc -m -q "5b2974ad-f932-499b-90a3-93577a9f0573,556e5e3f-0ab9-4b6c-aa62-c42f6a6cf20c" > my_manifest.txt
  bget api gdc -m -q "5b2974ad-f932-499b-90a3-93577a9f0573,556e5e3f-0ab9-4b6c-aa62-c42f6a6cf20c" -n
  # query data that only support the samll filesize
  bget api gdc -d -q "5b2974ad-f932-499b-90a3-93577a9f0573" -n`
}
