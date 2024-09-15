package adapters

import (
	"fmt"
	"strings"

	hn "static-site-generator/pkg/htmlnodes"
	md "static-site-generator/pkg/markdown"
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

func MarkdownToHTMLNode(markdown string) (*hn.ParentNode, error) {
	children := []hn.HTMLStringer{}
	blocks := md.MarkdownToBlocks(markdown)

	for _, block := range blocks {
		var child hn.HTMLStringer
		var err error = nil

		blockType := md.MarkdownBlockToBlockType(block)
		switch blockType {
		case md.MARKDOWN_BLOCK_TYPE_PARAGRAPH:
			child, err = paragraphBlockToHTMLNode(block)
		case md.MARKDOWN_BLOCK_TYPE_HEADING:
			child, err = headingBlockToHTMLNode(block)
		case md.MARKDOWN_BLOCK_TYPE_CODE:
			child, err = codeBlockToHTMLNode(block)
		case md.MARKDOWN_BLOCK_TYPE_QUOTE:
			child, err = quoteBlockToHTMLNode(block)
		case md.MARKDOWN_BLOCK_TYPE_UNORDERED_LIST:
			child, err = unorderedListBlockToHTMLNode(block)
		case md.MARKDOWN_BLOCK_TYPE_ORDERED_LIST:
			child, err = orderedListBlockToHTMLNode(block)
		}

		if err != nil {
			return nil, err
		}

		children = append(children, child)
	}

	return hn.NewParentNode("div", children, map[string]string{})
}

func paragraphBlockToHTMLNode(block string) (hn.HTMLStringer, error) {
	lines := strings.Split(block, "\n")
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}

	block = strings.Join(lines, " ")

	children, err := inlineTextToHTMLNodes(block)
	if err != nil {
		return nil, err
	}

	return hn.NewParentNode("p", children, map[string]string{})
}

func headingBlockToHTMLNode(block string) (hn.HTMLStringer, error) {
	headingLevel := 0
	for i := 0; i < len(block) && block[i] == '#'; i++ {
		headingLevel++
	}

	tag := fmt.Sprintf("h%v", headingLevel)
	block = strings.TrimSpace(block[headingLevel+1:])

	children, err := inlineTextToHTMLNodes(block)
	if err != nil {
		return nil, err
	}

	return hn.NewParentNode(tag, children, map[string]string{})
}

func codeBlockToHTMLNode(block string) (hn.HTMLStringer, error) {
	block = strings.TrimPrefix(block, "```")
	block = strings.TrimPrefix(block, "\n")
	block = strings.TrimSuffix(block, "```")

	children, err := inlineTextToHTMLNodes(block)
	if err != nil {
		return nil, err
	}

	codeNode, err := hn.NewParentNode("code", children, map[string]string{})
	if err != nil {
		return nil, err
	}

	return hn.NewParentNode("pre", []hn.HTMLStringer{codeNode}, map[string]string{})
}

func quoteBlockToHTMLNode(block string) (hn.HTMLStringer, error) {
	lines := strings.Split(block, "\n")
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i][1:])
	}

	block = strings.Join(lines, " ")

	children, err := inlineTextToHTMLNodes(block)
	if err != nil {
		return nil, err
	}

	return hn.NewParentNode("blockquote", children, map[string]string{})
}

func unorderedListBlockToHTMLNode(block string) (hn.HTMLStringer, error) {
	lines := strings.Split(block, "\n")
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i][2:])
	}

	listItems, err := linesToListItemHTMLNodes(lines)
	if err != nil {
		return nil, err
	}

	return hn.NewParentNode("ul", listItems, map[string]string{})
}

func orderedListBlockToHTMLNode(block string) (hn.HTMLStringer, error) {
	lines := strings.Split(block, "\n")
	for i, line := range lines {
		_, lines[i], _ = strings.Cut(line, ".")
		lines[i] = strings.TrimSpace(lines[i][1:])
	}

	listItems, err := linesToListItemHTMLNodes(lines)
	if err != nil {
		return nil, err
	}

	return hn.NewParentNode("ol", listItems, map[string]string{})
}

func linesToListItemHTMLNodes(lines []string) ([]hn.HTMLStringer, error) {
	listItems := make([]hn.HTMLStringer, len(lines))

	for i, line := range lines {
		children, err := inlineTextToHTMLNodes(line)
		if err != nil {
			return nil, err
		}

		listItems[i], err = hn.NewParentNode("li", children, map[string]string{})
		if err != nil {
			return nil, err
		}
	}

	return listItems, nil
}

func inlineTextToHTMLNodes(text string) ([]hn.HTMLStringer, error) {
	textNodes, err := TextToTextNodes(text)
	if err != nil {
		return nil, err
	}

	htmlNodes := make([]hn.HTMLStringer, len(textNodes))

	for i, textNode := range textNodes {
		htmlNode, err := TextNodeToHTMLNode(textNode)
		if err != nil {
			return nil, err
		}

		htmlNodes[i] = htmlNode
	}

	return htmlNodes, nil
}
