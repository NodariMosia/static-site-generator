package utils

import "regexp"

var (
	markdownImagesRegexp = regexp.MustCompile(`!\[(.*?)\]\((.*?)\)`)
	markdownLinksRegexp  = regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)
)

type MarkdownTextUrlPair struct{ Text, Url string }

func ExtractMarkdownImages(text string) []MarkdownTextUrlPair {
	submatches := markdownImagesRegexp.FindAllStringSubmatch(text, -1)
	result := make([]MarkdownTextUrlPair, len(submatches))

	for i, submatch := range submatches {
		imgAlt, imgSrc := submatch[1], submatch[2]
		result[i] = MarkdownTextUrlPair{imgAlt, imgSrc}
	}

	return result
}

func ExtractMarkdownLinks(text string) []MarkdownTextUrlPair {
	listOfSubmatchIndices := markdownLinksRegexp.FindAllStringSubmatchIndex(text, -1)
	result := []MarkdownTextUrlPair{}

	for _, submatchIndices := range listOfSubmatchIndices {
		fullStart, _ := submatchIndices[0], submatchIndices[1]
		linkTextStart, linkTextEnd := submatchIndices[2], submatchIndices[3]
		linkHrefStart, linkHrefEnd := submatchIndices[4], submatchIndices[5]

		isImageBlock := fullStart != 0 && text[fullStart-1] == '!'
		if isImageBlock {
			continue
		}

		linkText, linkHref := text[linkTextStart:linkTextEnd], text[linkHrefStart:linkHrefEnd]
		result = append(result, MarkdownTextUrlPair{linkText, linkHref})
	}

	return result
}
