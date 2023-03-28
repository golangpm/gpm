package main

import (
	"fmt"
	"log"
	"os"

	user "github.com/golangpm/pkg/User"
	"github.com/golangpm/pkg/utils"
)

func main() {
	// All os arguments...
	args := os.Args

	if len(args) < 0 {
		log.Fatal("Application don't have arguments...")
		return
	}
	// GPM defaults...
	if len(args) < 2 {
		utils.Logo()
	}

	// is goapp check...
	if args[1] == "goapp" {
		utils.CreateApp(args[2])
		fmt.Printf("The application %s has been created!\n", os.Args[2])
	}

	// is start command check...
	if args[1] == "start" {
		// utils.CurrentDir()
		Args()
	}
	// set config...
	if args[1] == "set-config" {
		user.SetConfig()
	}

	// get config data...
	if args[1] == "config" {
		// user.GetHomeDir()
		user.GetConfig()

		// user.GetUser()
	}

	// Print executable Args
	// Args()
}

// Commands list -h
func Commands() {

}

func Args() {
	fmt.Printf("\nARGS:\n%v\n", os.Args)
}
