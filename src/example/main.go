package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	port = 9111
)

func httpResponseOk(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	now := time.Now()

	// Print request information.
	fmt.Println(now)
	fmt.Println(request.Proto, request.Method, request.URL.Path)
	fmt.Println(request.RemoteAddr, "➡", request.Host)
	for k, v := range request.Form {
		fmt.Println(fmt.Sprintf("%s:", k), strings.Join(v, ""))
	}
	fmt.Println()

	// Send a response.
	fmt.Fprintf(writer, "OK.\r\n")
}

func main() {
	http.HandleFunc("/", httpResponseOk)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("Opening a listener failed: ", err)
	}
}
