package main

import (
	"fmt"
	"net/http"
)

func DemoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Go museum server - Demo! 👋"))
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", DemoHandler)

	err := http.ListenAndServe(":3333", server) // typically we use an IP Address or a Domain:Port
	if err != nil {
		fmt.Printf("Error while running server with value %v \n,", err)
	}

}
