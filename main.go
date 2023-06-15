package main

import (
	"fmt"
	"log"
	"net/http"
	"veda-backend/handlers"
	vedalan "veda-backend/vedaLan"
)

//"log"
//"net/http"
//"veda-backend/handlers"

type Quote struct {
	Quote  string `json:"text"`
	Author string `json:"author"`
}

type Usage struct {
	SessionID int `json:"sessionID"`
	FromTime  int `json:"fromTime"`
	ToTime    int `json:"toTime"`
}



func main() {

	fmt.Println("test")

	http.HandleFunc("/quote", handlers.HandleQuoteRequest)
	http.HandleFunc("/readusage", handlers.ReadUsageHandler)
	http.HandleFunc("/readtodo", handlers.ReadTodoHandler)
	http.HandleFunc("/deletetodo", handlers.DeleteTodoHandler)
	http.HandleFunc("/createtodo", handlers.CreateTodoHandler)
	http.HandleFunc("/logout", handlers.AddLogoutTimeHandler)
	http.HandleFunc("/login", handlers.AddLoginTimeHandler)
	http.HandleFunc("/getothers", vedalan.GetPeersJson)



	
	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", nil))



}

