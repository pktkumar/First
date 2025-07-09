package main

import (
	"fmt"
	//"io"
	"net/http"
)

///go mod init First
///go run .

func handlerOne(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is server on port 8083")
}
func handlerTwo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is server on port 8084")
}

func myNew() {
	fmt.Println("in myNew")
	resp, err := http.Get("http://localhost:8080/name?myName=Kumar")

	if err != nil {
		fmt.Printf("error reading reponse")
		return
	}

	defer resp.Body.Close()

	if err != nil {
		fmt.Printf("error reading body response")
		return
	}

	mux1 := http.NewServeMux()
	mux1.HandleFunc("/", handlerOne)

	go func() {
		fmt.Println("in 40....")

		http.ListenAndServe(":8083", mux1)

	}()

}

func myTwo() {

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", handlerTwo)

	go func() {

		fmt.Println("in 57....")

		http.ListenAndServe(":8084", mux2)
	}()

}
