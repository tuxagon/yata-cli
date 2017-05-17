package yata

// Database TODO docs
type Database interface {
	Read(collection string, v interface{}) error
	Write(collection string, v interface{}) error
}

// NewDatabase TODO docs
func NewDatabase(dir string) Database {
	return NewJSONDatabase(dir)
}
