package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/signup", SignupHandler).Methods("POST")
	router.HandleFunc("/login", LoginHandler).Methods("POST")
	router.HandleFunc("/protected", TokenVerifyMiddlewWare(ProtectedHandler)).Methods("GET")

	log.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func SignupHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("SignupHandler called"))
}

func LoginHandler(w http.ResponseWriter, r *http.Request){

}

func ProtectedHandler(w http.ResponseWriter, r *http.Request){

}

func TokenVerifyMiddlewWare(next http.HandlerFunc) http.HandlerFunc{
	return nil
}