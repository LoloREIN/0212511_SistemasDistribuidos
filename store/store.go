package store


import (
    "encoding/binary"
    "os"
    "sync"
    "0212511_SD/record"
)

// Store gestiona el almacenamiento físico para los registros de log.
type Store struct {
    file *os.File
    mu   sync.Mutex
    size uint64 // Mantiene un seguimiento del tamaño actual del archivo para optimizar las escrituras.
}

// NewStore crea y retorna una nueva instancia de Store.
func NewStore(filePath string) (*Store, error) {
    f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        return nil, err
    }

    // Obtener el tamaño actual del archivo para seguir agregando al final.
    fi, err := f.Stat()
    if err != nil {
        return nil, err
    }

    return &Store{
        file: f,
        size: uint64(fi.Size()),
    }, nil
}

// Append agrega un nuevo registro al final del archivo de log.
func (s *Store) Append(rec *record.Record) (uint64, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    // Asegúrate de que los registros se escriban de manera que puedan ser leídos de manera eficiente.
    pos := s.size
    data, err := rec.Serialize()  // Suponiendo que record tiene un método Serialize.
    if err != nil {
        return 0, err
    }
    dataLen := uint64(len(data))
    buf := make([]byte, dataLen+8) // 8 extra bytes para almacenar la longitud del registro.

    binary.BigEndian.PutUint64(buf[:8], dataLen)
    copy(buf[8:], data)

    _, err = s.file.Write(buf)
    if err != nil {
        return 0, err
    }

    s.size += uint64(len(buf))
    return pos, nil
}

// Read lee un registro desde una posición específica en el archivo.
func (s *Store) Read(pos uint64) (*record.Record, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    var size uint64
    sizeBuf := make([]byte, 8)
    _, err := s.file.ReadAt(sizeBuf, int64(pos))
    if err != nil {
        return nil, err
    }

    size = binary.BigEndian.Uint64(sizeBuf)
    data := make([]byte, size)
    _, err = s.file.ReadAt(data, int64(pos+8))
    if err != nil {
        return nil, err
    }

    return record.Deserialize(data)  // Suponiendo que record tiene un método Deserialize.
}

// Close cierra el archivo de log.
func (s *Store) Close() error {
    return s.file.Close()
}
