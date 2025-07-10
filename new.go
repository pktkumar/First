package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

//
///go mod init First
///go run .

func myNew() {

	mux1 := http.NewServeMux()
	mux1.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		resp, err := http.Get("http://localhost:8080/name?myName=Kumar")

		if err != nil {
			fmt.Printf("error reading reponse")
			return
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read response body", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "%s", body)
	})
	log.Println("Server started on :8083")
	http.ListenAndServe(":8083", mux1)
	//log.Fatal(http.ListenAndServe(":8083", mux1))

}

func myTwo() {

	mux2 := http.NewServeMux()

	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "This is server on port 8084 <br> test.....")
	})
	log.Println("Server started on :8084")
	http.ListenAndServe(":8084", mux2)
	//log.Fatal(http.ListenAndServe(":8084", mux2))
}
