package adapters

import (
	"errors"
	"reflect"
	"testing"

	tn "static-site-generator/pkg/textnodes"
)

func TestTextToTextNodes(t *testing.T) {
	tests := []struct {
		name      string
		text      string
		wantNodes []*tn.TextNode
		wantErr   error
	}{
		{
			name: "shouldParseMultipleTypesOfTextNodes",
			text: "This is **text** with an *italic* word and a `code block` and an ![obi wan image](https://i.imgur.com/fJRm4Vk.jpeg) and a [link](https://google.com)",
			wantNodes: []*tn.TextNode{
				tn.NewTextNode("This is ", tn.TEXT_NODE_TYPE_TEXT, ""),
				tn.NewTextNode("text", tn.TEXT_NODE_TYPE_BOLD, ""),
				tn.NewTextNode(" with an ", tn.TEXT_NODE_TYPE_TEXT, ""),
				tn.NewTextNode("italic", tn.TEXT_NODE_TYPE_ITALIC, ""),
				tn.NewTextNode(" word and a ", tn.TEXT_NODE_TYPE_TEXT, ""),
				tn.NewTextNode("code block", tn.TEXT_NODE_TYPE_CODE, ""),
				tn.NewTextNode(" and an ", tn.TEXT_NODE_TYPE_TEXT, ""),
				tn.NewTextNode("obi wan image", tn.TEXT_NODE_TYPE_IMAGE, "https://i.imgur.com/fJRm4Vk.jpeg"),
				tn.NewTextNode(" and a ", tn.TEXT_NODE_TYPE_TEXT, ""),
				tn.NewTextNode("link", tn.TEXT_NODE_TYPE_LINK, "https://google.com"),
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNodes, gotErr := TextToTextNodes(tt.text)
			if !reflect.DeepEqual(gotNodes, tt.wantNodes) || !errors.Is(gotErr, tt.wantErr) {
				t.Errorf(
					"TextToTextNodes() = (%v, %v), want (%v, %v)",
					gotNodes, gotErr, tt.wantNodes, tt.wantErr,
				)
			}
		})
	}
}

func TestTextNodeToHTMLNode(t *testing.T) {
	node1 := tn.NewTextNode("This should break", "someOtherType", "")
	node2 := tn.NewTextNode("Normal text", tn.TEXT_NODE_TYPE_TEXT, "")
	node3 := tn.NewTextNode("Bold text", tn.TEXT_NODE_TYPE_BOLD, "")
	node4 := tn.NewTextNode("Italic text", tn.TEXT_NODE_TYPE_ITALIC, "")
	node5 := tn.NewTextNode("Code block text", tn.TEXT_NODE_TYPE_CODE, "")
	node6 := tn.NewTextNode("Anchor text", tn.TEXT_NODE_TYPE_LINK, "https://google.com")
	node7 := tn.NewTextNode("Image alt text", tn.TEXT_NODE_TYPE_IMAGE, "https://avatars.githubusercontent.com/u/66739334?v=4")

	tests := []struct {
		name     string
		textNode *tn.TextNode
		wantHTML string
		wantErr  error
	}{
		{
			name:     "shouldReturn",
			textNode: node1,
			wantHTML: "",
			wantErr:  tn.ErrUnsupportedTextNodeType,
		},
		{
			name:     "shouldReturnRawText",
			textNode: node2,
			wantHTML: "Normal text",
			wantErr:  nil,
		},
		{
			name:     "shouldReturnBoldTag",
			textNode: node3,
			wantHTML: "<b>Bold text</b>",
			wantErr:  nil,
		},
		{
			name:     "shouldReturnItalicTag",
			textNode: node4,
			wantHTML: "<i>Italic text</i>",
			wantErr:  nil,
		},
		{
			name:     "shouldReturnCodeBlockTag",
			textNode: node5,
			wantHTML: "<code>Code block text</code>",
			wantErr:  nil,
		},
		{
			name:     "shouldReturnAnchorTag",
			textNode: node6,
			wantHTML: "<a href=\"https://google.com\">Anchor text</a>",
			wantErr:  nil,
		},
		{
			name:     "shouldReturnImageTag",
			textNode: node7,
			wantHTML: "<img alt=\"Image alt text\" src=\"https://avatars.githubusercontent.com/u/66739334?v=4\"></img>",
			wantErr:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			leafNode, gotErr := TextNodeToHTMLNode(tt.textNode)
			if !errors.Is(gotErr, tt.wantErr) {
				t.Errorf(
					"TextNodeToHTMLNode(%s) = (_, %v), want (_, %v)",
					tt.textNode.String(), gotErr, tt.wantErr,
				)
			}

			gotHTML, _ := leafNode.ToHTML()
			if gotHTML != tt.wantHTML {
				t.Errorf(
					"TextNodeToHTMLNode(%s) = (%v, _), want (%v, _)",
					tt.textNode.String(), gotHTML, tt.wantHTML,
				)
			}
		})
	}
}
