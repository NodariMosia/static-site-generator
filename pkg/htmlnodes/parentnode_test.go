package htmlnodes

import (
	"errors"
	"testing"
)

func TestNewParentNode(t *testing.T) {
	ln1, _ := NewLeafNode("b", "Bold text", nil)
	ln2, _ := NewLeafNode("", "Normal text", nil)
	ln3, _ := NewLeafNode("i", "italic text", map[string]string{"style": "color:red"})
	ln4, _ := NewLeafNode("", "Normal text", nil)

	_, err1 := NewParentNode("", []HTMLStringer{ln1, ln2, ln3}, nil)
	_, err2 := NewParentNode("p", []HTMLStringer{}, nil)
	_, err3 := NewParentNode("p", []HTMLStringer{ln1, ln2, ln3}, nil)
	_, err4 := NewParentNode(
		"p",
		[]HTMLStringer{ln1, ln2, ln3, ln4},
		map[string]string{"style": "background:gray"},
	)

	tests := []struct {
		name    string
		err     error
		wantErr error
	}{
		{
			name:    "shouldReturnErrEmptyTag",
			err:     err1,
			wantErr: ErrEmptyTag,
		},
		{
			name:    "shouldReturnErrEmptyChildren",
			err:     err2,
			wantErr: ErrEmptyChildren,
		},
		{
			name:    "shouldConstructWithTagAndChildren",
			err:     err3,
			wantErr: nil,
		},
		{
			name:    "shouldConstructWithAllParams",
			err:     err4,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !errors.Is(tt.err, tt.wantErr) {
				t.Errorf(
					"NewParentNode() = (_, %v), want (_, %v)",
					tt.err, tt.wantErr,
				)
			}
		})
	}
}

func TestParentNode_ToHTML(t *testing.T) {
	ln1, _ := NewLeafNode("b", "Bold text", nil)
	ln2, _ := NewLeafNode("", "Normal text", nil)
	ln3, _ := NewLeafNode("i", "italic text", map[string]string{"style": "color:red"})
	ln4, _ := NewLeafNode("", "Normal text", nil)

	pn1, _ := NewParentNode("p", []HTMLStringer{ln1}, nil)
	pn2, _ := NewParentNode("p", []HTMLStringer{ln1}, nil)
	pn3, _ := NewParentNode("p", []HTMLStringer{ln1}, nil)
	pn1.Tag = ""                    // should cause ErrEmptyTag
	pn2.Children = nil              // should cause ErrEmptyChildren
	pn3.Children = []HTMLStringer{} // should cause ErrEmptyChildren

	pn4, _ := NewParentNode("p", []HTMLStringer{ln1}, nil)
	pn5, _ := NewParentNode(
		"p",
		[]HTMLStringer{ln1, ln2, ln3, ln4},
		map[string]string{"style": "background:blue"},
	)
	pn6, _ := NewParentNode("div", []HTMLStringer{pn4, ln3}, nil)
	pn7, _ := NewParentNode(
		"div",
		[]HTMLStringer{pn5, ln3, pn6, ln4},
		map[string]string{"style": "background:gray"},
	)

	tests := []struct {
		name    string
		node    *ParentNode
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
			name:    "shouldReturnErrEmptyTag",
			node:    pn1,
			wantStr: "",
			wantErr: ErrEmptyTag,
		},
		{
			name:    "shouldReturnErrEmptyChildren",
			node:    pn2,
			wantStr: "",
			wantErr: ErrEmptyChildren,
		},
		{
			name:    "shouldReturnErrEmptyChildren",
			node:    pn3,
			wantStr: "",
			wantErr: ErrEmptyChildren,
		},
		{
			name:    "shouldFormatSingleChild",
			node:    pn4,
			wantStr: "<p><b>Bold text</b></p>",
			wantErr: nil,
		},
		{
			name:    "shouldFormatMultipleChildren",
			node:    pn5,
			wantStr: "<p style=\"background:blue\"><b>Bold text</b>Normal text<i style=\"color:red\">italic text</i>Normal text</p>",
			wantErr: nil,
		},
		{
			name:    "shouldFormatNestedParentNode",
			node:    pn6,
			wantStr: "<div><p><b>Bold text</b></p><i style=\"color:red\">italic text</i></div>",
			wantErr: nil,
		},
		{
			name:    "shouldFormatDeeplyNestedParentNode",
			node:    pn7,
			wantStr: "<div style=\"background:gray\"><p style=\"background:blue\"><b>Bold text</b>Normal text<i style=\"color:red\">italic text</i>Normal text</p><i style=\"color:red\">italic text</i><div><p><b>Bold text</b></p><i style=\"color:red\">italic text</i></div>Normal text</div>",
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStr, gotErr := tt.node.ToHTML()
			if gotStr != tt.wantStr || !errors.Is(gotErr, tt.wantErr) {
				t.Errorf(
					"ParentNode.ToHTML() = (%v, %v), want (%v, %v)",
					gotStr, gotErr, tt.wantStr, tt.wantErr,
				)
			}
		})
	}
}
