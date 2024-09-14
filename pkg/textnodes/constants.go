package textnodes

import "errors"

type TextNodeType string

const (
	TEXT_NODE_TYPE_TEXT   TextNodeType = "text"
	TEXT_NODE_TYPE_BOLD   TextNodeType = "bold"
	TEXT_NODE_TYPE_ITALIC TextNodeType = "italic"
	TEXT_NODE_TYPE_CODE   TextNodeType = "code"
	TEXT_NODE_TYPE_LINK   TextNodeType = "link"
	TEXT_NODE_TYPE_IMAGE  TextNodeType = "image"
)

var (
	ErrUnsupportedTextNodeType = errors.New("unsupported text node type")
	ErrInvalidMarkdownSyntax   = errors.New("invalid Markdown syntax")
)
