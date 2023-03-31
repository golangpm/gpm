package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"

	"github.com/golangpm/pkg/consts"
	cp "github.com/otiai10/copy"
)

type Project struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Author  string `json:"author"`
	Email   string `json:"email"`
	License string `json:"license"`
}

type UserConfig struct {
	Name  string `json:"username"`
	Email string `json:"email"`
}

/* Create a new App by Template
*  goapp command
* Example: gpm goapp `app name`
 */
func CreateApp(appName string) {
	// var users UserConfig
	fmt.Printf("Creating %s%s%s app...\n", consts.Bold, appName, consts.Reset)

	// Create th app folder
	CreateAppFolder(appName)

	CreateStructure(appName)

	// Create internal directory...
	MakeInternal(appName, "app")
	MakeInternal(appName, "pkg")

	// Create general file structure...
	CreateTemplFie(fmt.Sprintf("%s/cmd/main.go", appName), fmt.Sprintf("%s/cmd/%s.go", appName, appName))

	// MakeFile
	SetMakeFile(appName)
	// Remove MakeFile.txt
	RemoveFile(fmt.Sprintf("%s/MakeFile.txt", appName))
	// Create project.json
	ProjectJson(appName)
	// Get username and Email from gpm config...
	fileContent, err := os.Open(consts.ConfigPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer fileContent.Close()
	byteResult, _ := ioutil.ReadAll(fileContent)
	var users UserConfig
	json.Unmarshal(byteResult, &users)

	// Done output message
	fmt.Printf("cd %s%s%s\ngo mod init github.com/%s/%s\n", consts.Cyan, appName, consts.Reset, users.Name, appName)
}

// Create App Folder ...
func CreateAppFolder(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		if e := os.Mkdir(path, os.ModePerm); e != nil {
			log.Fatal(e)
		}
	}
}

// Create File Structure ...
func CreateStructure(appName string) {
	if runtime.GOOS == "darwin" {
		if err := cp.Copy("/usr/local/bin/Template", appName); err != nil {
			log.Fatal("Не удалось создать файловую структуру...")
		}
	} else if runtime.GOOS == "linux" {
		if err := cp.Copy("/usr/bin/Template", appName); err != nil {
			log.Fatal("Не удалось создать файловую структуру...")
		}
	}
}

// Rename main.go file ...
func CreateTemplFie(file, appName string) {
	if e := os.Rename(file, appName); e != nil {
		log.Fatal(e)
	}
}

// Create internal directory
func MakeInternal(appName, path string) {
	err := os.MkdirAll(fmt.Sprintf("%s/internal/%s", appName, path), 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
}

// Makefile variables
func SetMakeFile(appName string) {
	input, err := ioutil.ReadFile(fmt.Sprintf("%s/Makefile.txt", appName))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output := bytes.Replace(input, []byte("YOUR_APP_NAME"), []byte(appName), -1)

	if err = ioutil.WriteFile(fmt.Sprintf("%s/Makefile", appName), output, 0777); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ProjectJson(appName string) {
	fileContent, err := os.Open(consts.ConfigPath)

	if err != nil {
		log.Fatal(err)
		return
	}

	// log.Println("The File is opened successfully...")

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	var users UserConfig
	json.Unmarshal(byteResult, &users)

	author := fmt.Sprintf("https://github.com/%s/%s", users.Name, appName)

	userEmail := users.Email

	data := Project{
		Name:    appName,
		Version: "0.1.0",
		Author:  author,
		Email:   userEmail,
		License: "MIT License",
	}

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile(fmt.Sprintf("%s/project.json", appName), file, 0644)
}

// Get project name from config...
func GetProjectName() {
	fileContent, err := os.Open(consts.ConfigPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer fileContent.Close()
	byteResult, _ := ioutil.ReadAll(fileContent)
	var project Project
	json.Unmarshal(byteResult, &project)
	fmt.Println(project.Name)
}
