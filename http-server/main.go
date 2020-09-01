package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)
// http.Request is a pointer 
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	fmt.Print(req.Header)
	fmt.Print("Here")
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, name, h)
		}
	}
}

func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}()
}

func main() {
	SetupCloseHandler()
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers",  headers)

	http.ListenAndServe(":8080", nil)
}
