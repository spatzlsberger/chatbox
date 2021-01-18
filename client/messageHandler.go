package main

import (
	"bytes"
	"chatbox/chatboxutil"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// GetMessages is a function to get new messages for the user from the server
func GetMessages(username string) ([]chatboxutil.Message, error) {
	url := "http://localhost:5050/getMessages"
	jsonbody, _ := json.Marshal(chatboxutil.GetMessagesRequest{UserName:username})
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonbody))
	if err != nil{
		return nil, err
	}
	var messagesResponse chatboxutil.GetMessagesResponse
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

	jsonbody, err := json.Marshal(chatboxutil.Message{From:username, To:to, Message:message})
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