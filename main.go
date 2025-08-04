package main

import (
	"flag"

	"github.com/akrck02/valhalla-core/api"
	"github.com/akrck02/valhalla-core/cli"
	"github.com/akrck02/valhalla-core/logger"
)

func main() {

	modeFlag := flag.String("mode", "cli", "valhalla-core -mode=api")
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
