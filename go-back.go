package main

import (
	"flag"
	"net"
	"os/exec"
)

var targetIP net.IP
var temp string

func main() {
	flag.StringVar(&temp, "ip", "", "specify IP to connect back to")
	flag.Parse()

	if temp != "" {
		targetIP = net.ParseIP(temp)
		if targetIP == nil {
			panic(`Invalid IP format`)
		}
	} else if temp == "" {
		panic(`target not specified`)
	}

	c := exec.Command(`ping`, targetIP.String())
	if err := c.Run(); err != nil {
		panic(`could not ping target`)
	}
}