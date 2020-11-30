package aws

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/katallaxie/protoc-gen-cloud-proxy/internal/gen"
)

type loader struct {
}

// New ...
func New() gen.Loader {
	l := new(loader)

	return l
}

// Load ...
func (l *loader) Load(modelPaths []string) (gen.APIs, error) {
	apis := gen.APIs{}

	return apis, nil
}

func loadAPI(modelPath string) (*api, error) {
	a := &api{}

	modelFile := filepath.Base(modelPath)
	modelDir := filepath.Dir(modelPath)

	err := attachModelFiles(modelDir,
		modelLoader{modelFile, a.Attach, true},
		// modelLoader{"paginators-1.json", a.AttachPaginators, false},
		// modelLoader{"waiters-2.json", a.AttachWaiters, false},
	)
	if err != nil {
		return nil, err
	}

	return a, nil
}

type modelLoader struct {
	Filename string
	Loader   func(string) error
	Required bool
}

func attachModelFiles(modelPath string, modelFiles ...modelLoader) error {
	for _, m := range modelFiles {
		filepath := filepath.Join(modelPath, m.Filename)
		_, err := os.Stat(filepath)
		if os.IsNotExist(err) && !m.Required {
			continue
		} else if err != nil {
			return fmt.Errorf("failed to load model file %v, %v", m.Filename, err)
		}

		if err = m.Loader(filepath); err != nil {
			return fmt.Errorf("model load failed, %s, %v", modelPath, err)
		}
	}

	return nil
}

// ExpandModelGlobPath ...
func ExpandModelGlobPath(globs ...string) ([]string, error) {
	modelPaths := []string{}

	for _, g := range globs {
		filepaths, err := filepath.Glob(g)
		if err != nil {
			return nil, err
		}

		for _, p := range filepaths {
			modelPaths = append(modelPaths, p)
		}
	}

	return modelPaths, nil
}

// TrimModelServiceVersions ...
func TrimModelServiceVersions(modelPaths []string) (include, exclude []string) {
	sort.Strings(modelPaths)

	// Remove old API versions from list
	m := map[string]struct{}{}
	for i := len(modelPaths) - 1; i >= 0; i-- {
		// service name is 2nd-to-last component
		parts := strings.Split(modelPaths[i], string(filepath.Separator))
		svc := parts[len(parts)-3]

		if _, ok := m[svc]; ok {
			// Removed unused service version
			exclude = append(exclude, modelPaths[i])
			continue
		}
		include = append(include, modelPaths[i])
		m[svc] = struct{}{}
	}

	return include, exclude
}
