package resolvers

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

const (
	src       = "src"
	arc       = "arc"
	goPathSrc = "src"
	manifest  = ".cgentron.yml"
)

// Opts ...
type Opts struct {
	// Output ...
	Output string
	// Src ...
	Src string
	// Arc ...
	Arc string
	// GoPath ...
	GoPath string
}

// Opt ...
type Opt func(*Opts)

// NewClient ...
func NewClient(opts ...Opt) (*Client, error) {
	o := new(Opts)
	if err := o.Configure(opts...); err != nil {
		return nil, err
	}

	c := new(Client)
	c.opts = o
	c.HTTPClient = &http.Client{Timeout: 5 * time.Second}

	return c, nil
}

// Configure ...
func (s *Opts) Configure(opts ...Opt) error {
	for _, o := range opts {
		o(s)
	}

	srcRootPath := filepath.Join(filepath.FromSlash(s.Output), src)
	if err := deleteSource(srcRootPath); err != nil {
		return err
	}

	goPath, err := ioutil.TempDir(srcRootPath, "gop-*")
	if err != nil {
		return err
	}
	s.GoPath = goPath
	s.Src = filepath.Join(goPath, goPathSrc)

	arcPath := filepath.Join(filepath.FromSlash(s.Output), arc)
	if err := os.MkdirAll(arcPath, 0o755); err != nil {
		return err
	}
	s.Arc = arcPath

	return nil
}

// Client ...
type Client struct {
	HTTPClient *http.Client

	opts *Opts
}

// GoPath get the plugins GoPath.
func (c *Client) GoPath() string {
	return c.opts.GoPath
}

// Fetch a plugin archive.
func (c *Client) Fetch(ctx context.Context, name, version, archive string) (string, error) {
	filename := filepath.Join(c.opts.Arc, name, version+".zip")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, archive, nil)
	if err != nil {
		return "", err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: %d: %s", resp.StatusCode, resp.Status)
	}

	if err := os.MkdirAll(filepath.Dir(filename), 0o755); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	out, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	return "", nil
}

// Unzip extracts an archive
func (c *Client) Unzip(name, version string) error {
	if err := c.unzipArchive(name, version); err != nil {
		return err
	}

	return nil
}

// ReadManifest reads a plugin manifest.
func (c *Client) ReadManifest(moduleName string) (*Manifest, error) {
	return ReadManifest(c.opts.GoPath, moduleName)
}

// ReadManifest reads a plugin manifest.
func ReadManifest(goPath, moduleName string) (*Manifest, error) {
	p := filepath.Join(goPath, goPathSrc, filepath.FromSlash(moduleName), manifest)

	file, err := os.Open(p)
	if err != nil {
		return nil, fmt.Errorf("failed to open the plugin manifest %s: %w", p, err)
	}

	defer func() { _ = file.Close() }()

	m := &Manifest{}
	err = yaml.NewDecoder(file).Decode(m)
	if err != nil {
		return nil, fmt.Errorf("failed to decode the plugin manifest %s: %w", p, err)
	}

	return m, nil
}

func (c *Client) unzipArchive(name, version string) error {
	filename := filepath.Join(c.opts.Arc, name, version+".zip")

	archive, err := zip.OpenReader(filename)
	if err != nil {
		return err
	}
	defer archive.Close()

	dest := filepath.Join(c.opts.Src, filepath.FromSlash(name))

	for _, f := range archive.File {
		if err := unzipFile(f, dest); err != nil {
			return err
		}
	}

	return nil
}

func unzipFile(f *zip.File, dest string) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}

	defer func() { _ = rc.Close() }()

	pathParts := strings.SplitN(f.Name, string(os.PathSeparator), 2)
	p := filepath.Join(dest, pathParts[1])

	if f.FileInfo().IsDir() {
		return os.MkdirAll(p, f.Mode())
	}

	err = os.MkdirAll(filepath.Dir(p), 0o750)
	if err != nil {
		return err
	}

	elt, err := os.OpenFile(p, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}

	defer func() { _ = elt.Close() }()

	_, err = io.Copy(elt, rc)
	if err != nil {
		return err
	}

	return nil
}

func deleteSource(dir string) error {
	dirPath, err := filepath.Abs(dir)
	if err != nil {
		return err
	}

	currentPath, err := os.Getwd()
	if err != nil {
		return err
	}

	if strings.HasPrefix(currentPath, dirPath) {
		return fmt.Errorf("cannot be deleted: the directory path %s is the parent of the current path %s", dirPath, currentPath)
	}

	err = os.RemoveAll(dir)
	if err != nil {
		return err
	}

	return os.MkdirAll(dir, 0o755)
}
