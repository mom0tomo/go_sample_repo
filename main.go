package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("request received")
		fmt.Println(w, "Hello, 世界")
	})

	log.Println("start server")
	server := &http.Server{Addr: ":8080"}
	server.ListenAndServe()
}
