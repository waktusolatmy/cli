package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thediveo/klo"
	"github.com/waktusolatmy/cli/pkg/api"
)

var currentZoneCmd = &cobra.Command{
	Use:   "current",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		zoneCode := viper.GetString("zone")

		zones, err := api.GetZones()
		cobra.CheckErr(err)

		for _, z := range zones {
			if z.JakimCode != zoneCode {
				continue
			}

			printer, err := klo.PrinterFromFlag(output, &klo.Specs{
				DefaultColumnSpec: "JAKIMCODE:{.JakimCode},NEGERI:{.Negeri},DAERAH:{.Daerah}",
			})
			cobra.CheckErr(err)

			printer.Fprint(os.Stdout, z)
			break
		}
	},
}

func init() {
	currentZoneCmd.Flags().StringVarP(&output, "output", "o", "", "Output (json/yaml)")
	zonesCmd.AddCommand(currentZoneCmd)
}
