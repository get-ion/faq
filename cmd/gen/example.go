package gen

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Snippet struct {
	filename string
	contents []byte
}

// DocName returns the markdown simple name
func (s Snippet) DocName() string {
	newExt := ".md"
	if ext := filepath.Ext(s.filename); ext != "" {
		return strings.Replace(s.filename, ext, newExt, 1)
	}

	return s.filename + newExt
}

func (s Snippet) Contents() []byte {
	ext := filepath.Ext(s.filename)
	c := s.contents
	var langPrefix = "bash"
	if ext != "" {
		langPrefix = ext[1:]
	}

	// ```go
	// [...]
	// ```
	c = append([]byte("```"+langPrefix+"\n"), c...)
	c = append(c, []byte("\n```")...)
	return c
}

type SnippetMap struct {
	RelFolder string
	Snippets  []Snippet
}

func (s *SnippetMap) Contents() []byte {
	var c []byte
	for _, sn := range s.Snippets {
		c = append(c, append(sn.Contents(), []byte("\n")...)...)
	}

	return c
}

func (s *SnippetMap) DocName() string {
	return strings.Replace(s.RelFolder, string(filepath.Separator), "", -1) + ".md"
}

type ExampleGen struct {
	Dir     string
	Skipper func(path string, info os.FileInfo) bool // true to skip
}

var excludes = []string{"README.md", ".zip", ".ico", ".png", ".jpg", "jquery-2.1.1.js", "bootstrap.min.css", "bindata.go"}

var DefaultSkipper = func(path string, info os.FileInfo) bool {
	if info == nil || info.IsDir() {
		return true
	}

	for _, ex := range excludes {
		if strings.HasSuffix(path, ex) {
			return true
		}
	}

	return false
}

func (e *ExampleGen) Snippets() ([]Snippet, error) {
	root := e.Dir
	var snippets = make([]Snippet, 0)
	err := filepath.Walk(root, func(path string, info os.FileInfo, errW error) error {
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		contents, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		s := Snippet{
			filename: rel,
			contents: contents,
		}

		snippets = append(snippets, s)
		return errW
	})

	return snippets, err
}
