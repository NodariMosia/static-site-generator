package textnodes

import (
	"fmt"
	"strings"

	"static-site-generator/pkg/utils"
)

func SplitNodesByDelimiter(nodes []*TextNode, delimiter, textType string) ([]*TextNode, error) {
	if len(nodes) == 0 || delimiter == "" {
		return nodes, nil
	}

	result := []*TextNode{}

	for _, node := range nodes {
		if node == nil {
			continue
		}

		if node.TextType != TEXT_NODE_TYPE_TEXT || node.Text == "" {
			result = append(result, node)
			continue
		}

		parts := strings.Split(node.Text, delimiter)

		if len(parts)%2 == 0 {
			return []*TextNode{}, ErrInvalidMarkdownSyntax
		}

		if len(parts) == 1 {
			result = append(result, node)
			continue
		}

		for i, part := range parts {
			if i%2 == 0 {
				result = append(result, NewTextNode(part, TEXT_NODE_TYPE_TEXT, ""))
			} else {
				result = append(result, NewTextNode(part, textType, ""))
			}
		}
	}

	return result, nil
}

func SplitNodesByImages(nodes []*TextNode) []*TextNode {
	return splitNodesByTextUrlPairs(
		nodes,
		utils.ExtractMarkdownImages,
		"![%s](%s)",
		TEXT_NODE_TYPE_IMAGE,
	)
}

func SplitNodesByLinks(nodes []*TextNode) []*TextNode {
	return splitNodesByTextUrlPairs(
		nodes,
		utils.ExtractMarkdownLinks,
		"[%s](%s)",
		TEXT_NODE_TYPE_LINK,
	)
}

func splitNodesByTextUrlPairs(
	nodes []*TextNode,
	markdownTextUrlPairExtractor func(text string) []utils.MarkdownTextUrlPair,
	markdownTextUrlPairFormat string,
	textType string,
) []*TextNode {
	if len(nodes) == 0 {
		return nodes
	}

	result := []*TextNode{}

	for _, node := range nodes {
		if node == nil {
			continue
		}

		if node.TextType != TEXT_NODE_TYPE_TEXT {
			result = append(result, node)
			continue
		}

		if node.Text == "" {
			continue
		}

		remainingText := node.Text
		textUrlPairs := markdownTextUrlPairExtractor(node.Text)

		for _, pair := range textUrlPairs {
			before, after, found := strings.Cut(
				remainingText,
				fmt.Sprintf(markdownTextUrlPairFormat, pair.Text, pair.Url),
			)

			if !found {
				continue
			}

			if before != "" {
				result = append(result, NewTextNode(before, TEXT_NODE_TYPE_TEXT, ""))
			}

			result = append(result, NewTextNode(pair.Text, textType, pair.Url))

			remainingText = after
		}

		if remainingText != "" {
			result = append(result, NewTextNode(remainingText, TEXT_NODE_TYPE_TEXT, ""))
		}
	}

	return result
}
