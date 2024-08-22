package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/get_user", func(w http.ResponseWriter, r *http.Request) {
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		fmt.Fprintf(w, "%s %s is %d years old", user.Name, user.Email, user.Age)
	})

	r.HandleFunc("/add_user", func(w http.ResponseWriter, r *http.Request) {
		guy := User{
			Name:  "john",
			Email: "doe@doe.com",
			Age:   22,
		}

		json.NewEncoder(w).Encode(guy)
	})

	http.ListenAndServe(":8080", r)
}
