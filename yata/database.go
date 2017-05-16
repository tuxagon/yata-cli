package yata

// Database TODO docs
type Database interface {
	Read(v interface{}) error
	Placeholder()
}

// NewDatabase TODO docs
func NewDatabase(dir string) Database {
	return NewJSONDatabase(dir)
}
