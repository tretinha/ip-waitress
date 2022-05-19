package cmd

import (
	"strings"

	"github.com/tretinha/ip-waitress/models"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run ip-waitress",
	Run: func(cmd *cobra.Command, args []string) {
		rangesFile := "./ranges.yaml"
		device := ""

		for _, v := range args {
			argumentsParts := strings.Split(v, "=")
			if len(argumentsParts) == 2 {
				if argumentsParts[0] == "ranges" {
					rangesFile = argumentsParts[1]
				}
				if argumentsParts[0] == "device" {
					device = argumentsParts[1]
				}
			}
		}

		ipRange := models.IPRange{
			RangesFile: rangesFile,
			Device:     device,
		}

		ipRange.SetIps()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags()
}
