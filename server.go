package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"github.com/gorilla/mux"
)

type User struct {
	Name  string `json:"name"`
	ID_login string `json:"ID_login"`
}

var mu sync.Mutex

var usuarios []User

func main() {
	r := mux.NewRouter()

	// endpoint para obtener la información
	r.HandleFunc("/agregar_usuario", func(w http.ResponseWriter, r *http.Request) {
		var user User
		json.NewDecoder(r.Body).Decode(&user) // unmarshaler para JSON
		fmt.Fprintf(w, "%s tienes el id: %s ", user.Name, user.ID_login)
		mu.Lock()
		usuarios = append(usuarios, user)
		mu.Unlock()
	})

	// Endpoint para agregar un usuario
	r.HandleFunc("/obtener_usuario", func(w http.ResponseWriter, r *http.Request) {
		ID_Login_var := r.URL.Query().Get("ID_Login")
		mu.Lock()
		defer mu.Unlock()

		for _, user := range usuarios {
			if user.ID_login == ID_Login_var {
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(user)
				return
			}
		}

	http.Error(w, "Usuario no encontrado", http.StatusNotFound)
	})
	 // Lógica para poder recibir peticiones en http
	http.ListenAndServe(":1108", r)
	fmt.Println("Se inicio el servidor en el puerto:1108") // avisar que se inicio el servidor y en que puerto

}
/*
# Una vez aprendido hagan un pequeño proyecto

1. Crear un modulo de go, será nuestro principal repositorio de trabajo
2. Crear un archivo llamado server.go
3. En el server.go
    1. Agregar lógica para recibir peticiones http
    2. Tener un endpoint para agregar un usuario
    3. Otro endpoint para obtener la información
    4. La información que se mande y que se reciba tiene que empaquetarse en un JSON
*/