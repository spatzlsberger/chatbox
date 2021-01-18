package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
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
	newUserBody := newUser{strings.TrimRight(name,"\r\n"), strings.TrimRight(userName, "\r\n")}
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

func mainloop([]string) {
	fmt.Println("entering main loop")
	for {
		selection, _ := promptSelection()
		fmt.Println(selection)

	}
}

func promptSelection() (int64, error) {
	fmt.Println("What would you like to do?")
	fmt.Println("1. See all registered users")
	fmt.Println("2. Send Message")
	fmt.Println("3. Check Messages from User")
	// TODO add more options
	fmt.Print("Your Selection: ")
	reader := bufio.NewReader(os.Stdin)
	selection, _ := reader.ReadString('\n')
	selection = strings.TrimRight(selection, "\r\n")
	fmt.Println(selection)
	valSelection, err := validateInput(selection)
	if err != nil{
		fmt.Println("Invalid selection, try again.")
		return 0, err
	}
	return valSelection, nil
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
	registerUser()
	users, err := getUsers()
	if err != nil {
		panic(err)
	}
	mainloop(users)
}