package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var currentZoneCmd = &cobra.Command{
	Use:   "current",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.GetString("zone"))
	},
}

func init() {
	zonesCmd.AddCommand(currentZoneCmd)
}
