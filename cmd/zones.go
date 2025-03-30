package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/thediveo/klo"
	"github.com/waktusolatmy/cli/pkg/api"
)

var output string

var zonesCmd = &cobra.Command{
	Use:     "zones",
	Aliases: []string{"zone"},
	Short:   "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		zones, err := api.GetZones()
		cobra.CheckErr(err)

		printer, err := klo.PrinterFromFlag(output, &klo.Specs{
			DefaultColumnSpec: "JAKIMCODE:{.JakimCode},NEGERI:{.Negeri},DAERAH:{.Daerah}",
		})
		cobra.CheckErr(err)

		printer.Fprint(os.Stdout, zones)
	},
}

func init() {
	zonesCmd.Flags().StringVarP(&output, "output", "o", "", "Output (json/yaml)")
	rootCmd.AddCommand(zonesCmd)
}
