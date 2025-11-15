module github.com/gurtejbir694/gomark/cmd/md2html

go 1.25.3

require (
	github.com/gurtejbir694/gomark/internal/parser v0.0.0
	github.com/gurtejbir694/gomark/internal/renderer v0.0.0
	github.com/gurtejbir694/gomark/pkg/config v0.0.0
)

replace github.com/gurtejbir694/gomark/internal/parser => ../../internal/parser

replace github.com/gurtejbir694/gomark/internal/renderer => ../../internal/renderer

replace github.com/gurtejbir694/gomark/pkg/config => ../../pkg/config
