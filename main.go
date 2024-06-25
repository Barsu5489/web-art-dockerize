package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// The server runs asynchronously on port 8080 using a goroutine.
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Println(err)
		}
	}()
	fmt.Println("server running on port 8080")
	select {} // keep the main function running indefinitely.
}
