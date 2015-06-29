package configstore

import (
	"fmt"
	"time"
)

type Store interface {
	StoreWriter
	StoreReader
	fmt.Stringer
}

type StoreWriter interface {
	Add(Entry) error
	Delete(Entry) error
}

type StoreReader interface {
	Names() ([]string, error)
	Get(*Entry) error
	Dates(string) ([]time.Time, error)
}
