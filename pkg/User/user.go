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

	err := ioutil.WriteFile(consts.ConfigPath, file, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

// Get config file...
func GetConfig() {
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
