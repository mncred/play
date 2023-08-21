package logger

// Field provides a key/value pair to attach with logs
type Field struct {
	Key   string
	Value any
}

const (
	FieldKeyName = "name"
)
