package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		_, err := fmt.Fprintf(responseWriter, "ParseForm() err: %v", err)
		if err != nil {
			return
		}
		return
	}

	name := request.FormValue("name")
	reason := request.FormValue("reason")

	if name == "" || reason == "" {
		_, err := fmt.Fprintf(responseWriter, "REJECTED!")
		if err != nil {
			return
		}
		return
	}

	_, err := fmt.Fprintf(responseWriter, "WELCOME!\n%s You've successfully apply to the Goddess Makima Fan Club!\n", name)
	if err != nil {
		return
	}
	_, err = fmt.Fprintf(responseWriter, "Your reason to join \"%s\" shall be embraced by our Goddess", reason)
	if err != nil {
		return
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
