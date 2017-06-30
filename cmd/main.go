package main

import (
	"fmt"
	"os"

	"github.com/get-ion/faq/cmd/gen"
)

func main() {
	var (
		src = "C:\\mygopath\\src\\github.com\\get-ion\\ion\\_examples\\"
		dst = "C:\\mygopath\\src\\github.com\\get-ion\\faq\\result\\"
	)

	if len(os.Args) == 3 {
		args := os.Args[1:]

		src = args[0]
		dst = args[1]
	}

	g := gen.New(src, dst)
	if err := g.Examples().Save(); err != nil {
		panic(err)
	}
}

func help() {
	fmt.Println("Usage: go run main.go ./_examples ./_examples_md")
}
