package main

import (
	"flag"

	"github.com/akrck02/valhalla-core/modules/api"
	"github.com/akrck02/valhalla-core/modules/cli"
	"github.com/akrck02/valhalla-core/sdk/logger"
)

func main() {
	modeFlag := flag.String("mode", "api", "valhalla-core -mode=api")
	flag.Parse()

	// If startmode is api, start valhalla api
	mode := *modeFlag
	logger.Log("Selected mode: ", mode)
	switch mode {
	case "cli":
		cli.Start()
	case "api":
		api.Start()
	default:
		panic("Selected mode does not exist.")
	}
}
