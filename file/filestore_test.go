package file

import (
	"io"
	"strings"
	"testing"
	"time"

	"github.com/nrolans/configstore"
)

func TestStore(t *testing.T) {

	// Prepare a new entry
	e := configstore.NewEntry()
	e.Name = "hello.example.org"
	e.Date = time.Now()
	config := strings.NewReader("Hello Foo!")
	_, err := io.Copy(e, config)
	if err != nil {
		t.Fatal("Failed to copy to config to entry")
	}

	// Initiate a store and write to it
	fs := NewFileStore("/var/tmp/data", "20060102-150405")
	if err := fs.Add(*e); err != nil {
		t.Fatalf("Failed to write to filestore: %s", err)
	}

}
