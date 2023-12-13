package file

import (
	"io/ioutil"
	"path/filepath"

	file "github.com/eric-tech01/simple-file"
)

// fileDataSource file provider.
type fileDataSource struct {
	path        string
	dir         string
	enableWatch bool
	changed     chan struct{}
}

// NewDataSource returns new fileDataSource.
func NewDataSource(path string, watch bool) (*fileDataSource, error) {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return nil, err

	}
	dir := file.CheckAndGetParentDir(absolutePath)
	ds := &fileDataSource{path: absolutePath, dir: dir, enableWatch: watch}
	return ds, nil
}

// ReadConfig ...
func (fp *fileDataSource) ReadConfig() (content []byte, err error) {
	return ioutil.ReadFile(fp.path)
}

// Close ...
func (fp *fileDataSource) Close() error {
	close(fp.changed)
	return nil
}

// IsConfigChanged ...
func (fp *fileDataSource) IsConfigChanged() <-chan struct{} {
	return fp.changed
}
