package adapters

import (
	hn "static-site-generator/pkg/htmlnodes"
	tn "static-site-generator/pkg/textnodes"
)

func TextToTextNodes(text string) ([]*tn.TextNode, error) {
	textNodes := []*tn.TextNode{tn.NewTextNode(text, tn.TEXT_NODE_TYPE_TEXT, "")}

	textNodes = tn.SplitNodesByImages(textNodes)
	textNodes = tn.SplitNodesByLinks(textNodes)

	textNodes, err := tn.SplitNodesByDelimiter(textNodes, "**", tn.TEXT_NODE_TYPE_BOLD)
	if err != nil {
		return nil, err
	}

	textNodes, err = tn.SplitNodesByDelimiter(textNodes, "*", tn.TEXT_NODE_TYPE_ITALIC)
	if err != nil {
		return nil, err
	}

	textNodes, err = tn.SplitNodesByDelimiter(textNodes, "`", tn.TEXT_NODE_TYPE_CODE)
	if err != nil {
		return nil, err
	}

	return textNodes, nil
}

func TextNodeToHTMLNode(textNode *tn.TextNode) (*hn.LeafNode, error) {
	switch textNode.TextType {
	case tn.TEXT_NODE_TYPE_TEXT:
		return hn.NewLeafNode("", textNode.Text, nil)
	case tn.TEXT_NODE_TYPE_BOLD:
		return hn.NewLeafNode("b", textNode.Text, nil)
	case tn.TEXT_NODE_TYPE_ITALIC:
		return hn.NewLeafNode("i", textNode.Text, nil)
	case tn.TEXT_NODE_TYPE_CODE:
		return hn.NewLeafNode("code", textNode.Text, nil)
	case tn.TEXT_NODE_TYPE_LINK:
		return hn.NewLeafNode(
			"a", textNode.Text,
			map[string]string{"href": textNode.Url},
		)
	case tn.TEXT_NODE_TYPE_IMAGE:
		return hn.NewLeafNode(
			"img", "",
			map[string]string{"alt": textNode.Text, "src": textNode.Url},
		)
	default:
		return nil, tn.ErrUnsupportedTextNodeType
	}
}
