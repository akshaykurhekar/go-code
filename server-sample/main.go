package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter,r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "hello Akshay from server...")
}

func formHandler(w http.ResponseWriter ,r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful \n")
	name := r.FormValue("name") 
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name) 
	fmt.Fprintf(w, "Address = %s\n", address)
}

func main()  {
	
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer);
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server is live on port 8080\n")
	if err := http.ListenAndServe(":8080",nil); err != nil {
		log.Fatal(err)
	}
}