package configstore

import (
	"bytes"
	"fmt"
	"time"
)

// Entry including Name, Date and Content
type Entry struct {
	Name    string
	Date    time.Time
	Content bytes.Buffer
}

// NewEntry Create a new Entry to add in the store
func NewEntry() *Entry {
	e := new(Entry)
	return e
}

func (b *Entry) Write(p []byte) (n int, err error) {
	return b.Content.Write(p)
}

func (b *Entry) Read(p []byte) (n int, err error) {
	return b.Content.Read(p)
}

func (b Entry) String() string {
	return fmt.Sprintf("%s @ %s", b.Name, b.Date)
}

// Custom error types

// NotFound in the store
type NotFound struct {
	Name string
	Date time.Time
}
