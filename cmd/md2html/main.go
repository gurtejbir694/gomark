package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gurtejbir694/gomark/internal/parser"
	"github.com/gurtejbir694/gomark/internal/renderer"
	"github.com/gurtejbir694/gomark/pkg/config"
)

func main() {
	var css, title string
	flag.StringVar(&css, "css", "", "Path to custom CSS")
	flag.StringVar(&title, "title", "Markdown Document", "HTML page title")
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		fmt.Println(os.Stderr, "Usage: md2html [flags] <input.md> <output.html>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	inPath, outPath := args[0], args[1]

	f, err := os.Open(inPath)
	check(err)
	defer f.Close()

	bodyHTML, err := parser.Parser(f)
	check(err)

	opts := &config.Options{
		InputPath:  inPath,
		OutputPath: outPath,
		CSSPath:    css,
		Title:      title,
	}
	fullHTML, err := renderer.FullHTML(bodyHTML, opts)
	check(err)

	check(os.WriteFile(outPath, fullHTML, 0644))
	fmt.Printf("Converted: %s -> %s\n", inPath, outPath)
}

func check(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
