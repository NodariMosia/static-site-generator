package htmlnodes

import "testing"

func TestHTMLNode_PropsToHTML(t *testing.T) {
	tests := []struct {
		name string
		node *HTMLNode
		want string
	}{
		{
			name: "shouldReturnEmptyStringForNilProps",
			node: NewHTMLNode("a", "Google", nil, nil),
			want: "",
		},
		{
			name: "shouldReturnEmptyStringForEmptyProps",
			node: NewHTMLNode("a", "Google", nil, map[string]string{}),
			want: "",
		},
		{
			name: "shouldFormatOneProp",
			node: NewHTMLNode(
				"a",
				"Google",
				nil,
				map[string]string{"href": "https://www.google.com"},
			),
			want: " href=\"https://www.google.com\"",
		},
		{
			name: "shouldFormatMultipleProps",
			node: NewHTMLNode(
				"a",
				"Google",
				nil,
				map[string]string{
					"href":   "https://www.google.com",
					"target": "_blank",
				},
			),
			want: " href=\"https://www.google.com\" target=\"_blank\"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.node.PropsToHTML()
			if got != tt.want {
				t.Errorf("HTMLNode.PropsToHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTMLNode_String(t *testing.T) {
	tests := []struct {
		name string
		node *HTMLNode
		want string
	}{
		{
			name: "shouldFormatAllParams",
			node: NewHTMLNode(
				"nav",
				"Links:",
				[]HTMLStringer{
					NewHTMLNode("a", "Home", nil, map[string]string{"href": "/"}),
					NewHTMLNode("a", "About", nil, map[string]string{"href": "/about"}),
				},
				map[string]string{"style": "background:gray"},
			),
			want: "HTMLNode(Tag: `nav`, Value: `Links:`, Children: [HTMLNode(Tag: `a`, Value: `Home`, Children: [], Props: ` href=\"/\"`), HTMLNode(Tag: `a`, Value: `About`, Children: [], Props: ` href=\"/about\"`)], Props: ` style=\"background:gray\"`)",
		},
		{
			name: "shouldFormatAllDefaultParams",
			node: &HTMLNode{},
			want: "HTMLNode(Tag: ``, Value: ``, Children: [], Props: ``)",
		},
		{
			name: "shouldFormatAllDefaultParamsFromConstructor",
			node: NewHTMLNode("", "", nil, nil),
			want: "HTMLNode(Tag: ``, Value: ``, Children: [], Props: ``)",
		},
		{
			name: "shouldFormatEmptyTag",
			node: NewHTMLNode("", "Hello world!", nil, nil),
			want: "HTMLNode(Tag: ``, Value: `Hello world!`, Children: [], Props: ``)",
		},
		{
			name: "shouldFormatEmptyValue",
			node: NewHTMLNode(
				"nav",
				"",
				[]HTMLStringer{
					NewHTMLNode("a", "About", nil, map[string]string{"href": "/about"}),
				},
				nil,
			),
			want: "HTMLNode(Tag: `nav`, Value: ``, Children: [HTMLNode(Tag: `a`, Value: `About`, Children: [], Props: ` href=\"/about\"`)], Props: ``)",
		},
		{
			name: "shouldFormatNilChildren",
			node: NewHTMLNode("a", "About", nil, map[string]string{"href": "/about"}),
			want: "HTMLNode(Tag: `a`, Value: `About`, Children: [], Props: ` href=\"/about\"`)",
		},
		{
			name: "shouldFormatEmptyChildren",
			node: NewHTMLNode("a", "About", []HTMLStringer{}, map[string]string{"href": "/about"}),
			want: "HTMLNode(Tag: `a`, Value: `About`, Children: [], Props: ` href=\"/about\"`)",
		},
		{
			name: "shouldFormatNilProps",
			node: NewHTMLNode("div", "Hello world!", nil, nil),
			want: "HTMLNode(Tag: `div`, Value: `Hello world!`, Children: [], Props: ``)",
		},
		{
			name: "shouldFormatEmptyProps",
			node: NewHTMLNode("div", "Hello world!", nil, map[string]string{}),
			want: "HTMLNode(Tag: `div`, Value: `Hello world!`, Children: [], Props: ``)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.node.String()
			if got != tt.want {
				t.Errorf("HTMLNode.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
