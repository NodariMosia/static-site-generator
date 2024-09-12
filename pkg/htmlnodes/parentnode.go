package htmlnodes

import (
	"fmt"
	"strings"
)

type ParentNode struct {
	HTMLNode
}

func NewParentNode(tag string, children []HTMLStringer, props map[string]string) (*ParentNode, error) {
	if tag == "" {
		return nil, ErrEmptyTag
	}

	if len(children) == 0 {
		return nil, ErrEmptyChildren
	}

	return &ParentNode{HTMLNode{tag, "", children, props}}, nil
}

func (node *ParentNode) ToHTML() (string, error) {
	if node == nil {
		return "", ErrToHTMLCalledOnNilNodePointer
	}

	if node.Tag == "" {
		return "", ErrEmptyTag
	}

	if len(node.Children) == 0 {
		return "", ErrEmptyChildren
	}

	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("<%s%s>", node.Tag, node.PropsToHTML()))

	for _, child := range node.Children {
		childHTML, err := child.ToHTML()
		if err != nil {
			return "", err
		}

		sb.WriteString(childHTML)
	}

	sb.WriteString(fmt.Sprintf("</%s>", node.Tag))

	return sb.String(), nil
}
