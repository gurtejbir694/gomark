package renderer

import (
	"bytes"
	"embed"
	"html/template"
	"os"

	"github.com/gurtejbir694/markdown-to-html/pkg/config"
)

//go:embed assets/template.html assets/style.css
var assets embed.FS

func FullHTML(mdHTML []byte, opts *config.Options) ([]byte, error) {
	tmpl, err := template.ParseFS(assets, "assets/template.html")
	if err != nil {
		return nil, err
	}

	css := getCSS(opts.CSSPath)

	data := struct {
		Title string
		Body  template.HTML
		CSS   template.CSS
	}{
		Title: opts.Title,
		Body:  template.HTML(mdHTML),
		CSS:   template.CSS(css),
	}

	var out []byte
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, data); err != nil {
		return nil, err
	}
	out = buf.Bytes()
	return out, nil
}

func getCSS(path string) string {
	if path != "" {
		if b, err := os.ReadFile(path); err == nil {
			return string(b)
		}
	}
	if b, err := assets.ReadFile("assets/style.css"); err == nil {
		return string(b)
	}
	return ""
}
