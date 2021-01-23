package main

import (
	"bufio"
	"bytes"
	"chatbox/chatboxutil"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func registerUser() (string, error){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Enter your username: ")
	userName, _ := reader.ReadString('\n')

	name = strings.TrimRight(name, "\r\n")
	userName = strings.TrimRight(userName, "\r\n")

	url := "http://localhost:5050/addNewUser"
	newUserBody := chatboxutil.NewUser{Name:name, UserName:userName}
	jsonBody, _ := json.Marshal(newUserBody)
	req, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))

	if err != nil {
		return "", err
	}

	defer req.Body.Close()
	fmt.Println("response status: ", req.Status)
	io.Copy(os.Stdout, req.Body)

	return userName, nil
}

func getUsers() ([]string, error){
	url := "http://localhost:5050/getUsers"
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	var availableUsers chatboxutil.GetUsersReponse
	err = json.NewDecoder(req.Body).Decode(&availableUsers)
	return availableUsers.Users, nil
}

func mainloop(users []string, username string) {
	fmt.Println("entering main loop")
	for {
		promptSelection()
		selection, _ := getInput()
		valSelection, err := validateInput(selection)
		if err != nil{
			fmt.Println("Invalid selection, try again.")
			continue
		}
		if valSelection == 1{
			showAllUsers(username)
		} else if valSelection == 2{
			sendMessage(username)
		} else if valSelection == 3{
			messages, _ := GetMessages(username)
			displayMessages(messages)
		}
	}
}

func getInput() (string, error){
	reader := bufio.NewReader(os.Stdin)
	selection, _ := reader.ReadString('\n')
	selection = strings.TrimRight(selection, "\r\n")
	return selection, nil
}

func promptSelection() {
	fmt.Println("What would you like to do?")
	fmt.Println("1. See all registered users")
	fmt.Println("2. Send Message")
	fmt.Println("3. Check Messages from User")
	// TODO add more options
	fmt.Print("Your Selection: ")

}

func validateInput(input string) (int64, error){
	number, err := strconv.ParseInt(input, 10, 0)
	if err != nil{
		return 0, err
	}
	if number < 1 || number > 3{
		return 0, errors.New("Invalid entry")
	}
	return number, nil
}

func main(){
	username, err := registerUser()
	users, err := getUsers()
	if err != nil {
		panic(err)
	}
	mainloop(users, username)
}