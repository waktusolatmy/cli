package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thediveo/klo"
	"github.com/waktusolatmy/cli/pkg/api"
	"github.com/waktusolatmy/cli/pkg/formatter"
)

var zoneCode string

var rootCmd = &cobra.Command{
	Use:   "waktusolat",
	Short: "A brief description of your application",
	Run: func(cmd *cobra.Command, args []string) {
		if zoneCode == "" {
			configZone := viper.GetString("zone")
			if configZone == "" {
				fmt.Println("Zone not set")
				os.Exit(1)
			}

			zoneCode = configZone
		}

		prayerTimes, err := api.GetPrayerTimesByZone(zoneCode)
		cobra.CheckErr(err)

		now := time.Now()
		day := now.Day()

		for _, pt := range prayerTimes.Prayers {
			if day != pt.Day {
				continue
			}

			type prayerTime struct {
				Zone    string `json:"zone"`
				Subuh   string `json:"subuh"`
				Syuruk  string `json:"syuruk"`
				Zohor   string `json:"zohor"`
				Asar    string `json:"asar"`
				Maghrib string `json:"maghrib"`
				Isyak   string `json:"isyak"`
			}

			times := prayerTime{
				Zone:    zoneCode,
				Subuh:   formatter.EpochToKitchen(pt.Fajr),
				Syuruk:  formatter.EpochToKitchen(pt.Syuruk),
				Zohor:   formatter.EpochToKitchen(pt.Dhuhr),
				Asar:    formatter.EpochToKitchen(pt.Asr),
				Maghrib: formatter.EpochToKitchen(pt.Maghrib),
				Isyak:   formatter.EpochToKitchen(pt.Isha),
			}

			printer, err := klo.PrinterFromFlag(output, &klo.Specs{
				DefaultColumnSpec: "ZONE:{.Zone},SUBUH:{.Subuh},SYURUK:{.Syuruk},ZOHOR:{.Zohor},ASAR:{.Asar},MAGHRIB:{.Maghrib},ISYAK:{.Isyak}",
			})
			cobra.CheckErr(err)

			printer.Fprint(os.Stdout, times)
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringVarP(&output, "output", "o", "", "Output (json/yaml)")
	rootCmd.Flags().StringVarP(&zoneCode, "zone-code", "z", "", "Overwrite zone code (zone code from config file will be ignored)")
}

func initConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".waktusolatmy")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			cobra.CheckErr(viper.SafeWriteConfig())
		} else {
			cobra.CheckErr(err)
		}
	}
}
