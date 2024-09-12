package textnodes

import "strings"

func SplitNodesByDelimiter(oldNodes []*TextNode, delimiter, textType string) ([]*TextNode, error) {
	if len(oldNodes) == 0 || delimiter == "" {
		return oldNodes, nil
	}

	result := []*TextNode{}

	for _, oldNode := range oldNodes {
		if oldNode == nil {
			continue
		}

		if oldNode.TextType != TEXT_NODE_TYPE_TEXT || oldNode.Text == "" {
			result = append(result, oldNode)
			continue
		}

		parts := strings.Split(oldNode.Text, delimiter)

		if len(parts)%2 == 0 {
			return []*TextNode{}, ErrInvalidMarkdownSyntax
		}

		if len(parts) == 1 {
			result = append(result, oldNode)
			continue
		}

		for i, part := range parts {
			if i%2 == 0 {
				result = append(result, NewTextNode(part, oldNode.TextType, ""))
			} else {
				result = append(result, NewTextNode(part, textType, ""))
			}
		}
	}

	return result, nil
}
