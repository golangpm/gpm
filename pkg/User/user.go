package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/golangpm/pkg/consts"
)

type UserConfig struct {
	Name  string `json:"username"`
	Email string `json:"email"`
}

// Set Global gpm configuration...
func SetConfig() {
	var User, Email string
	// homeDir, _ := os.UserHomeDir()

	// CreateGPMFolder()

	// Set Github username...
	fmt.Print("Enter your github Username: ")
	fmt.Scan(&User)
	// Set Email...
	fmt.Print("Enter your Email: ")
	fmt.Scan(&Email)

	// Write configuration...
	data := UserConfig{
		Name:  User,
		Email: Email,
	}

	// path := fmt.Sprintf("%s/gpm-conf/gpm-config.json", consts.ConfigPath)

	file, _ := json.MarshalIndent(data, "", " ")

	// _ = ioutil.WriteFile(path, file, 7777)

	err := ioutil.WriteFile(consts.ConfigPath, file, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

// Get config file...
func GetConfig() {

	// CreateGPMFolder()

	// homeDir, _ := os.UserHomeDir()

	// config := fmt.Sprintf("%s/gpm-conf/gpm-config.json", consts.ConfigPath)
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
	fmt.Println("User Name: " + users.Name)
	fmt.Println("User Email: " + users.Email)
}

// Get username from config...
func GetUser() {
	// homeDir := GetHomeDir
	// path := fmt.Sprintf("%s/gpm-conf/gpm-config.json", consts.ConfigPath)

	fileContent, err := os.Open(consts.ConfigPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer fileContent.Close()
	byteResult, _ := ioutil.ReadAll(fileContent)
	var users UserConfig
	json.Unmarshal(byteResult, &users)
	fmt.Println(users.Name)
}

// Getting a home dir...
func GetHomeDir() {
	Home, _ := os.UserHomeDir()
	fmt.Println("Home Directory:", Home)
}

// func CreateGPMFolder() error {
// 	homeDir := GetHomeDir
// 	config := fmt.Sprintf("%s/gpm-conf/gpm-config.json", consts.ConfigPath)

// 	if _, err := os.Stat(config); os.IsNotExist(err) {
// 		return os.Mkdir(config, os.ModeDir|0755)
// 	}
// 	return nil
// }
