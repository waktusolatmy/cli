package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/waktusolatmy/cli/pkg/api"
)

var setZoneCmd = &cobra.Command{
	Use:   "set",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		newZone := args[0]

		zones, err := api.GetZones()
		cobra.CheckErr(err)

		found := false
		for _, zone := range zones {
			if zone.JakimCode == newZone {
				found = true
				break
			}
		}

		if !found {
			fmt.Println("Invalid zone code")
			os.Exit(1)
		}

		viper.Set("zone", newZone)
		cobra.CheckErr(viper.WriteConfig())
		fmt.Printf("New zone set to %s\n", newZone)
	},
}

func init() {
	zonesCmd.AddCommand(setZoneCmd)
}
