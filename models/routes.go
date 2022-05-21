package models

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	"gopkg.in/yaml.v3"
)

type RTable struct {
	RoutesFile string
	Via        string
	Routes     []string
}

func (rTable *RTable) SetRoutes() {
	routesFile, err := ioutil.ReadFile(rTable.RoutesFile)
	if err != nil {
		log.Fatalln("Error loading the routes file", err)
		return
	}

	yaml.Unmarshal(routesFile, &rTable.Routes)

	for _, route := range rTable.Routes {
		cmd := exec.Command("/sbin/ip", "r", "a", route, "via", rTable.Via)
		_, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
