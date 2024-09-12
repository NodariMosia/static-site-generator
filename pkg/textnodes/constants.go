package textnodes

import "errors"

const (
	TEXT_NODE_TYPE_TEXT   = "text"
	TEXT_NODE_TYPE_BOLD   = "bold"
	TEXT_NODE_TYPE_ITALIC = "italic"
	TEXT_NODE_TYPE_CODE   = "code"
	TEXT_NODE_TYPE_LINK   = "link"
	TEXT_NODE_TYPE_IMAGE  = "image"
)

var (
	ErrUnsupportedTextNodeType = errors.New("unsupported text node type")
	ErrInvalidMarkdownSyntax   = errors.New("invalid Markdown syntax")
)
