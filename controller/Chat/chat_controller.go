package Chat

import (
	"encoding/json"
	"go-live-chat/middleware"
	"go-live-chat/model"
	"net/http"
	"strconv"
	"time"
)

//  chatController handles chat logic

type chatController struct{}

// sendMsg handles sending a message from a user to another user
func (c * chatController) sendMessage(w http.ResponseWriter, r * http.Request) {
	// use middleware to auth
	middleware.SocketAuthMiddleware(func(w http.ResponseWriter, r * http.Request) {
		var msg model.Message
		if err :=json.NewDecoder(r.Body).Decode(&msg); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// get the userID from the req header
		userID, err := strconv.Atoi(r.Header.Get("UserID"))
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusUnauthorized)
			return
		}

		msg.SenderID = userID
		msg.Timestamp = time.Now()

		// save msg to db
		if err := msg.Save(); err != nil {
			http.Error(w, "Failed to send message", http.StatusInternalServerError)
			return
		}

		// respon with success msg
		w.Header().Set("Content-Type", "aplication/json")
		json.NewEncoder(w).Encode(map[string] interface{} {
			"status" : "success",
			"message" : "message send successfully",
			"data" : msg,
		} )
	}) (w,r)
}

// getMsg
func (c * chatController )getMsg (w http.ResponseWriter, r * http.Request) {
	// use middleware
	middleware.SocketAuthMiddleware(func (w http.ResponseWriter, r * http.Request)  {
		userID, err := strconv.Atoi(r.Header.Get("UserID"))

		if err != nil {
			http.Error(w, "Invalid  user id", http.StatusUnauthorized)
			return
		}

		receiverIDStr := r.URL.Query().Get("receiver_id")
		receiverID, err := strconv.Atoi(receiverIDStr)
		if err != nil || receiverID == 0 {
			http.Error(w, "Invalid receiver ID", http.StatusBadRequest)
			return
		}

		// Get the chat history between the user and the receiver
		messages, err := model.GetChatHistory(userID, receiverID)
		if err != nil {
			http.Error(w, "Failed to retrieve chat history", http.StatusInternalServerError)
			return
		}

		// respond with chat
		w.Header().Set("Content-Type","aplication/json")
		json.NewEncoder(w).Encode(map[string]interface{} {
			"status" : "Success",
			 "message" : messages,
		})	
	}) (w,r)
}