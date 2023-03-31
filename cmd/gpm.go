package main

import (
	"fmt"
	"os"

	user "github.com/golangpm/pkg/User"
	"github.com/golangpm/pkg/utils"
)

func main() {
	// All os arguments...
	args := os.Args

	// if len(args) < 0 {
	// 	log.Fatal("Application don't have arguments...")
	// 	return
	// }
	// GPM defaults...
	if len(args) < 2 {
		utils.Logo()
		Help()
	}

	// is goapp check...
	if len(args) > 1 && args[1] == "goapp" {
		utils.Logo()
		utils.CreateApp(args[2])
		fmt.Printf("The application %s has been created!\n", os.Args[2])
		return
	}
	// is start command check...
	if len(args) > 1 && args[1] == "start" {
		utils.StartApp()
		return
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
