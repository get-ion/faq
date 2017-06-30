package gen

import (
	"os"
	"path/filepath"
)

type Gen struct {
	src string
	dst string
}

func New(src, dst string) *Gen {
	return &Gen{
		src: src,
		dst: dst,
	}
}

func (g *Gen) Examples() GenResult {
	root := g.src
	var examples = make([]ExampleGen, 0)
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			ex := ExampleGen{Dir: root, Skipper: DefaultSkipper}
			examples = append(examples, ex)

			return filepath.SkipDir
		}
		return nil
	})

	return GenResult{
		dst:      g.dst,
		examples: examples,
	}
}

type GenResult struct {
	dst      string
	examples []ExampleGen
}

func (r GenResult) Get() []ExampleGen {
	return r.examples
}

func (r GenResult) Save() error {
	examples := r.examples
	os.RemoveAll(r.dst)

	for _, e := range examples {
		snippets, err := e.Snippets()
		if err != nil {
			return err
		}

		for _, s := range snippets {
			filename := s.DocName()
			fullname := filepath.Join(r.dst, filename)
			dir := filepath.Dir(fullname)

			os.MkdirAll(dir, os.ModePerm)

			f, err := os.Create(fullname)
			if err != nil {
				return err
			}

			_, err = f.Write(s.Contents())
			if err != nil {
				return err
			}

			// Sync, Close or ioutil.WriteFile skips a file if has the same prefix I don't know why.
		}

	}

	return nil
}
