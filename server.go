// Lorenzo Reinoso Fuentes 						0212511

package main

// Zona de mis imports donde
import (
	"encoding/json" // Esta me permite usar el encoding de JSON
	"fmt"           // Me permite usar print y scanf, en si es formatear Strings
	"github.com/gorilla/mux"
	"log"
	"net/http" // Uno de los requisitos es que sea un servidor HTTP
	"strconv"
	"sync"
)

type Log struct {
	mu      sync.Mutex
	records []Record
}

type Record struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"`
}

/*
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}
*/

// Crear estructura de Logs "locker de registros"
var Bitacora = &Log{} // con puntero a Log

func main() {
	r := mux.NewRouter()

	// Handler de escritura: endpoint que "escribe"
	r.HandleFunc("/write_log", func(w http.ResponseWriter, r *http.Request) {
		var record Record
		err := json.NewDecoder(r.Body).Decode(&record)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		Bitacora.mu.Lock()
		record.Offset = uint64(len(Bitacora.records))
		Bitacora.records = append(Bitacora.records, record)
		Bitacora.mu.Unlock()

		// Step 3: Marshal el resultado y enviar respuesta
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]uint64{"offset": record.Offset})
	})

	r.HandleFunc("/read_log", func(w http.ResponseWriter, r *http.Request) {
		offsetStr := r.URL.Query().Get("offset")
		if offsetStr == "" {
			http.Error(w, "Offset parameter is required", http.StatusBadRequest)
			return
		}

		offset, err := strconv.ParseUint(offsetStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid offset", http.StatusBadRequest)
			return
		}

		// Procesar el dato
		Bitacora.mu.Lock()
		defer Bitacora.mu.Unlock()

		if offset >= uint64(len(Bitacora.records)) {
			http.Error(w, "Offset out of range", http.StatusBadRequest)
			return
		}

		record := Bitacora.records[offset]

		// Enviar respuesta
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(record)
	})

	fmt.Println("Se inicio el servidor en el puerto:8080") // avisar que se inicio el servidor y en que puerto
	log.Fatal(http.ListenAndServe(":8080", r))

}

// Marshaler: convertir una estructura de datos a un JSON, XML u otro formato
// Unmarshaler: de un JSON, XML u otro convertir a una estructura de datos del sistema

/*
## Características del servidor simple

Vamos a crear un servidor que escriba y lea nuestro log simple, pero es más fácil decirlo que hacerlo.

- Para esto necesitaremos dos endpoints que tendrán que procesar lo que necesitamos en tres pasos
    - Primero Desmarshalear/Unmarshal el JSON a una estructura
    - Segundo La lógica como tal del endpoint para procesar el dato
    - Tercero Marshalear/Marshal el resultado para dar una respuesta al request
- Al final solo tendremos dos endpoints en nuestro API uno para leer y otro para escribir

/*

/*
## Mí primer servidor general para guardar commit logs y recibir JSON

- Necesito una forma de mandar mensajes
- Necesito una estructura en la cuál mandar estos mensajes

Vamos a usar JSON por peticiones HTTP ya que:

- JSON sirve para cualquier cosa
- Es fácil de estructurar a nivel código y a nivel de HTTP

Por lo que necesitamos un servidor

- Al inicio tendremos dos endpoints
- Uno que consume y otro que produce
    - Uno que escribe y otro que va a leer pue
- Nuestro servidor va a pasar una petición en JSON y la tendrá que reestructurar en un dato que podemos guardar y viceversa
    - A este proceso de transformar datos duros ya sea de red, puertos y bases de datos a una estructura declara se le conoce como **marshalling**
- Pero necesitamos una forma de estructurar nuestros datos, ya que al momento de empezar nuestro sistema principal tendremos mejor manejo de que esperar en los servicios

Y entra a nuestro escenario: →
*/
