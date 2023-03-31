package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/golangpm/pkg/consts"
)

type ProjectName struct {
	Name string `json:"name"`
}

func StartApp() {
	fileContent, err := os.Open(consts.ProjectJson)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer fileContent.Close()
	byteResult, _ := ioutil.ReadAll(fileContent)
	var project ProjectName
	json.Unmarshal(byteResult, &project)

	// path to the application...
	path := fmt.Sprintf("cmd/%s.go", project.Name)
	// Run the application...
	c := exec.Command("go", "run", path)
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr
	c.Run()
}
