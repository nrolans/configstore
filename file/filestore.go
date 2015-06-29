package file

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"time"

	"github.com/nrolans/configstore"
)

// Sensible default for date format
const (
	DefaultDateFormat = "20060102-150405"
)

// FileFileStore stores entries under <root>/<name>/<date>
type FileStore struct {
	root       string
	dateFormat string
}

func NewFileStore(root string, dateFormat string) *FileStore {
	f := new(FileStore)
	f.root = root
	f.dateFormat = dateFormat
	return f
}

func (f FileStore) Add(e configstore.Entry) error {
	err := os.Mkdir(path.Join(f.root, e.Name), 0700)
	if err != nil {
		return err
	}
	fo, err := os.Create(f.getPath(e.Name, e.Date))
	if err != nil {
		return err
	}
	defer fo.Close()

	// Write to file
	_, err = io.Copy(fo, &e)
	if err != nil {
		return err
	}

	return nil
}

func (f FileStore) Delete(e configstore.Entry) error {
	return os.Remove(f.getPath(e.Name, e.Date))
}

func (f FileStore) Names() ([]string, error) {
	names := make([]string, 0)
	dirEnts, err := ioutil.ReadDir(f.root)
	if err != nil {
		return names, err
	}

	for _, e := range dirEnts {
		// Skip non-files
		if !e.Mode().IsDir() {
			continue
		}

		names = append(names, e.Name())
	}

	sort.Sort(sort.StringSlice(names))

	return names, nil
}

func (f FileStore) Dates(name string) ([]time.Time, error) {
	dates := make([]time.Time, 0)
	dirEnts, err := ioutil.ReadDir(path.Join(f.root, name))
	if err != nil {
		return dates, err
	}

	for _, e := range dirEnts {
		// Skip non-files
		if !e.Mode().IsRegular() {
			continue
		}

		// Skip invalid date format
		date, err := time.Parse(f.dateFormat, e.Name())
		if err != nil {
			continue
		}

		dates = append(dates, date)

	}

	sort.Sort(TimeSlice(dates))

	return dates, nil
}

func (f FileStore) String() string {
	return fmt.Sprintf("Filestore in <%s> with date format <%s>", f.root, f.dateFormat)
}

func (f FileStore) Get(e *configstore.Entry) error {
	fi, err := os.Open(f.getPath(e.Name, e.Date))
	if err != nil {
		return nil
	}
	defer fi.Close()
	_, err = io.Copy(e, fi)
	if err != nil {
		return err
	}
	return nil
}

func (f FileStore) getPath(name string, date time.Time) string {
	fullpath := path.Join(f.root, name, date.Format(f.dateFormat))
	return fullpath
}

// Implement sort.Interface for slices of time.TIme
type TimeSlice []time.Time

func (t TimeSlice) Len() int           { return len(t) }
func (t TimeSlice) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t TimeSlice) Less(i, j int) bool { return t[i].After(t[j]) }
