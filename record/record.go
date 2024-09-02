// record/record.go
package record

import (
    "bytes"
    "encoding/gob"
)

// Record representa un registro en el sistema de log.
type Record struct {
    Value  []byte
    Offset uint64 // Offset opcional, puede usarse dependiendo de la implementación del índice.
}

// Serialize convierte el Record en una secuencia de bytes para su almacenamiento.
func (r *Record) Serialize() ([]byte, error) {
    var buf bytes.Buffer
    enc := gob.NewEncoder(&buf)
    err := enc.Encode(r)
    if err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}

// Deserialize convierte una secuencia de bytes en un Record.
func Deserialize(data []byte) (*Record, error) {
    var r Record
    buf := bytes.NewReader(data)
    dec := gob.NewDecoder(buf)
    err := dec.Decode(&r)
    if err != nil {
        return nil, err
    }
    return &r, nil
}
