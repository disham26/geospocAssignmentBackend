package main

import (
	"log"
	"net/http"

	controllers "geospocAssignmentBackend/controllers"

	"github.com/gorilla/mux"
)

func main() {
	STATIC_DIR := "/uploads/"
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/submitReview", controllers.SubmitReview).Methods("POST")
	router.HandleFunc("/checkEmail/{email}", controllers.CheckEmailExists).Methods("GET")
	router.HandleFunc("/checkUser/{email}/{password}", controllers.ValidateUser).Methods("GET")
	router.HandleFunc("/approval/{email}/{password}", controllers.ApproveUser).Methods("GET")
	router.HandleFunc("/getAllProfiles", controllers.GetAllProfiles).Methods("GET")
	router.HandleFunc("/getProfile/{email}", controllers.GetProfileById).Methods("GET")
	router.HandleFunc("/addComment/{emailOfProfile}/{comment}/{by}", controllers.AddComment).Methods("GET")
	router.HandleFunc("/addRating/{emailOfProfile}/{rating}/{by}", controllers.AddRating).Methods("GET")
	router.
		PathPrefix(STATIC_DIR).
		Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("."+STATIC_DIR))))
	log.Fatal(http.ListenAndServe(":8000", router))
}
