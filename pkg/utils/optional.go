package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func DirFiles() {
	files, err := ioutil.ReadDir("./cmd/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}
}

func RemoveFile(file string) error {
	err := os.Remove(file)
	if err != nil {
		panic(err)
	}
	return nil
}

// Clear screen
func Clear() {
	out, _ := exec.Command("clear").Output()
	fmt.Println(string(out))
}

// Getting a home dir...
func GetHomeDir() {
	Home, _ := os.UserHomeDir()
	fmt.Println("Home Directory:", Home)
}
