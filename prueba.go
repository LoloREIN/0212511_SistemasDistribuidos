// Lorenzo Reinoso Fuentes 0212511

// Declaración del paquete principal
package main

// Zona de imports
import (
	"encoding/json" // Para manipular JSON
	"fmt"           // Para imprimir mensajes
	"log"           // Para registrar errores
	"net/http"      // Para manejar el servidor HTTP
	"strconv"       // Para convertir strings a enteros
	"sync"          // Para sincronización en concurrencia
)

// Definición de la estructura Log
type Log struct {
	mu      sync.Mutex
	records []Record
}

// Definición de la estructura Record
type Record struct {
	Value  string `json:"value"` // hice cambio por string para probarlo mas sencillo jajajXD
	Offset uint64 `json:"offset"`
}

// Instancia del log con nuevo nombre
var simpleLog = &Log{}

// Endpoint que escribe
func writeHandler(w http.ResponseWriter, r *http.Request) {
	// Step 1: Unmarshal JSON a estructura
	var record Record
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Step 2: Procesar el dato
	simpleLog.mu.Lock()
	record.Offset = uint64(len(simpleLog.records))
	simpleLog.records = append(simpleLog.records, record)
	simpleLog.mu.Unlock()

	// Step 3: Marshal el resultado y enviar respuesta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]uint64{"offset": record.Offset})
}

// Endpoint que lee
func readHandler(w http.ResponseWriter, r *http.Request) {
	// Obtener parámetro offset
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
	simpleLog.mu.Lock()
	defer simpleLog.mu.Unlock()

	if offset >= uint64(len(simpleLog.records)) {
		http.Error(w, "Offset out of range", http.StatusBadRequest)
		return
	}

	record := simpleLog.records[offset]

	// Enviar respuesta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}

// Función principal
func main() {
	http.HandleFunc("/write", writeHandler) // invocación de escritura
	http.HandleFunc("/read", readHandler)   // invoación de lectura

	fmt.Println("Se inicio el servidor en el puerto:8080") // avisar que se inicio el servidor y en que puerto
	log.Fatal(http.ListenAndServe(":8080", nil))
}
