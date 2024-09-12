package htmlnodes

import (
	"errors"
	"testing"
)

func TestNewLeafNode(t *testing.T) {
	_, err1 := NewLeafNode("", "", nil)
	_, err2 := NewLeafNode("", "This is a raw text.", nil)
	_, err3 := NewLeafNode("a", "About", map[string]string{"href": "/about"})

	tests := []struct {
		name    string
		err     error
		wantErr error
	}{
		{
			name:    "shouldConstructWithDefaultValues",
			err:     err1,
			wantErr: nil,
		},
		{
			name:    "shouldConstructWithJustValue",
			err:     err2,
			wantErr: nil,
		},
		{
			name:    "shouldConstructWithAllParams",
			err:     err3,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !errors.Is(tt.err, tt.wantErr) {
				t.Errorf(
					"NewLeafNode() = (_, %v), want (_, %v)",
					tt.err, tt.wantErr,
				)
			}
		})
	}
}

func TestLeafNode_ToHTML(t *testing.T) {
	ln1, _ := NewLeafNode("", "", nil)
	ln2, _ := NewLeafNode("", "This is a raw text.", nil)
	ln3, _ := NewLeafNode("p", "This is a paragraph of text.", nil)
	ln4, _ := NewLeafNode("p", "This is a paragraph of text.", map[string]string{})
	ln5, _ := NewLeafNode("a", "About", map[string]string{"href": "/about"})
	ln6, _ := NewLeafNode("a", "About", map[string]string{"href": "/about", "target": "_blank"})

	tests := []struct {
		name    string
		node    *LeafNode
		wantStr string
		wantErr error
	}{
		{
			name:    "shouldReturnErrToHTMLCalledOnNilNodePointer",
			node:    nil,
			wantStr: "",
			wantErr: ErrToHTMLCalledOnNilNodePointer,
		},
		{
			name:    "shouldReturnEmptyString",
			node:    ln1,
			wantStr: "",
			wantErr: nil,
		},
		{
			name:    "shouldReturnValueAsRawText",
			node:    ln2,
			wantStr: "This is a raw text.",
			wantErr: nil,
		},
		{
			name:    "shouldFormatWithNilProps",
			node:    ln3,
			wantStr: "<p>This is a paragraph of text.</p>",
			wantErr: nil,
		},
		{
			name:    "shouldFormatWithEmptyProps",
			node:    ln4,
			wantStr: "<p>This is a paragraph of text.</p>",
			wantErr: nil,
		},
		{
			name:    "shouldFormatWithOneProp",
			node:    ln5,
			wantStr: "<a href=\"/about\">About</a>",
			wantErr: nil,
		},
		{
			name:    "shouldFormatWithMultipleProps",
			node:    ln6,
			wantStr: "<a href=\"/about\" target=\"_blank\">About</a>",
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStr, gotErr := tt.node.ToHTML()
			if gotStr != tt.wantStr || !errors.Is(gotErr, tt.wantErr) {
				t.Errorf(
					"LeafNode.ToHTML() = (%v, %v), want (%v, %v)",
					gotStr, gotErr, tt.wantStr, tt.wantErr,
				)
			}
		})
	}
}
