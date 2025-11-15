module github.com/gurtejbir694/markdown-to-html/cmd/md2html

go 1.25.3

require (
	github.com/yourname/markdown-to-html/internal/parser v0.0.0
	github.com/yourname/markdown-to-html/internal/renderer v0.0.0
	github.com/yourname/markdown-to-html/pkg/config v0.0.0
)

replace github.com/yourname/markdown-to-html/internal/parser => ../../internal/parser

replace github.com/yourname/markdown-to-html/internal/renderer => ../../internal/renderer

replace github.com/yourname/markdown-to-html/pkg/config => ../../pkg/config
