package controllers

import (
	"encoding/json"
	model "geospocAssignmentBackend/models"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

//ValidateResult struct is the dummy struct for all the outputs
type ValidateResult struct {
	Status  bool
	Message string
}

func GetProfileById(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	Vars := mux.Vars(r)
	email := Vars["email"]
	result := model.GetProfileById(email)
	jData, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)

}

//AddComment controller takes the input and adds it to the simdb
func AddComment(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var validateRequest ValidateResult
	Vars := mux.Vars(r)
	email := Vars["emailOfProfile"]
	comment := Vars["comment"]
	by := Vars["by"]
	result := model.AddComment(email, comment, by)
	if result {
		validateRequest.Status = true
		validateRequest.Message = "Success"
	} else {
		validateRequest.Status = false
		validateRequest.Message = "Failed"
	}
	w.Header().Set("Content-Type", "application/json")
	jData, _ := json.Marshal(validateRequest)
	w.Write(jData)

}

//AddRating controller takes the input and adds it to the simdb
func AddRating(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var validateRequest ValidateResult
	Vars := mux.Vars(r)
	email := Vars["emailOfProfile"]
	rating := Vars["rating"]
	by := Vars["by"]
	result := model.AddRating(email, rating, by)
	if result {
		validateRequest.Status = true
		validateRequest.Message = "Success"
	} else {
		validateRequest.Status = false
		validateRequest.Message = "Failed"
	}
	w.Header().Set("Content-Type", "application/json")
	jData, _ := json.Marshal(validateRequest)
	w.Write(jData)

}

//GetAllProfiles function gets all the profiles
func GetAllProfiles(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	result := model.GetAllProfiles()
	w.Header().Set("Content-Type", "application/json")
	jData, _ := json.Marshal(result)
	w.Write(jData)

}
func SubmitReview(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	r.ParseMultipartForm(32 << 20) // Parses the request body
	file, _, err := r.FormFile("file")

	if err != nil {
		log.Println("Error retrieving file", err)
	}
	defer file.Close()
	email := r.Form.Get("email")
	fileBytes, err := ioutil.ReadAll(file)
	filetype := http.DetectContentType(fileBytes)
	fileName := email
	fileEndings, err := mime.ExtensionsByType(filetype)
	newPath := filepath.Join("uploads", fileName+fileEndings[0])
	newFile, err := os.Create(newPath)
	if _, err := newFile.Write(fileBytes); err != nil {
		log.Println("File writing error is:", err)
		return
	}
	defer newFile.Close()

	coverletter := r.Form.Get("coverletter")
	boolean := r.Form.Get("boolean")

	name := r.Form.Get("name")
	ip := r.Form.Get("ip")
	location := r.Form.Get("location")
	webAddresss := r.Form.Get("webAddress")

	if err != nil {
		log.Println(err)
		return
	}
	var like bool
	if boolean == "true" {
		like = true
	} else {
		like = false
	}

	Data := "Hi"
	jData, err := json.Marshal(Data)
	if err != nil {
	}
	model.SaveReview(name, email, newPath, coverletter, webAddresss, like, ip, location)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

//CheckEmailExists is the API used to check if there is an email address already used
func CheckEmailExists(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	Vars := mux.Vars(r)
	email := Vars["email"]
	w.Header().Set("Content-Type", "application/json")
	jData, _ := json.Marshal(model.CheckEmail(email))

	w.Write(jData)

}

//ValidateUser controller
func ValidateUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	Vars := mux.Vars(r)
	email := Vars["email"]
	pass := Vars["password"]
	result, message := model.ValidateEmail(email, pass)
	var validateResult ValidateResult
	validateResult.Message = message
	validateResult.Status = result
	w.Header().Set("Content-Type", "application/json")
	jData, _ := json.Marshal(validateResult)
	w.Write(jData)
}

//ApprovUser controller
func ApproveUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	Vars := mux.Vars(r)
	email := Vars["email"]
	password := Vars["password"]
	result, message := model.InsertUser(email, password)
	var validateResult ValidateResult
	validateResult.Message = message
	validateResult.Status = result
	w.Header().Set("Content-Type", "application/json")
	jData, _ := json.Marshal(validateResult)
	w.Write(jData)

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
