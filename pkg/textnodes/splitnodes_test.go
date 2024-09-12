package textnodes

import (
	"errors"
	"fmt"
	"testing"
)

func TestSplitNodesByDelimiter(t *testing.T) {
	tests := []struct {
		name      string
		oldNodes  []*TextNode
		delimiter string
		textType  string
		wantStr   string
		wantErr   error
	}{
		{
			name:      "shouldReturnErrInvalidMarkdownSyntax",
			oldNodes:  []*TextNode{{"This text `has invalid delimiter", TEXT_NODE_TYPE_TEXT, ""}},
			delimiter: "`",
			textType:  TEXT_NODE_TYPE_CODE,
			wantStr:   "[]",
			wantErr:   ErrInvalidMarkdownSyntax,
		},
		{
			name:      "shouldReturnEmptySliceForNilOldNodes",
			oldNodes:  nil,
			delimiter: "`",
			textType:  TEXT_NODE_TYPE_CODE,
			wantStr:   "[]",
			wantErr:   nil,
		},
		{
			name:      "shouldReturnEmptySliceForEmptyOldNodes",
			oldNodes:  []*TextNode{},
			delimiter: "`",
			textType:  TEXT_NODE_TYPE_CODE,
			wantStr:   "[]",
			wantErr:   nil,
		},
		{
			name:      "shouldReturnOldNodesSliceForEmptyDelimiter",
			oldNodes:  []*TextNode{{"Hello world", TEXT_NODE_TYPE_TEXT, ""}},
			delimiter: "",
			textType:  TEXT_NODE_TYPE_CODE,
			wantStr:   "[TextNode(Hello world, text)]",
			wantErr:   nil,
		},
		{
			name:      "shouldReturnSameTextNodeForNonTextTypeNode",
			oldNodes:  []*TextNode{{"Hello `world`!", TEXT_NODE_TYPE_ITALIC, ""}},
			delimiter: "`",
			textType:  TEXT_NODE_TYPE_CODE,
			wantStr:   "[TextNode(Hello `world`!, italic)]",
			wantErr:   nil,
		},
		{
			name:      "shouldReturnSameTextNodeForNoDelimiterMatch",
			oldNodes:  []*TextNode{{"Hello `world`!", TEXT_NODE_TYPE_TEXT, ""}},
			delimiter: "*",
			textType:  TEXT_NODE_TYPE_ITALIC,
			wantStr:   "[TextNode(Hello `world`!, text)]",
			wantErr:   nil,
		},
		{
			name:      "shouldSplitOnBoldDelimiter",
			oldNodes:  []*TextNode{{"Hello **world**! `Just` **some** words", TEXT_NODE_TYPE_TEXT, ""}},
			delimiter: "**",
			textType:  TEXT_NODE_TYPE_BOLD,
			wantStr:   "[TextNode(Hello , text) TextNode(world, bold) TextNode(! `Just` , text) TextNode(some, bold) TextNode( words, text)]",
			wantErr:   nil,
		},
		{
			name:      "shouldSplitOnItalicDelimiter",
			oldNodes:  []*TextNode{{"Hello *world*! `Just` *some* words", TEXT_NODE_TYPE_TEXT, ""}},
			delimiter: "*",
			textType:  TEXT_NODE_TYPE_ITALIC,
			wantStr:   "[TextNode(Hello , text) TextNode(world, italic) TextNode(! `Just` , text) TextNode(some, italic) TextNode( words, text)]",
			wantErr:   nil,
		},
		{
			name: "shouldSplitOnInlineCodeBlockDelimiter",
			oldNodes: []*TextNode{
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
			newNodes, gotErr := SplitNodesByDelimiter(tt.oldNodes, tt.delimiter, tt.textType)
			gotStr := fmt.Sprintf("%v", newNodes)

			if gotStr != tt.wantStr || !errors.Is(gotErr, tt.wantErr) {
				t.Errorf(
					"SplitNodesByDelimiter(%v, %s, %s) = (%v, %v), want (%v, %v)",
					tt.oldNodes, tt.delimiter, tt.textType, gotStr, gotErr, tt.wantStr, tt.wantErr,
				)
			}
		})
	}

	t.Run("shouldSplitNodesMultipleTimesOnDifferentDelimiters", func(t *testing.T) {
		oldNodes := []*TextNode{
			{"This is *text* with a `code block` word", TEXT_NODE_TYPE_TEXT, ""},
			{"**Click me** and go to *about* page!", TEXT_NODE_TYPE_LINK, "/about"},
			{"This *is text* **with** a `code block` word", TEXT_NODE_TYPE_TEXT, ""},
			{"This *is text* **with** a `code block` word", TEXT_NODE_TYPE_CODE, ""},
			{"This is text with a `code block` word", TEXT_NODE_TYPE_ITALIC, ""},
			{"This is text with a `code block word`", TEXT_NODE_TYPE_TEXT, ""},
			{"`This is text with a code block word`", TEXT_NODE_TYPE_TEXT, ""},
		}

		newNodes, err := SplitNodesByDelimiter(oldNodes, "**", TEXT_NODE_TYPE_BOLD)
		gotStr := fmt.Sprintf("%v", newNodes)
		wantStr := "[TextNode(This is *text* with a `code block` word, text) TextNode(**Click me** and go to *about* page!, link, /about) TextNode(This *is text* , text) TextNode(with, bold) TextNode( a `code block` word, text) TextNode(This *is text* **with** a `code block` word, code) TextNode(This is text with a `code block` word, italic) TextNode(This is text with a `code block word`, text) TextNode(`This is text with a code block word`, text)]"

		if gotStr != wantStr || err != nil {
			t.Errorf(
				"SplitNodesByDelimiter(%v, %s, %s) = (%v, %v), want (%v, nil)",
				oldNodes, "**", TEXT_NODE_TYPE_BOLD, gotStr, err, wantStr,
			)
		}

		newNodes, err = SplitNodesByDelimiter(newNodes, "*", TEXT_NODE_TYPE_ITALIC)
		gotStr = fmt.Sprintf("%v", newNodes)
		wantStr = "[TextNode(This is , text) TextNode(text, italic) TextNode( with a `code block` word, text) TextNode(**Click me** and go to *about* page!, link, /about) TextNode(This , text) TextNode(is text, italic) TextNode( , text) TextNode(with, bold) TextNode( a `code block` word, text) TextNode(This *is text* **with** a `code block` word, code) TextNode(This is text with a `code block` word, italic) TextNode(This is text with a `code block word`, text) TextNode(`This is text with a code block word`, text)]"

		if gotStr != wantStr || err != nil {
			t.Errorf(
				"SplitNodesByDelimiter(%v, %s, %s) = (%v, %v), want (%v, nil)",
				oldNodes, "*", TEXT_NODE_TYPE_ITALIC, gotStr, err, wantStr,
			)
		}

		newNodes, err = SplitNodesByDelimiter(newNodes, "`", TEXT_NODE_TYPE_CODE)
		gotStr = fmt.Sprintf("%v", newNodes)
		wantStr = "[TextNode(This is , text) TextNode(text, italic) TextNode( with a , text) TextNode(code block, code) TextNode( word, text) TextNode(**Click me** and go to *about* page!, link, /about) TextNode(This , text) TextNode(is text, italic) TextNode( , text) TextNode(with, bold) TextNode( a , text) TextNode(code block, code) TextNode( word, text) TextNode(This *is text* **with** a `code block` word, code) TextNode(This is text with a `code block` word, italic) TextNode(This is text with a , text) TextNode(code block word, code) TextNode(, text) TextNode(, text) TextNode(This is text with a code block word, code) TextNode(, text)]"

		if gotStr != wantStr || err != nil {
			t.Errorf(
				"SplitNodesByDelimiter(%v, %s, %s) = (%v, %v), want (%v, %v)",
				oldNodes, "`", TEXT_NODE_TYPE_CODE, gotStr, err, wantStr, nil,
			)
		}
	})
}
