package models

import (
	"log"
	"strconv"
	"time"

	"github.com/sonyarouje/simdb/db"
)

type Customer struct {
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Path        string    `json:"path"`
	CoverLetter string    `json:"coverletter"`
	WebAddress  string    `json:"webaddress"`
	Like        bool      `json:"like"`
	UTS         time.Time `json:"uts"`
	IP          string    `json:"IP"`
	Location    string    `json:"location"`
	Reviews     []Review  `json:"reviews"`
	Ratings     []Rating  `json:"ratings"`
}

type Review struct {
	Content string `json:"content"`
	By      string `json:"by"`
}

type Rating struct {
	Rating int    `json:"rating"`
	By     string `json:"by"`
}
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Approved bool   `json:approved`
}

func (c Customer) ID() (jsonField string, value interface{}) {
	value = c.Email
	jsonField = "email"
	return
}

func (u User) ID() (jsonField string, value interface{}) {
	value = u.Email
	jsonField = "email"
	return
}

//AddComment function adds a comment to the profile
func AddComment(email string, comment string, by string) bool {
	var customers []Customer
	driver, err := db.New("data")
	if err != nil {
		panic(err)
	}

	err = driver.Open(Customer{}).Where("email", "=", email).Get().AsEntity(&customers)
	if len(customers) > 0 {
		customer := customers[0]
		var review Review
		review.By = by
		review.Content = comment
		customer.Reviews = append(customer.Reviews, review)
		err = driver.Update(customer)
		if err != nil {
			log.Println(err)
		}
		return true
	}
	return false

}

//AddRating function adds a rating to the profile
func AddRating(email string, comment string, by string) bool {
	var customers []Customer
	driver, err := db.New("data")
	if err != nil {
		panic(err)
	}

	err = driver.Open(Customer{}).Where("email", "=", email).Get().AsEntity(&customers)
	if len(customers) > 0 {
		customer := customers[0]
		var rating Rating
		rating.By = by
		rating.Rating, _ = strconv.Atoi(comment)
		customer.Ratings = append(customer.Ratings, rating)
		err = driver.Update(customer)
		if err != nil {
			log.Println(err)
		}
		return true
	}
	return false

}

//SaveReview Function to save the details in SimDB
func SaveReview(name string, email string, path string, coverletter string, webaddress string, like bool, IP string, location string) bool {

	driver, err := db.New("data")
	if err != nil {
		panic(err)
	}
	customer := Customer{
		Name:        name,
		Email:       email,
		Path:        path,
		CoverLetter: coverletter,
		WebAddress:  webaddress,
		Like:        like,
		IP:          IP,
		UTS:         time.Now(),
		Location:    location,
	}
	err = driver.Insert(customer)
	if err != nil {
		return false
	}
	return true
}
func GetAllProfiles() []Customer {
	var customers []Customer
	driver, err := db.New("data")
	if err != nil {
		panic(err)
	}
	err = driver.Open(Customer{}).Get().AsEntity(&customers)
	return customers

}
func GetProfileById(email string) Customer {
	var customer []Customer
	driver, err := db.New("data")
	if err != nil {
		log.Println(err)
	}
	err = driver.Open(Customer{}).Where("email", "=", email).Get().AsEntity(&customer)

	return customer[0]

}

//CheckEmail function checks if any new email is present
func CheckEmail(email string) bool {
	var customers []Customer
	driver, err := db.New("data")
	if err != nil {
		log.Println(err)
	}
	err = driver.Open(Customer{}).Where("email", "=", email).Get().AsEntity(&customers)
	if err != nil {
		log.Println(err)
	}
	if len(customers) == 0 {
		return false
	}
	log.Println("found")
	return true
}

//ValidateEmail function checks if an email and password match
func ValidateEmail(email string, pass string) (bool, string) {
	var user []User
	driver, err := db.New("data")
	if err != nil {
		log.Println(err)
	}
	err = driver.Open(User{}).Where("email", "=", email).Get().AsEntity(&user)
	if err != nil {
		log.Println(err)
	}
	if len(user) == 0 {
		return false, "No Email Found"
	}
	if user[0].Password == pass {
		return true, "Match"
	}
	return false, "Do Not Match"

}

//Insert User
func InsertUser(email string, pass string) (bool, string) {
	boolean, result := ValidateEmail(email, pass)
	if boolean || result == "Do Not Match" {
		return false, "Already Exists"
	}
	driver, err := db.New("data")
	if err != nil {
		panic(err)
	}
	user := User{
		Email:    email,
		Password: pass,
		Approved: true,
	}
	err = driver.Insert(user)
	if err != nil {
		return false, "Could Not Insert"
	}
	return true, "Inserted Successfully"

}
