package markdown

import (
	"regexp"
	"strconv"
	"strings"
)

func MarkdownToBlocks(markdown string) []string {
	if markdown == "" {
		return []string{}
	}

	lines := strings.Split(markdown, "\n")

	blocks := []string{}
	var sb strings.Builder

	for _, line := range lines {
		if line != "" {
			if sb.Len() != 0 {
				sb.WriteByte('\n')
			}

			sb.WriteString(strings.TrimSpace(line))
			continue
		}

		if sb.Len() != 0 {
			blocks = append(blocks, sb.String())
			sb.Reset()
		}
	}

	if sb.Len() != 0 {
		blocks = append(blocks, sb.String())
		sb.Reset()
	}

	return blocks
}

var (
	// Headings start with 1-6 # characters, followed by a space and then the heading text.
	markdownHeadingRegexp = regexp.MustCompile("^#{1,6} .+")

	// Code blocks must start with 3 backticks (```) and end with 3 backticks (```).
	markdownCodeRegexp = regexp.MustCompile("(?s)^```.*```$")

	// Every line in a quote block must start with a > character.
	markdownQuoteRegexp = regexp.MustCompile("^(>.*)(?:\n>.*)*$")

	// Every line in an unordered list block must start with a * or - character, followed by a space.
	markdownUnorderedListRegexp = regexp.MustCompile("^([\\*-] .*)(?:\n[\\*-] .*)*$")

	// Every line in an ordered list block must start with a number followed by a . character and a space.
	markdownOrderedListRegexp = regexp.MustCompile("^(\\d+\\. .*)(?:\n\\d+\\. .*)*$")
)

func MarkdownBlockToBlockType(block string) MarkdownBlockType {
	if markdownHeadingRegexp.MatchString(block) {
		return MARKDOWN_BLOCK_TYPE_HEADING
	}

	if markdownCodeRegexp.MatchString(block) {
		return MARKDOWN_BLOCK_TYPE_CODE
	}

	if markdownQuoteRegexp.MatchString(block) {
		return MARKDOWN_BLOCK_TYPE_QUOTE
	}

	if markdownUnorderedListRegexp.MatchString(block) {
		return MARKDOWN_BLOCK_TYPE_UNORDERED_LIST
	}

	if isOrderedListBlockValid(block) {
		return MARKDOWN_BLOCK_TYPE_ORDERED_LIST
	}

	return MARKDOWN_BLOCK_TYPE_PARAGRAPH
}

func isOrderedListBlockValid(block string) bool {
	if !markdownOrderedListRegexp.MatchString(block) {
		return false
	}

	// Numbers at the start of lines must start at 1 and increment by 1 for each line.

	for i, line := range strings.Split(block, "\n") {
		lineNumberStr, _, found := strings.Cut(line, ".")

		if !found {
			return false
		}

		lineNumber, err := strconv.Atoi(lineNumberStr)

		if err != nil || lineNumber != i+1 {
			return false
		}
	}

	return true
}
