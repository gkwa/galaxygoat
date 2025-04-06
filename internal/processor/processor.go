package processor

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/gkwa/galaxygoat/internal/parser"
	"github.com/gkwa/galaxygoat/internal/transformer"
)

// HTMLProcessor handles the HTML processing workflow
type HTMLProcessor struct {
	parser      parser.HTMLParser
	transformer transformer.HTMLTransformer
	elements    map[string]bool
}

// NewHTMLProcessor creates a new HTMLProcessor with the specified configuration
func NewHTMLProcessor(
	elementsToRemove string,
	parser parser.HTMLParser,
	transformer transformer.HTMLTransformer,
) (*HTMLProcessor, error) {
	// Parse comma-separated elements to remove
	elementsSet := make(map[string]bool)
	for _, el := range strings.Split(elementsToRemove, ",") {
		trimmed := strings.TrimSpace(el)
		if trimmed != "" {
			elementsSet[trimmed] = true
		}
	}

	if len(elementsSet) == 0 {
		return nil, fmt.Errorf("no elements specified to remove")
	}

	return &HTMLProcessor{
		parser:      parser,
		transformer: transformer,
		elements:    elementsSet,
	}, nil
}

// ProcessHTML reads HTML from the reader, processes it, and writes the result to the writer
func (p *HTMLProcessor) ProcessHTML(reader io.Reader, writer io.Writer) error {
	// Read all content
	content, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("error reading input: %w", err)
	}

	// Parse HTML
	doc, err := p.parser.Parse(bytes.NewReader(content))
	if err != nil {
		return fmt.Errorf("error parsing HTML: %w", err)
	}

	// Transform HTML
	err = p.transformer.Transform(doc, p.elements)
	if err != nil {
		return fmt.Errorf("error transforming HTML: %w", err)
	}

	// Render the modified HTML
	err = p.parser.Render(writer, doc)
	if err != nil {
		return fmt.Errorf("error rendering HTML: %w", err)
	}

	return nil
}
