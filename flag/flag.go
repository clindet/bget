package flag

import "github.com/spf13/cobra"

func CheckStdInFlag(cmd *cobra.Command) (args []string, hasStdin bool) {
	for _, v := range cmd.Flags().Args() {
		if v != "-" {
			args = append(args, v)
		} else {
			hasStdin = true
		}
	}
	return args, hasStdin
}
