package htmlnodes

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

type HTMLStringer interface {
	ToHTML() (string, error)
	PropsToHTML() string
	String() string
}

type HTMLNode struct {
	Tag      string
	Value    string
	Children []HTMLStringer
	Props    map[string]string
}

func NewHTMLNode(tag, value string, children []HTMLStringer, props map[string]string) *HTMLNode {
	return &HTMLNode{tag, value, children, props}
}

func (node *HTMLNode) ToHTML() (string, error) {
	return "", ErrUnimplementedMethodCall
}

func (node *HTMLNode) childrenToString() string {
	if node == nil || len(node.Children) == 0 {
		return "[]"
	}

	var sb strings.Builder

	sb.WriteByte('[')
	for i, child := range node.Children {
		if i != 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(child.String())
	}
	sb.WriteByte(']')

	return sb.String()
}

func (node *HTMLNode) PropsToHTML() string {
	if node == nil || len(node.Props) == 0 {
		return ""
	}

	var sb strings.Builder

	for _, k := range slices.Sorted(maps.Keys(node.Props)) {
		sb.WriteString(fmt.Sprintf(" %s=\"%s\"", k, node.Props[k]))
	}

	return sb.String()
}

func (node *HTMLNode) String() string {
	if node == nil {
		return "*HTMLNode<nil>"
	}

	return fmt.Sprintf(
		"HTMLNode(Tag: `%s`, Value: `%s`, Children: %s, Props: `%s`)",
		node.Tag, node.Value, node.childrenToString(), node.PropsToHTML(),
	)
}
