package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func StartApp(appName string) {
	// path to the application...
	path := fmt.Sprintf("%s/cmd/%s", appName, appName)
	// Run the application...
	c := exec.Command("go", "run", path)
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr
	c.Run()
}
