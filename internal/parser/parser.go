package parser

import (
	"io"

	"golang.org/x/net/html"
)

// HTMLParser defines the interface for parsing and rendering HTML
type HTMLParser interface {
	Parse(r io.Reader) (*html.Node, error)
	Render(w io.Writer, n *html.Node) error
}

// NetHTMLParser implements HTMLParser using golang.org/x/net/html
type NetHTMLParser struct{}

// Parse parses HTML content from a reader
func (p *NetHTMLParser) Parse(r io.Reader) (*html.Node, error) {
	return html.Parse(r)
}

// Render renders an HTML node to a writer
func (p *NetHTMLParser) Render(w io.Writer, n *html.Node) error {
	return html.Render(w, n)
}
