package chatboxutil

import "fmt"

// Message struct used on client and server
type Message struct {
	From string `json:"From"`
	To string `json:"To"`
	Message string `json:"Message"`
}

func (mess Message) String() string{
	return fmt.Sprintf("From: %s\nMessage:%s", mess.From, mess.Message)
}

// NewUser struct used on client and server
type NewUser struct {
	Name string `json:"name"`
	UserName string `json:"username"`
}

// GetUsersReponse used on client and server
type GetUsersReponse struct {
	Users []string `json:"Users"`
}

// GetMessagesResponse used on client and server
type GetMessagesResponse struct {
	Messages []Message `json:"Messages"`
}

// GetMessagesRequest used on client and server
type GetMessagesRequest struct {
	UserName string `json:"username"`
}