package main

import (
	"fmt"
	"os"

	user "github.com/golangpm/pkg/User"
	"github.com/golangpm/pkg/consts"
	"github.com/golangpm/pkg/utils"
)

func main() {
	// All os arguments...
	args := os.Args

	// GPM defaults...
	if len(args) < 2 {
		utils.Logo()
		Help()
		// Args()
	} else if len(args) > 2 && args[1] == "goapp" && args[2] != "" {
		// is goapp check...
		utils.Logo()
		utils.CreateApp(args[2])
		fmt.Printf("The application %s has been created!\n", os.Args[2])
		return
	} else if len(args) > 0 && args[1] == "version" || args[1] == "-v" {
		fmt.Printf("%vGPM %v", consts.Yellow, consts.Reset)
		Version()
	} else if len(args) > 1 && args[1] == "start" { // is start command check...
		utils.StartApp()
		return
	} else {
		utils.Logo()
		fmt.Printf("\tVersion: %v%s%v\nPlease enter the application name...\n", consts.Green, consts.Version, consts.Reset)
	}
	// set config...
	if len(args) > 1 && args[1] == "set-config" {
		user.SetConfig()
	}

	// get config data...
	if len(args) > 1 && args[1] == "config" {
		user.GetConfig()
	}

	// Print executable Args
	// Args()
}

// Commands list -h
func Help() {
	fmt.Println(`GPM Commands:
	General Commands:
	gpm goapp - create a new application
	gpm start - Start your application

	User Configuration:
	gpm config - show the configuration
	gpm set-config - set the configuration
	`)
}

func Args() {
	fmt.Printf("\nARGS:\n%v\n", os.Args)
}

func Version() {
	fmt.Printf("Version: %v%s%v\n\n", consts.Green, consts.Version, consts.Reset)
}
