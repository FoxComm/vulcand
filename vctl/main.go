package main

import (
	"os"

	"github.com/FoxComm/vulcand/log"
	"github.com/FoxComm/vulcand/plugin/registry"
	"github.com/FoxComm/vulcand/vctl/command"
)

var vulcanUrl string

func main() {
	log.EnsureLoggerExist("console", "error")

	cmd := command.NewCommand(registry.GetRegistry())
	err := cmd.Run(os.Args)
	if err != nil {
		log.Errorf("error: %s\n", err)
	}
}
