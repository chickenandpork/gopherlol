package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chickenandpork/gopherlol/commands"
)

// func handler() moved to commands/handler.go to allow unittesting against it

func main() {
	http.HandleFunc("/", commands.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
