package cmd

import (
	"github.com/openbiox/bget/bapi/fetch"
	"github.com/openbiox/bget/bapi/types"
	"github.com/spf13/cobra"
)

var endp types.GdcEndpoints

var gdcCmd = &cobra.Command{
	Use:   "gdc",
	Short: "Query GDC portal website APIs.",
	Long:  `Query GDC portal APIs. More see here https://github.com/openbiox/bget/bapi.`,
	Run: func(cmd *cobra.Command, args []string) {
		gdcCmdRunOptions(cmd)
	},
}

func gdcCmdRunOptions(cmd *cobra.Command) {
	endp.ExtraParams.From = BapiClis.From
	endp.ExtraParams.Size = BapiClis.Size
	endp.ExtraParams.Format = BapiClis.Format
	endp.ExtraParams.Query = BapiClis.Query
	endp.ExtraParams.Pretty = BapiClis.PrettyJSON
	if endp.ExtraParams.JSON {
		endp.ExtraParams.Format = "json"
	}
	if endp.Status || endp.Projects || endp.Cases || endp.Files || endp.Annotations || endp.Data || endp.Manifest || endp.Slicing {
		fetch.Gdc(&endp, &BapiClis)
		BapiClis.HelpFlags = false
	}
	if BapiClis.HelpFlags {
		cmd.Help()
	}
}

func init() {
	gdcCmd.Flags().BoolVarP(&endp.ExtraParams.RemoteName, "remote-name", "n", false, "Use remote defined filename.")
	gdcCmd.Flags().BoolVarP(&endp.Status, "status", "s", false, "Check GDC portal status (https://portal.gdc.cancer.gov/).")
	gdcCmd.Flags().BoolVarP(&endp.Cases, "cases", "c", false, "Retrive cases info from GDC portal.")
	gdcCmd.Flags().BoolVarP(&endp.Files, "files", "f", false, "Retrive files info from GDC portal.")
	gdcCmd.Flags().BoolVarP(&endp.Projects, "projects", "p", false, "Retrive projects meta info from GDC portal.")
	gdcCmd.Flags().BoolVarP(&endp.Annotations, "annotations", "a", false, "Retrive annotations info from GDC portal.")
	gdcCmd.Flags().BoolVarP(&endp.Data, "data", "d", false, "Retrive /data from GDC portal.")
	gdcCmd.Flags().BoolVarP(&endp.Manifest, "manifest", "m", false, "Retrive /manifest data from GDC portal.")
	gdcCmd.Flags().BoolVarP(&endp.Slicing, "slicing", "", false, "Retrive BAM slicing from GDC portal.")
	gdcCmd.Flags().StringVarP(&endp.ExtraParams.Filter, "filter", "", "", "Retrive data with GDC filter.")
	gdcCmd.Flags().BoolVarP(&endp.Legacy, "legacy", "l", false, "Use legacy API of GDC portal.")
	gdcCmd.Flags().StringVarP(&endp.ExtraParams.Token, "token", "", "", "Token to access GDC.")
	gdcCmd.Flags().StringVarP(&endp.ExtraParams.Sort, "sort", "", "", "Sort parameters.")
	gdcCmd.Flags().StringVarP(&endp.ExtraParams.Fields, "fields", "", "", "Fields parameters.")
	gdcCmd.Flags().StringVarP(&BapiClis.Outfn, "outfn", "o", "", "Out specifies destination of the returned data (default to stdout).")
	gdcCmd.Example = `  # retrive projects meta info from GDC portal
  bget api gdc -p
  bget api gdc -p --json-pretty
  bget api gdc -p -q TARGET-NBL --json-pretty
  bget api gdc -p --format tsv > tcga_projects.tsv
  bget api gdc -p --format csv > tcga_projects.csv
  bget api gdc -p --from 1 --szie 2
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
