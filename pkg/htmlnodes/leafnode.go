package htmlnodes

import "fmt"

type LeafNode struct {
	HTMLNode
}

func NewLeafNode(tag, value string, props map[string]string) (*LeafNode, error) {
	return &LeafNode{HTMLNode{tag, value, nil, props}}, nil
}

func (node *LeafNode) ToHTML() (string, error) {
	if node == nil {
		return "", ErrToHTMLCalledOnNilNodePointer
	}

	if node.Tag == "" {
		return node.Value, nil
	}

	return fmt.Sprintf("<%s%s>%s</%s>", node.Tag, node.PropsToHTML(), node.Value, node.Tag), nil
}
