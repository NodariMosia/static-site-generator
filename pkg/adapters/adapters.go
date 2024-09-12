package adapters

import (
	"static-site-generator/pkg/htmlnodes"
	"static-site-generator/pkg/textnodes"
)

func TextNodeToHTMLNode(textNode *textnodes.TextNode) (*htmlnodes.LeafNode, error) {
	switch textNode.TextType {
	case textnodes.TEXT_NODE_TYPE_TEXT:
		return htmlnodes.NewLeafNode("", textNode.Text, nil)
	case textnodes.TEXT_NODE_TYPE_BOLD:
		return htmlnodes.NewLeafNode("b", textNode.Text, nil)
	case textnodes.TEXT_NODE_TYPE_ITALIC:
		return htmlnodes.NewLeafNode("i", textNode.Text, nil)
	case textnodes.TEXT_NODE_TYPE_CODE:
		return htmlnodes.NewLeafNode("code", textNode.Text, nil)
	case textnodes.TEXT_NODE_TYPE_LINK:
		return htmlnodes.NewLeafNode(
			"a", textNode.Text,
			map[string]string{"href": textNode.Url},
		)
	case textnodes.TEXT_NODE_TYPE_IMAGE:
		return htmlnodes.NewLeafNode(
			"img", "",
			map[string]string{"alt": textNode.Text, "src": textNode.Url},
		)
	default:
		return nil, textnodes.ErrUnsupportedTextNodeType
	}
}
