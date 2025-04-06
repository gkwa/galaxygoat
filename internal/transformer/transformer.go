package transformer

import (
	"golang.org/x/net/html"
)

// HTMLTransformer defines the interface for transforming HTML
type HTMLTransformer interface {
	Transform(node *html.Node, options map[string]bool) error
}

// ElementRemover implements HTMLTransformer to remove specified elements
type ElementRemover struct{}

// Transform removes elements specified in the options map
func (t *ElementRemover) Transform(node *html.Node, elementsToRemove map[string]bool) error {
	removeElements(node, elementsToRemove)
	return nil
}

// removeElements removes all elements of the specified types from the HTML node tree
func removeElements(n *html.Node, elementsToRemove map[string]bool) {
	// Process children first (we need to traverse depth-first)
	var next *html.Node
	for c := n.FirstChild; c != nil; c = next {
		// Save the next node before potentially removing the current one
		next = c.NextSibling
		removeElements(c, elementsToRemove)
	}

	// Check if this node should be removed
	if n.Type == html.ElementNode && elementsToRemove[n.Data] {
		// If this node should be removed, unlink it from the tree
		if n.Parent != nil {
			n.Parent.RemoveChild(n)
		}
	}
}
