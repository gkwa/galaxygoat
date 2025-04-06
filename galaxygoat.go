package galaxygoat

import (
	"bytes"
	"io"
	"strings"

	"golang.org/x/net/html"
)

// RemoveElements removes specified HTML elements from HTML content
func RemoveElements(htmlContent []byte, elementsToRemove string) ([]byte, error) {
	// Parse elements to remove
	elementsSet := parseElementsToRemove(elementsToRemove)

	// Parse the HTML
	doc, err := html.Parse(bytes.NewReader(htmlContent))
	if err != nil {
		return nil, err
	}

	// Transform the HTML
	removeElementsFromNode(doc, elementsSet)

	// Render the transformed HTML
	var buf bytes.Buffer
	if err := html.Render(&buf, doc); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// RemoveElementsFromReader processes HTML from a reader and removes specified elements
func RemoveElementsFromReader(reader io.Reader, elementsToRemove string, writer io.Writer) error {
	// Read all content
	content, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	// Process the content
	result, err := RemoveElements(content, elementsToRemove)
	if err != nil {
		return err
	}

	// Write the result
	_, err = writer.Write(result)
	return err
}

// parseElementsToRemove converts a comma-separated string into a map for quick lookups
func parseElementsToRemove(elementsToRemove string) map[string]bool {
	elementsSet := make(map[string]bool)
	for _, el := range strings.Split(elementsToRemove, ",") {
		trimmed := strings.TrimSpace(el)
		if trimmed != "" {
			elementsSet[trimmed] = true
		}
	}
	return elementsSet
}

// removeElementsFromNode removes all elements of the specified types from the HTML node tree
func removeElementsFromNode(n *html.Node, elementsToRemove map[string]bool) {
	// Process children first (we need to traverse depth-first)
	var next *html.Node
	for c := n.FirstChild; c != nil; c = next {
		// Save the next node before potentially removing the current one
		next = c.NextSibling
		removeElementsFromNode(c, elementsToRemove)
	}

	// Check if this node should be removed
	if n.Type == html.ElementNode && elementsToRemove[n.Data] {
		// If this node should be removed, unlink it from the tree
		if n.Parent != nil {
			n.Parent.RemoveChild(n)
		}
	}
}
