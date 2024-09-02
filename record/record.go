// record/record.go
package record

type Record struct {
    Value  []byte
    Offset uint64
}

func NewRecord(value []byte, offset uint64) *Record {
    return &Record{
        Value:  value,
        Offset: offset,
    }
}
