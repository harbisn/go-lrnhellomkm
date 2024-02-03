package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if error := request.ParseForm(); error != nil {
		fmt.Fprintf(responseWriter, "ParseForm() error: %v", error)
		return
	}

	name := request.FormValue("name")
	reason := request.FormValue("reason")

	if name == "" || reason == "" {
		fmt.Fprintf(responseWriter, "REJECTED!")
		return
	}

	fmt.Fprintf(responseWriter, "WELCOME!\n%s You've successfully apply to the Goddess Makima Fan Club!\n", name)
	fmt.Fprintf(responseWriter, "Your reason to join \"%s\" shall be embraced by our Goddess", reason)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at port 8080\n")
	if error := http.ListenAndServe(":8080", nil); error != nil {
		log.Fatal(error)
	}
}
