package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	cp "github.com/otiai10/copy"
)

func main() {
	CheckOS()
}

func SetupGPM(file, path string) {
	if err := cp.Copy(file, path); err != nil {
		log.Fatal("Не удалось скопировать файлы при установке...")
	}
}

func CheckOS() {
	os := runtime.GOOS
	switch os {
	case "windows":
		fmt.Println("Installing GPM for Windows...")

		fmt.Println("GPM is installed!")
	case "darwin":
		fmt.Println("Installing GPM for MAC operating system...")
		ExecMacScript()
		SetConfigFile()
		fmt.Println("GPM is installed!")
	case "linux":
		fmt.Println("Installing GPM for Linux...")
		ExecLinuxScript()
		SetConfigFile()
		fmt.Println("GPM is installed!")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func ExecMacScript() {
	// sudo cp -rf Template bin/goapp /usr/local/bin/
	c := exec.Command("sudo", "cp", "-rf", "assets/Template", "bin/gpm", "/usr/local/bin/")
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr
	c.Run()
}
func ExecLinuxScript() {
	// sudo cp -rf Template bin/goapp /usr/local/bin/
	c := exec.Command("sudo", "cp", "-rf", "assets/Template", "bin/gpm", "/usr/bin/")
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr
	c.Run()
}

func SetConfigFile() {
	// configs/gpm-config.json
	c := exec.Command("sh", "./Scripts/install.sh")
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr
	c.Run()
}

// Getting a home dir...
func GetHomeDir() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(homeDir)
}
