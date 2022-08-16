package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Rest API with GO!")

	http.HandleFunc("/status", handleStatus)
	http.HandleFunc("/", handleRoot)

	/*  BELOW IS AN EXAMPLE OF USING AN ANONYMOUS FUNCTION   */
	/*http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		log.Println("'/status' was called")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})*/

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleStatus(w http.ResponseWriter, r *http.Request) {

	log.Println("'/status' was called")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	status := map[string]string{"status": "OK"}
	json.NewEncoder(w).Encode(status)
	//w.Write([]byte("OK"))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	log.Println("'/' was called")
	w.WriteHeader(http.StatusNotImplemented)
}