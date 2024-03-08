package main

import (
	"fmt"
	"log"
	"net/http"

	"UTS/Controller"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/getallrooms/{id}", Controller.GetAllRooms).Methods("GET")
	router.HandleFunc("/getdetailrooms/{id}", Controller.GetDetailRooms).Methods("GET")
	router.HandleFunc("/insertrooms/", Controller.InsertRoom).Methods("POST")

	http.Handle("/", router)
	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
