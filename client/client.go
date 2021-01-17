package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type newMessage struct {
	From string `json:"From"`
	To string `json:"To"`
	Message string `json:"Message"`
}

type newUser struct {
	Name string `json:"name"`
	UserName string `json:"username"`
}

type getUsersReponse struct {
	Users []string `json:"Users"`
}

func registerUser() (bool, error){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Enter your username: ")
	userName, _ := reader.ReadString('\n')
	
	url := "http://localhost:5050/addNewUser"
	newUserBody := newUser{name, userName}
	jsonBody, _ := json.Marshal(newUserBody)
	req, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))

	if err != nil {
		return false, err
	}

	defer req.Body.Close()
	fmt.Println("response status: ", req.Status)
	io.Copy(os.Stdout, req.Body)

	return true, nil
}

func getUsers() ([]string, error){
	url := "http://localhost:5050/getUsers"
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	var availableUsers getUsersReponse
	err = json.NewDecoder(req.Body).Decode(&availableUsers)
	for _, returnedUser := range availableUsers.Users {
		fmt.Println("User: ", returnedUser)
	}
	return availableUsers.Users, nil
}

func main() {
	registerUser()
	getUsers()
}