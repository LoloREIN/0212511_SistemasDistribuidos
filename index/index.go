// index/index.go
package index

import (
    "encoding/binary"
    "os"
    "errors"
)

type Index struct {
    file *os.File
    size uint64 // Total size of the index file in bytes.
}

// NewIndex opens or creates a new index file.
func NewIndex(filePath string) (*Index, error) {
    f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
        return nil, err
    }

    // Get the current size of the file to know where to write the next index.
    fi, err := f.Stat()
    if err != nil {
        return nil, err
    }

    return &Index{
        file: f,
        size: uint64(fi.Size()),
    }, nil
}

// Write adds a new index entry at the end of the index file.
func (idx *Index) Write(offset, position uint64) error {
    if _, err := idx.file.Seek(0, os.SEEK_END); err != nil {
        return err
    }

    buf := make([]byte, 16) // Assume each index entry is 16 bytes (8 bytes each for offset and position).
    binary.BigEndian.PutUint64(buf[0:8], offset)
    binary.BigEndian.PutUint64(buf[8:16], position)

    _, err := idx.file.Write(buf)
    return err
}

// GetLastIndexOffset retrieves the last offset recorded in the index.
func (idx *Index) GetLastIndexOffset() (uint64, error) {
    if idx.size == 0 {
        return 0, errors.New("index file is empty")
    }

    // Seek to the last index entry
    _, err := idx.file.Seek(-16, os.SEEK_END)
    if err != nil {
        return 0, err
    }

    buf := make([]byte, 8)
    if _, err := idx.file.Read(buf); err != nil {
        return 0, err
    }

    lastOffset := binary.BigEndian.Uint64(buf)
    return lastOffset, nil
}
