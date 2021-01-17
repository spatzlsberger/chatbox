package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

type message struct {
	From string `json:"From"`
	To string `json:"To"`
	Message string `json:"Message"`
}

type getUsersResponse struct {
	Users []string `json:"Users"`
}

type getMessagesRequest struct {
	Username string
}

type getMessagesResponse struct {
	Messages []message `json:"Messages"`
}

//map of users the system is aware of
var userNames map[string]int64
// map of waiting messages for a user with that ID
var messages map[int64][]message

func sendNewMessage(w http.ResponseWriter, r *http.Request) {
	var mess message
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
	var request getMessagesRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		panic(err)
	}
	fmt.Println(messages)
	fmt.Println(request.Username)
	fmt.Println(userNames)
	id := request.Username
	fmt.Println(userNames[id])
	messagesToReturn := messages[userNames[id]]

	json.NewEncoder(w).Encode(getMessagesResponse{messagesToReturn})
	messages[userNames[id]] = nil

}

func getUsers(w http.ResponseWriter, r *http.Request) {
	/*
	Returns all the users that system has registered so far. 
	*/
	w.WriteHeader(http.StatusOK)
	var userarray []string
	for k := range userNames {
		userarray = append(userarray, k)
	}
	resp := getUsersResponse{userarray}
	json.NewEncoder(w).Encode(resp)
}

func addNewUser(w http.ResponseWriter, r *http.Request) {
	
	var u User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Created"))

	id := rand.Int63n(50000)

	userNames[u.Username] = id
	w.Write([]byte(strconv.FormatInt(id, 10)))

	for k := range userNames {
		fmt.Println(k)
	}
}

func main () {
	fmt.Println("Starting server on port 5050")
	userNames = map[string]int64{}
	messages = map[int64][]message{}
	http.HandleFunc("/getUsers", getUsers)
	http.HandleFunc("/sendMessage", sendNewMessage)
	http.HandleFunc("/getMessages", getMessages)
	http.HandleFunc("/addNewUser", addNewUser)
	http.ListenAndServe("localhost:5050", nil)
}