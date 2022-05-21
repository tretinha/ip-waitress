package cmd

import (
	"strings"

	"github.com/tretinha/ip-waitress/models"

	"github.com/spf13/cobra"
)

var addRouteCmd = &cobra.Command{
	Use:   "add-route",
	Short: "Add routes to the routing table",
	Run: func(cmd *cobra.Command, args []string) {
		routesFile := "./routes.yaml"
		via := ""

		for _, v := range args {
			argumentsParts := strings.Split(v, "=")
			if len(argumentsParts) == 2 {
				if argumentsParts[0] == "routes" {
					routesFile = argumentsParts[1]
				}
				if argumentsParts[0] == "via" {
					via = argumentsParts[1]
				}
			}
		}

		rTable := models.RTable{
			RoutesFile: routesFile,
			Via:        via,
		}

		rTable.SetRoutes()
	},
}

func init() {
	rootCmd.AddCommand(addRouteCmd)
	addRouteCmd.Flags()
}
