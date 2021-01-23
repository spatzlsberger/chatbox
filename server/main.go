package main

import (
	"chatbox/chatboxutil"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

//map of users the system is aware of
var userNames map[string]int64
// map of waiting messages for a user with that ID
var messages map[int64][]chatboxutil.Message

func sendNewMessage(w http.ResponseWriter, r *http.Request) {
	var mess chatboxutil.Message
	err := json.NewDecoder(r.Body).Decode(&mess)
	if err != nil {
		panic(err)
	}

	id := userNames[mess.To]
	messages[id] = append(messages[id], mess)
	fmt.Println(messages[id])
	w.WriteHeader(http.StatusCreated)
}

func getMessages(w http.ResponseWriter, r *http.Request) {
	var request chatboxutil.GetMessagesRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		panic(err)
	}
	fmt.Println(messages)
	fmt.Println(request.UserName)
	fmt.Println(userNames)
	id := request.UserName
	fmt.Println(userNames[id])
	messagesToReturn := messages[userNames[id]]

	json.NewEncoder(w).Encode(chatboxutil.GetMessagesResponse{Messages:messagesToReturn})
	messages[userNames[id]] = []chatboxutil.Message{}

}

func getUsers(w http.ResponseWriter, r *http.Request) {
	/*
	Returns all the users that system has registered so far. 
	*/
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	var userarray []string
	for k := range userNames {
		userarray = append(userarray, k)
	}
	resp := chatboxutil.GetUsersReponse{Users:userarray}
	json.NewEncoder(w).Encode(resp)
}

func addNewUser(w http.ResponseWriter, r *http.Request) {
	
	var u chatboxutil.NewUser
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Created"))
	id := rand.Int63n(50000)

	userNames[u.UserName] = id
	w.Write([]byte(strconv.FormatInt(id, 10)))

	for k := range userNames {
		fmt.Println(k)
	}
}

func main () {
	fmt.Println("Starting server on port 5050")
	userNames = map[string]int64{}
	messages = map[int64][]chatboxutil.Message{}
	http.HandleFunc("/getUsers", getUsers)
	http.HandleFunc("/sendMessage", sendNewMessage)
	http.HandleFunc("/getMessages", getMessages)
	http.HandleFunc("/addNewUser", addNewUser)
	http.ListenAndServe("localhost:5050", nil)
}