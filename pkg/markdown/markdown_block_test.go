package markdown

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMarkdownToBlocks(t *testing.T) {
	tests := []struct {
		name string
		text string
		want []string
	}{
		{
			name: "shouldReturnEmptySliceForEmptyText",
			text: "",
			want: []string{},
		},
		{
			name: "shouldReturnEmptySliceForEmptyLines",
			text: "\n \t \n    \t\t\t\n\n",
			want: []string{},
		},
		{
			name: "shouldGroupBlocks",
			text: "# This is a heading\n\nThis is a paragraph of text. It has some **bold** and *italic* words inside of it.\n\n* This is the first list item in a list block\n* This is a list item\n* This is another list item",
			want: []string{
				"# This is a heading",
				"This is a paragraph of text. It has some **bold** and *italic* words inside of it.",
				"* This is the first list item in a list block\n* This is a list item\n* This is another list item",
			},
		},
		{
			name: "shouldGroupAndSanitizeBlocks",
			text: "   # This is a heading   \n\n       \t      \n\n \t  This is a paragraph of text. It has some **bold** and *italic* words inside of it.\n\n* This is the first list item in a list block\n  * This is a list item     \n * This is another list item   \t  \n",
			want: []string{
				"# This is a heading",
				"This is a paragraph of text. It has some **bold** and *italic* words inside of it.",
				"* This is the first list item in a list block\n* This is a list item\n* This is another list item",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MarkdownToBlocks(tt.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarkdownToBlocks(%s) = %v, want %v", tt.text, got, tt.want)
			}
		})
	}
}

func TestMarkdownBlockToBlockType(t *testing.T) {
	type testGroup []struct {
		block string
		want  MarkdownBlockType
	}

	tests := map[string]testGroup{
		"headingTests": {
			{
				block: "# Hello",
				want:  MARKDOWN_BLOCK_TYPE_HEADING,
			},
			{
				block: "###### Hello",
				want:  MARKDOWN_BLOCK_TYPE_HEADING,
			},
			{
				block: "#",
				want:  MARKDOWN_BLOCK_TYPE_PARAGRAPH,
			},
			{
				block: "######Title is not separated from #s with a space",
				want:  MARKDOWN_BLOCK_TYPE_PARAGRAPH,
			},
			{
				block: "####### Too much #s",
				want:  MARKDOWN_BLOCK_TYPE_PARAGRAPH,
			},
			{
				block: ".# Symbols before #",
				want:  MARKDOWN_BLOCK_TYPE_PARAGRAPH,
			},
		},
		"codeTests": {
			{
				block: "```\nmultiline code goes here\nsecond line of code\n```",
				want:  MARKDOWN_BLOCK_TYPE_CODE,
			},
			{
				block: "```code goes here```",
				want:  MARKDOWN_BLOCK_TYPE_CODE,
			},
			{
				block: "``````",
				want:  MARKDOWN_BLOCK_TYPE_CODE,
			},
			{
				block: "`````",
				want:  MARKDOWN_BLOCK_TYPE_PARAGRAPH,
			},
			{
				block: "```doesn't end with ```backticks",
				want:  MARKDOWN_BLOCK_TYPE_PARAGRAPH,
			},
			{
				block: "doesn't ```start end with backticks```",
				want:  MARKDOWN_BLOCK_TYPE_PARAGRAPH,
			},
		},
		"quoteTests": {
			{
				block: ">Single line quote",
				want:  MARKDOWN_BLOCK_TYPE_QUOTE,
			},
			{
				block: "> Single line quote",
				want:  MARKDOWN_BLOCK_TYPE_QUOTE,
			},
			{
				block: ">Multiline quote\n>second line\n>third line",
				want:  MARKDOWN_BLOCK_TYPE_QUOTE,
			},
			{
				block: ">>Still a valid quote block",
				want:  MARKDOWN_BLOCK_TYPE_QUOTE,
			},
			{
				block: ".> doesn't start with > character",
				want:  MARKDOWN_BLOCK_TYPE_PARAGRAPH,
			},
			{
				block: ">first line is valid\n.> invalid second line\n>valid last line",
				want:  MARKDOWN_BLOCK_TYPE_PARAGRAPH,
			},
		},
		"unorderedListTests": {
			{
				block: "* Single line of unordered list with * char",
				want:  MARKDOWN_BLOCK_TYPE_UNORDERED_LIST,
			},
			{
				block: "- Single line of unordered list with - char",
				want:  MARKDOWN_BLOCK_TYPE_UNORDERED_LIST,
			},
			{
				block: "* Multiline unordered list with * chars\n* second unordered list item",
				want:  MARKDOWN_BLOCK_TYPE_UNORDERED_LIST,
			},
			{
				block: "- Multiline unordered list with - chars\n- second unordered list item",
				want:  MARKDOWN_BLOCK_TYPE_UNORDERED_LIST,
			},
			{
				block: "* Mixed multiline unordered list\n- with * and - chars\n* last line",
				want:  MARKDOWN_BLOCK_TYPE_UNORDERED_LIST,
			},
			{
				block: "*no space after unordered list item indicator",
				want:  MARKDOWN_BLOCK_TYPE_PARAGRAPH,
			},
			{
				block: "- no space after unordered list item indicator\n*on second line",
				want:  MARKDOWN_BLOCK_TYPE_PARAGRAPH,
			},
			{
				block: ".- doesn't start with unordered list item indicator",
				want:  MARKDOWN_BLOCK_TYPE_PARAGRAPH,
			},
		},
		"orderedListTests": {
			{
				block: "1. Single line of ordered list",
				want:  MARKDOWN_BLOCK_TYPE_ORDERED_LIST,
			},
			{
				block: "1. Multiple lines of ordered list\n2. Second\n3. Third",
				want:  MARKDOWN_BLOCK_TYPE_ORDERED_LIST,
			},
			{
				block: "1.No space after ordered list item index",
				want:  MARKDOWN_BLOCK_TYPE_PARAGRAPH,
			},
			{
				block: "1 No dot after ordered list item index",
				want:  MARKDOWN_BLOCK_TYPE_PARAGRAPH,
			},
			{
				block: "2. Single line of ordered list starting at wrong index",
				want:  MARKDOWN_BLOCK_TYPE_PARAGRAPH,
			},
			{
				block: "1. multiple lines of ordered list\n3. containing wrong index\n4. in middle",
				want:  MARKDOWN_BLOCK_TYPE_PARAGRAPH,
			},
		},
	}

	for groupName, group := range tests {
		for i, tt := range group {
			t.Run(fmt.Sprintf("%s#%v", groupName, i+1), func(t *testing.T) {
				if got := MarkdownBlockToBlockType(tt.block); got != tt.want {
					t.Errorf(
						"MarkdownBlockToBlockType(%s) = %v, want %v",
						tt.block, got, tt.want,
					)
				}
			})
		}
	}
}
