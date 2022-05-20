package models

import (
	"io/ioutil"
	"log"
	"net/netip"
	"os/exec"

	"gopkg.in/yaml.v3"
)

type IPRange struct {
	RangesFile string
	Device     string
	Ranges     []string
}

func (ipr *IPRange) SetIps() {
	rangeFile, err := ioutil.ReadFile(ipr.RangesFile)
	if err != nil {
		log.Fatalln("Error loading the ranges file", err)
		return
	}

	yaml.Unmarshal(rangeFile, &ipr.Ranges)

	for _, cidrRange := range ipr.Ranges {
		prefix, err := netip.ParsePrefix(cidrRange)
		if err != nil {
			log.Println(err)
		}

		for addr := prefix.Addr(); prefix.Contains(addr); addr = addr.Next() {
			exec.Command("/sbin/ip", "a", "a", addr.String(), "dev", ipr.Device)
		}
	}
}
