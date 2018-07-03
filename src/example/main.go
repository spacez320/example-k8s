//
// A very simple Go webserver. Requires the environment variable PORT be set.
//
// Usage:
//
//   GOPATH=$PWD go install example
//   PORT=9000 ./bin/example

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	port = os.Getenv("PORT")
)

func httpResponseOk(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	now := time.Now()

	// Print request information.
	fmt.Println(request.Proto, request.Method, request.URL.Path)
	fmt.Println(now)
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

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal("Opening a listener failed: ", err)
	}
}
