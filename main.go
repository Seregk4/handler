package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//var message string

type requestBody struct {
	Message string `json:"message"`
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var reqBody requestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	messages := Message{
		Text: reqBody.Message,
	}

	if err := DB.Create(&messages).Error; err != nil {
		http.Error(w, "Unable to save message", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(messages)
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	var messages []Message
	if err := DB.Find(&messages).Error; err != nil {
		http.Error(w, "Unable to fetch messages", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages) // Возврат списка сообщений
}

func deleteMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if result := DB.Delete(&Message{}, vars["id"]); result.RowsAffected == 0 {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent) // 204 No Content
}

func main() {

	InitDB()

	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	// наше приложение будет слушать запросы на localhost:8080/api/hello
	router.HandleFunc("/api/hello", GetMessage).Methods("GET")
	router.HandleFunc("/api/hello", CreateMessage).Methods("POST")
	http.ListenAndServe(":8080", router)
}
