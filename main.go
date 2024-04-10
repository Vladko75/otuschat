package main

import (
	"log"
	"net/http"
	"os"
)

const (
	port string = ":8989"
)

var dbUser *os.File
var dbMessage *os.File

func main() {
	var err error
	// Data for users
	dbUser, err = os.OpenFile("./data/users.txt",
		os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Panic("Error open file users.txt: " + err.Error())
		return
	}
	defer dbUser.Close()

	// Data for messages
	dbMessage, err = os.OpenFile("./data/messages.txt",
		os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Panic("Error open file messages.txt: " + err.Error())
		return
	}
	defer dbMessage.Close()

	http.HandleFunc("/", checkService)
	http.HandleFunc("/api/sign-in", signIn)
	//	http.HandleFunc("/ws", handleConnections)

	readAllUsers()

	log.Printf("Server started on %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Panic("Error starting server: " + err.Error())
	}

}

func checkService(w http.ResponseWriter, r *http.Request) {
	responseString(w, `{"success": true}`)
}

func responseString(w http.ResponseWriter, text string) {
	responseJson(w, []byte(text))
}

func responseJson(w http.ResponseWriter, v []byte) {
	w.Header().Set("Content-Type", "application/json;  charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(v); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Error`))
	}
}
