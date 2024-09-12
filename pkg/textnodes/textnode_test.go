package textnodes

import "testing"

func TestTextNode_Equals(t *testing.T) {
	tests := []struct {
		name   string
		first  *TextNode
		second *TextNode
		want   bool
	}{
		{
			name:   "shouldEqualWhenAllPropertiesAreEqual",
			first:  NewTextNode("This is a text node", TEXT_NODE_TYPE_BOLD, "/"),
			second: NewTextNode("This is a text node", TEXT_NODE_TYPE_BOLD, "/"),
			want:   true,
		},
		{
			name:   "shouldNotEqualWhenSomePropertiesAreDifferent",
			first:  NewTextNode("This is a text node", TEXT_NODE_TYPE_BOLD, "/"),
			second: NewTextNode("This is a different text node", TEXT_NODE_TYPE_BOLD, "/"),
			want:   false,
		},
		{
			name:   "shouldNotEqualWhenSomePropertiesAreDifferent",
			first:  NewTextNode("This is a text node", TEXT_NODE_TYPE_BOLD, ""),
			second: NewTextNode("This is a text node", TEXT_NODE_TYPE_ITALIC, ""),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.first.Equals(tt.second)
			if got != tt.want {
				t.Errorf("TextNode.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTextNode_String(t *testing.T) {
	tests := []struct {
		name string
		node *TextNode
		want string
	}{
		{
			name: "shouldFormatTextNodeWithoutUrl",
			node: NewTextNode("This is a text node", TEXT_NODE_TYPE_BOLD, ""),
			want: "TextNode(This is a text node, bold)",
		},
		{
			name: "shouldFormatTextNodeWithUrl",
			node: NewTextNode("This is a text node", TEXT_NODE_TYPE_BOLD, "/"),
			want: "TextNode(This is a text node, bold, /)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.node.String()
			if got != tt.want {
				t.Errorf("TextNode.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
