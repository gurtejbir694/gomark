# Gomark

**Fast, modular Markdown to standalone HTML in Go.**

> Converts any `.md` file into a **beautiful, self-contained HTML page** — perfect for documentation, blogs, or static sites.

```bash
gomark input.md output.html
```
## Features

| Feature | Status |
| ------- | ------ |
| GitHub Flavored Markdown (GFM) | Supported |
| Syntax Highlighting (100+ languages) | Supported |
| Embedded CSS & Template | Supported |
| Images, Links, Tables, Blockquotes | Supported |
| Zero runtime dependencies | Supported |
| `go install` -able CLI | Supported |

## Installation

### From Source

```bash
git clone https://github.com/gurtejbir694/gomark.git
cd gomark
go build -o gomark ./cmd/md2html
sudo mv gomark /usr/local/bin/
```
### Global Install

```bash
go install github.com/gurtejbir694/gomark/cmd/md2html@latest
```
> Now `gomark` is available in your terminal.

## Usage

```bash
gomark [options] <input.md> <output.html>
```
## Options

| Flag | Description | Default |
| ------ | ------ | ------ |
| `-title` | HTML `<title>` | Markdown Document |
| `-css` | Path to custom CSS | Uses embedded GitHub-style CSS |

