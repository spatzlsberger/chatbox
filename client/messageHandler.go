package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type getMessagesRequest struct {
	Username string `json:"username"`
}

type getMessagesResponse struct {
	Messages []NewMessage `json:"Messages"`
}

// GetMessages is a function to get new messages for the user from the server
func GetMessages(username string) ([]NewMessage, error) {
	url := "http://localhost:5050/getMessages"
	jsonbody, _ := json.Marshal(getMessagesRequest{username})
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonbody))
	if err != nil{
		return nil, err
	}
	var messagesResponse getMessagesResponse
	err = json.NewDecoder(resp.Body).Decode(&messagesResponse)

	if err != nil{
		return nil, err
	}

	return messagesResponse.Messages, nil
}


func showAllUsers(username string){
	users, err := getUsers()
	if err != nil{
		panic("error occured during getting users")
	}
	for index, user := range users{
		fmt.Println("User ", index, ": ", user)
	}
}

func sendMessage(username string) (bool, error){
	fmt.Print("Enter the username that you want to send the message to: ")
	to, err := getInput()
	if err != nil{
		return false, err
	}
	fmt.Print("Enter your message you wish to send: ")
	message, err := getInput()

	jsonbody, err := json.Marshal(NewMessage{username, to, message})
	url := "http://localhost:5050/sendMessage"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonbody))
	if err != nil{
		return false, err
	}
	if resp.StatusCode != 201{
		return false, errors.New("Message request sent to server, but failed to create")
	}
	return true, nil
}