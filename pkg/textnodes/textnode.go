package textnodes

import "fmt"

type TextNode struct {
	Text     string
	TextType string
	Url      string
}

func NewTextNode(text, textType, url string) *TextNode {
	return &TextNode{text, textType, url}
}

func (node *TextNode) Equals(other *TextNode) bool {
	if node == nil || other == nil {
		return node == other
	}

	return node.Text == other.Text &&
		node.TextType == other.TextType &&
		node.Url == other.Url
}

func (node *TextNode) String() string {
	if node == nil {
		return "*TextNode<nil>"
	}

	if node.Url == "" {
		return fmt.Sprintf("TextNode(%s, %s)", node.Text, node.TextType)
	}

	return fmt.Sprintf("TextNode(%s, %s, %s)", node.Text, node.TextType, node.Url)
}
