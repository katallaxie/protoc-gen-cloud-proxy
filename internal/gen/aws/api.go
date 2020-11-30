package aws

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type api struct {
	path string
}

// Attach opens a file by name, and unmarshal its JSON data.
// Will proceed to setup the API if not already done so.
func (a *api) Attach(filename string) error {
	a.path = filepath.Dir(filename)
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(a); err != nil {
		return fmt.Errorf("failed to decode %s, err: %v", filename, err)
	}

	return nil
}
