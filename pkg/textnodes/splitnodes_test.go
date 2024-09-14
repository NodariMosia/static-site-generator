package textnodes

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestSplitNodesByDelimiter(t *testing.T) {
	tests := []struct {
		name      string
		nodes     []*TextNode
		delimiter string
		textType  TextNodeType
		wantStr   string
		wantErr   error
	}{
		{
			name:      "shouldReturnErrInvalidMarkdownSyntax",
			nodes:     []*TextNode{{"This text `has invalid delimiter", TEXT_NODE_TYPE_TEXT, ""}},
			delimiter: "`",
			textType:  TEXT_NODE_TYPE_CODE,
			wantStr:   "[]",
			wantErr:   ErrInvalidMarkdownSyntax,
		},
		{
			name:      "shouldReturnEmptySliceForNilNodes",
			nodes:     nil,
			delimiter: "`",
			textType:  TEXT_NODE_TYPE_CODE,
			wantStr:   "[]",
			wantErr:   nil,
		},
		{
			name:      "shouldReturnEmptySliceForEmptyNodes",
			nodes:     []*TextNode{},
			delimiter: "`",
			textType:  TEXT_NODE_TYPE_CODE,
			wantStr:   "[]",
			wantErr:   nil,
		},
		{
			name:      "shouldReturnNodesSliceForEmptyDelimiter",
			nodes:     []*TextNode{{"Hello world", TEXT_NODE_TYPE_TEXT, ""}},
			delimiter: "",
			textType:  TEXT_NODE_TYPE_CODE,
			wantStr:   "[TextNode(Hello world, text)]",
			wantErr:   nil,
		},
		{
			name:      "shouldReturnSameTextNodeForNonTextTypeNode",
			nodes:     []*TextNode{{"Hello `world`!", TEXT_NODE_TYPE_ITALIC, ""}},
			delimiter: "`",
			textType:  TEXT_NODE_TYPE_CODE,
			wantStr:   "[TextNode(Hello `world`!, italic)]",
			wantErr:   nil,
		},
		{
			name:      "shouldReturnSameTextNodeForNoDelimiterMatch",
			nodes:     []*TextNode{{"Hello `world`!", TEXT_NODE_TYPE_TEXT, ""}},
			delimiter: "*",
			textType:  TEXT_NODE_TYPE_ITALIC,
			wantStr:   "[TextNode(Hello `world`!, text)]",
			wantErr:   nil,
		},
		{
			name:      "shouldSplitOnBoldDelimiter",
			nodes:     []*TextNode{{"Hello **world**! `Just` **some** words", TEXT_NODE_TYPE_TEXT, ""}},
			delimiter: "**",
			textType:  TEXT_NODE_TYPE_BOLD,
			wantStr:   "[TextNode(Hello , text) TextNode(world, bold) TextNode(! `Just` , text) TextNode(some, bold) TextNode( words, text)]",
			wantErr:   nil,
		},
		{
			name:      "shouldSplitOnItalicDelimiter",
			nodes:     []*TextNode{{"Hello *world*! `Just` *some* words", TEXT_NODE_TYPE_TEXT, ""}},
			delimiter: "*",
			textType:  TEXT_NODE_TYPE_ITALIC,
			wantStr:   "[TextNode(Hello , text) TextNode(world, italic) TextNode(! `Just` , text) TextNode(some, italic) TextNode( words, text)]",
			wantErr:   nil,
		},
		{
			name: "shouldSplitOnInlineCodeBlockDelimiter",
			nodes: []*TextNode{
				{"Hello `world*! *Just` `some` words", TEXT_NODE_TYPE_TEXT, ""},
				{"Hello *world*! `Just` *some* words", TEXT_NODE_TYPE_TEXT, ""},
			},
			delimiter: "`",
			textType:  TEXT_NODE_TYPE_CODE,
			wantStr:   "[TextNode(Hello , text) TextNode(world*! *Just, code) TextNode( , text) TextNode(some, code) TextNode( words, text) TextNode(Hello *world*! , text) TextNode(Just, code) TextNode( *some* words, text)]",
			wantErr:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newNodes, gotErr := SplitNodesByDelimiter(tt.nodes, tt.delimiter, tt.textType)
			gotStr := fmt.Sprintf("%v", newNodes)

			if gotStr != tt.wantStr || !errors.Is(gotErr, tt.wantErr) {
				t.Errorf(
					"SplitNodesByDelimiter(%v, %s, %s) = (%v, %v), want (%v, %v)",
					tt.nodes, tt.delimiter, tt.textType, gotStr, gotErr, tt.wantStr, tt.wantErr,
				)
			}
		})
	}

	t.Run("shouldSplitNodesMultipleTimesOnDifferentDelimiters", func(t *testing.T) {
		nodes := []*TextNode{
			{"This is *text* with a `code block` word", TEXT_NODE_TYPE_TEXT, ""},
			{"**Click me** and go to *about* page!", TEXT_NODE_TYPE_LINK, "/about"},
			{"This *is text* **with** a `code block` word", TEXT_NODE_TYPE_TEXT, ""},
			{"This *is text* **with** a `code block` word", TEXT_NODE_TYPE_CODE, ""},
			{"This is text with a `code block` word", TEXT_NODE_TYPE_ITALIC, ""},
			{"This is text with a `code block word`", TEXT_NODE_TYPE_TEXT, ""},
			{"`This is text with a code block word`", TEXT_NODE_TYPE_TEXT, ""},
		}

		newNodes, err := SplitNodesByDelimiter(nodes, "**", TEXT_NODE_TYPE_BOLD)
		gotStr := fmt.Sprintf("%v", newNodes)
		wantStr := "[TextNode(This is *text* with a `code block` word, text) TextNode(**Click me** and go to *about* page!, link, /about) TextNode(This *is text* , text) TextNode(with, bold) TextNode( a `code block` word, text) TextNode(This *is text* **with** a `code block` word, code) TextNode(This is text with a `code block` word, italic) TextNode(This is text with a `code block word`, text) TextNode(`This is text with a code block word`, text)]"

		if gotStr != wantStr || err != nil {
			t.Errorf(
				"SplitNodesByDelimiter(%v, %s, %s) = (%v, %v), want (%v, nil)",
				nodes, "**", TEXT_NODE_TYPE_BOLD, gotStr, err, wantStr,
			)
		}

		newNodes, err = SplitNodesByDelimiter(newNodes, "*", TEXT_NODE_TYPE_ITALIC)
		gotStr = fmt.Sprintf("%v", newNodes)
		wantStr = "[TextNode(This is , text) TextNode(text, italic) TextNode( with a `code block` word, text) TextNode(**Click me** and go to *about* page!, link, /about) TextNode(This , text) TextNode(is text, italic) TextNode( , text) TextNode(with, bold) TextNode( a `code block` word, text) TextNode(This *is text* **with** a `code block` word, code) TextNode(This is text with a `code block` word, italic) TextNode(This is text with a `code block word`, text) TextNode(`This is text with a code block word`, text)]"

		if gotStr != wantStr || err != nil {
			t.Errorf(
				"SplitNodesByDelimiter(%v, %s, %s) = (%v, %v), want (%v, nil)",
				nodes, "*", TEXT_NODE_TYPE_ITALIC, gotStr, err, wantStr,
			)
		}

		newNodes, err = SplitNodesByDelimiter(newNodes, "`", TEXT_NODE_TYPE_CODE)
		gotStr = fmt.Sprintf("%v", newNodes)
		wantStr = "[TextNode(This is , text) TextNode(text, italic) TextNode( with a , text) TextNode(code block, code) TextNode( word, text) TextNode(**Click me** and go to *about* page!, link, /about) TextNode(This , text) TextNode(is text, italic) TextNode( , text) TextNode(with, bold) TextNode( a , text) TextNode(code block, code) TextNode( word, text) TextNode(This *is text* **with** a `code block` word, code) TextNode(This is text with a `code block` word, italic) TextNode(This is text with a , text) TextNode(code block word, code) TextNode(, text) TextNode(, text) TextNode(This is text with a code block word, code) TextNode(, text)]"

		if gotStr != wantStr || err != nil {
			t.Errorf(
				"SplitNodesByDelimiter(%v, %s, %s) = (%v, %v), want (%v, %v)",
				nodes, "`", TEXT_NODE_TYPE_CODE, gotStr, err, wantStr, nil,
			)
		}
	})
}

func TestSplitNodesByImages(t *testing.T) {
	tests := []struct {
		name  string
		nodes []*TextNode
		want  []*TextNode
	}{
		{
			name:  "shouldReturnEmptySliceForNoNodes",
			nodes: []*TextNode{},
			want:  []*TextNode{},
		},
		{
			name: "shouldReturnSameSliceForNoImages",
			nodes: []*TextNode{
				{"This is text with [no images](https://google.com)!", TEXT_NODE_TYPE_TEXT, ""},
			},
			want: []*TextNode{
				{"This is text with [no images](https://google.com)!", TEXT_NODE_TYPE_TEXT, ""},
			},
		},
		{
			name: "shouldSplitOneImage",
			nodes: []*TextNode{
				{"This is text with a ![rick roll](https://i.imgur.com/aKaOqIh.gif)!", TEXT_NODE_TYPE_TEXT, ""},
			},
			want: []*TextNode{
				{"This is text with a ", TEXT_NODE_TYPE_TEXT, ""},
				{"rick roll", TEXT_NODE_TYPE_IMAGE, "https://i.imgur.com/aKaOqIh.gif"},
				{"!", TEXT_NODE_TYPE_TEXT, ""},
			},
		},
		{
			name: "shouldSplitMultipleImages",
			nodes: []*TextNode{
				{"This is text with a ![rick roll](https://i.imgur.com/aKaOqIh.gif) and ![obi wan](https://i.imgur.com/fJRm4Vk.jpeg)", TEXT_NODE_TYPE_TEXT, ""},
			},
			want: []*TextNode{
				{"This is text with a ", TEXT_NODE_TYPE_TEXT, ""},
				{"rick roll", TEXT_NODE_TYPE_IMAGE, "https://i.imgur.com/aKaOqIh.gif"},
				{" and ", TEXT_NODE_TYPE_TEXT, ""},
				{"obi wan", TEXT_NODE_TYPE_IMAGE, "https://i.imgur.com/fJRm4Vk.jpeg"},
			},
		},
		{
			name: "shouldSplitMultipleNodes",
			nodes: []*TextNode{
				{"This is text with [no images](https://google.com)!", TEXT_NODE_TYPE_TEXT, ""},
				{"", TEXT_NODE_TYPE_TEXT, ""},
				{"This is text with a ![rick roll](https://i.imgur.com/aKaOqIh.gif)!", TEXT_NODE_TYPE_TEXT, ""},
				{"This is bold text with a ![rick roll](https://i.imgur.com/aKaOqIh.gif)!", TEXT_NODE_TYPE_BOLD, ""},
				{"This is text with a ![rick roll](https://i.imgur.com/aKaOqIh.gif) and ![obi wan](https://i.imgur.com/fJRm4Vk.jpeg)", TEXT_NODE_TYPE_TEXT, ""},
			},
			want: []*TextNode{
				{"This is text with [no images](https://google.com)!", TEXT_NODE_TYPE_TEXT, ""},
				{"This is text with a ", TEXT_NODE_TYPE_TEXT, ""},
				{"rick roll", TEXT_NODE_TYPE_IMAGE, "https://i.imgur.com/aKaOqIh.gif"},
				{"!", TEXT_NODE_TYPE_TEXT, ""},
				{"This is bold text with a ![rick roll](https://i.imgur.com/aKaOqIh.gif)!", TEXT_NODE_TYPE_BOLD, ""},
				{"This is text with a ", TEXT_NODE_TYPE_TEXT, ""},
				{"rick roll", TEXT_NODE_TYPE_IMAGE, "https://i.imgur.com/aKaOqIh.gif"},
				{" and ", TEXT_NODE_TYPE_TEXT, ""},
				{"obi wan", TEXT_NODE_TYPE_IMAGE, "https://i.imgur.com/fJRm4Vk.jpeg"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitNodesByImages(tt.nodes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitNodesByImages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitNodesByLinks(t *testing.T) {
	tests := []struct {
		name  string
		nodes []*TextNode
		want  []*TextNode
	}{
		{
			name:  "shouldReturnEmptySliceForNoNodes",
			nodes: []*TextNode{},
			want:  []*TextNode{},
		},
		{
			name: "shouldReturnSameSliceForNoLinks",
			nodes: []*TextNode{
				{"This is text with ![no links](https://i.imgur.com/aKaOqIh.gif)!", TEXT_NODE_TYPE_TEXT, ""},
			},
			want: []*TextNode{
				{"This is text with ![no links](https://i.imgur.com/aKaOqIh.gif)!", TEXT_NODE_TYPE_TEXT, ""},
			},
		},
		{
			name: "shouldSplitOneLink",
			nodes: []*TextNode{
				{"This is text with a link [to gmail](https://gmail.com)!", TEXT_NODE_TYPE_TEXT, ""},
			},
			want: []*TextNode{
				{"This is text with a link ", TEXT_NODE_TYPE_TEXT, ""},
				{"to gmail", TEXT_NODE_TYPE_LINK, "https://gmail.com"},
				{"!", TEXT_NODE_TYPE_TEXT, ""},
			},
		},
		{
			name: "shouldSplitMultipleLinks",
			nodes: []*TextNode{
				{"This is text with a link [to google](https://google.com) and [to youtube](https://youtube.com)", TEXT_NODE_TYPE_TEXT, ""},
			},
			want: []*TextNode{
				{"This is text with a link ", TEXT_NODE_TYPE_TEXT, ""},
				{"to google", TEXT_NODE_TYPE_LINK, "https://google.com"},
				{" and ", TEXT_NODE_TYPE_TEXT, ""},
				{"to youtube", TEXT_NODE_TYPE_LINK, "https://youtube.com"},
			},
		},
		{
			name: "shouldSplitMultipleNodes",
			nodes: []*TextNode{
				{"This is text with ![no links](https://i.imgur.com/aKaOqIh.gif)!", TEXT_NODE_TYPE_TEXT, ""},
				{"", TEXT_NODE_TYPE_TEXT, ""},
				{"This is text with a link [to gmail](https://gmail.com)!", TEXT_NODE_TYPE_TEXT, ""},
				{"This is bold text with a link [to gmail](https://gmail.com)!", TEXT_NODE_TYPE_BOLD, ""},
				{"This is text with a link [to google](https://google.com) and [to youtube](https://youtube.com)", TEXT_NODE_TYPE_TEXT, ""},
			},
			want: []*TextNode{
				{"This is text with ![no links](https://i.imgur.com/aKaOqIh.gif)!", TEXT_NODE_TYPE_TEXT, ""},
				{"This is text with a link ", TEXT_NODE_TYPE_TEXT, ""},
				{"to gmail", TEXT_NODE_TYPE_LINK, "https://gmail.com"},
				{"!", TEXT_NODE_TYPE_TEXT, ""},
				{"This is bold text with a link [to gmail](https://gmail.com)!", TEXT_NODE_TYPE_BOLD, ""},
				{"This is text with a link ", TEXT_NODE_TYPE_TEXT, ""},
				{"to google", TEXT_NODE_TYPE_LINK, "https://google.com"},
				{" and ", TEXT_NODE_TYPE_TEXT, ""},
				{"to youtube", TEXT_NODE_TYPE_LINK, "https://youtube.com"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitNodesByLinks(tt.nodes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitNodesByLinks() = %v, want %v", got, tt.want)
			}
		})
	}
}
