// Lorenzo Reinoso Fuentes                      0212511

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