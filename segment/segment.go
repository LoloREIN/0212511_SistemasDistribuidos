// segment/segment.go
package segment

import (
	"fmt"
	"path"
	"0212511_SD/record"
	"0212511_SD/store"
	"0212511_SD/index"
	"0212511_SD/config"
)

type Segment struct {
	Store                  *store.Store
	Index                  *index.Index
	BaseOffset, NextOffset uint64
	Config                 config.Config
}

// Asigna correctamente la ruta como string, y maneja la creación del objeto Store/Index internamente
func NewSegment(dir string, baseOffset uint64, c config.Config) (*Segment, error) {
	s := &Segment{
		BaseOffset: baseOffset,
		Config:     c,
	}

	storeFilePath := path.Join(dir, fmt.Sprintf("%d.store", baseOffset))
	indexFilePath := path.Join(dir, fmt.Sprintf("%d.index", baseOffset))

	var err error
	s.Store, err = store.NewStore(storeFilePath)
	if err != nil {
		return nil, err
	}

	s.Index, err = index.NewIndex(indexFilePath)
	if err != nil {
		return nil, err
	}

	// Encuentra el último índice para establecer el siguiente offset
	// Debe usar un método adecuado para obtener el último índice si está disponible
	if lastIndexOffset, err := s.Index.GetLastIndexOffset(); err == nil {
		s.NextOffset = baseOffset + lastIndexOffset + 1
	} else {
		s.NextOffset = baseOffset // Si no hay índices, comienza desde BaseOffset
	}

	return s, nil
}

func (s *Segment) Append(rec *record.Record) (uint64, error) {
	offset, err := s.Store.Append(rec)
	if err != nil {
		return 0, err
	}
	err = s.Index.Write(offset, rec.Offset)
	if err != nil {
		return 0, err
	}
	return offset, nil
}

func (s *Segment) Read(offset uint64) (*record.Record, error) {
	return s.Store.Read(offset)
}
