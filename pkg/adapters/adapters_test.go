package adapters

import (
	"errors"
	"testing"

	"static-site-generator/pkg/textnodes"
)

func TestTextNodeToHTMLNode(t *testing.T) {
	tn1 := textnodes.NewTextNode("This should break", "someOtherType", "")
	tn2 := textnodes.NewTextNode("Normal text", textnodes.TEXT_NODE_TYPE_TEXT, "")
	tn3 := textnodes.NewTextNode("Bold text", textnodes.TEXT_NODE_TYPE_BOLD, "")
	tn4 := textnodes.NewTextNode("Italic text", textnodes.TEXT_NODE_TYPE_ITALIC, "")
	tn5 := textnodes.NewTextNode("Code block text", textnodes.TEXT_NODE_TYPE_CODE, "")
	tn6 := textnodes.NewTextNode("Anchor text", textnodes.TEXT_NODE_TYPE_LINK, "https://google.com")
	tn7 := textnodes.NewTextNode("Image alt text", textnodes.TEXT_NODE_TYPE_IMAGE, "https://avatars.githubusercontent.com/u/66739334?v=4")

	tests := []struct {
		name     string
		textNode *textnodes.TextNode
		wantHTML string
		wantErr  error
	}{
		{
			name:     "shouldReturn",
			textNode: tn1,
			wantHTML: "",
			wantErr:  textnodes.ErrUnsupportedTextNodeType,
		},
		{
			name:     "shouldReturnRawText",
			textNode: tn2,
			wantHTML: "Normal text",
			wantErr:  nil,
		},
		{
			name:     "shouldReturnBoldTag",
			textNode: tn3,
			wantHTML: "<b>Bold text</b>",
			wantErr:  nil,
		},
		{
			name:     "shouldReturnItalicTag",
			textNode: tn4,
			wantHTML: "<i>Italic text</i>",
			wantErr:  nil,
		},
		{
			name:     "shouldReturnCodeBlockTag",
			textNode: tn5,
			wantHTML: "<code>Code block text</code>",
			wantErr:  nil,
		},
		{
			name:     "shouldReturnAnchorTag",
			textNode: tn6,
			wantHTML: "<a href=\"https://google.com\">Anchor text</a>",
			wantErr:  nil,
		},
		{
			name:     "shouldReturnImageTag",
			textNode: tn7,
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
