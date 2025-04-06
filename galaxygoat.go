package galaxygoat

import (
	"bytes"
	"io"

	"github.com/gkwa/galaxygoat/internal/parser"
	"github.com/gkwa/galaxygoat/internal/processor"
	"github.com/gkwa/galaxygoat/internal/transformer"
)

// RemoveElements removes specified HTML elements from HTML content
func RemoveElements(htmlContent []byte, elementsToRemove string) ([]byte, error) {
	// Create an HTML processor with the specified elements to remove
	proc, err := processor.NewHTMLProcessor(
		elementsToRemove,
		&parser.NetHTMLParser{},
		&transformer.ElementRemover{},
	)
	if err != nil {
		return nil, err
	}

	// Create in-memory buffers
	reader := bytes.NewReader(htmlContent)
	var buf bytes.Buffer

	// Process the HTML
	if err := proc.ProcessHTML(reader, &buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// RemoveElementsFromReader processes HTML from a reader and removes specified elements
func RemoveElementsFromReader(reader io.Reader, elementsToRemove string, writer io.Writer) error {
	// Create an HTML processor with the specified elements to remove
	proc, err := processor.NewHTMLProcessor(
		elementsToRemove,
		&parser.NetHTMLParser{},
		&transformer.ElementRemover{},
	)
	if err != nil {
		return err
	}

	// Process the HTML
	return proc.ProcessHTML(reader, writer)
}
