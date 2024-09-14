package markdown

import (
	"reflect"
	"testing"
)

func TestExtractMarkdownImages(t *testing.T) {
	tests := []struct {
		name string
		text string
		want []MarkdownTextUrlPair
	}{
		{
			name: "shouldReturnEmptySliceForEmptyText",
			text: "",
			want: []MarkdownTextUrlPair{},
		},
		{
			name: "shouldReturnEmptySliceForNoMatches",
			text: "This is text with a link [to google](https://www.google.com) and [to youtube](https://www.youtube.com)",
			want: []MarkdownTextUrlPair{},
		},
		{
			name: "shouldCaptureImagesButNotLinks",
			text: "This is text with a ![rick roll](https://i.imgur.com/aKaOqIh.gif) and ![obi wan](https://i.imgur.com/fJRm4Vk.jpeg) images but also includes link [to google](https://www.google.com) that should not be captured",
			want: []MarkdownTextUrlPair{
				{"rick roll", "https://i.imgur.com/aKaOqIh.gif"},
				{"obi wan", "https://i.imgur.com/fJRm4Vk.jpeg"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractMarkdownImages(tt.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractMarkdownImages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractMarkdownLinks(t *testing.T) {
	tests := []struct {
		name string
		text string
		want []MarkdownTextUrlPair
	}{
		{
			name: "shouldReturnEmptySliceForEmptyText",
			text: "",
			want: []MarkdownTextUrlPair{},
		},
		{
			name: "shouldReturnEmptySliceForNoMatches",
			text: "This is text with a ![rick roll](https://i.imgur.com/aKaOqIh.gif) and ![obi wan](https://i.imgur.com/fJRm4Vk.jpeg) images",
			want: []MarkdownTextUrlPair{},
		},
		{
			name: "shouldCaptureLinksButNotImages",
			text: "This is text with a link [to google](https://www.google.com) and [to youtube](https://www.youtube.com) but also includes ![rick roll](https://i.imgur.com/aKaOqIh.gif) image that should not be captured",
			want: []MarkdownTextUrlPair{
				{"to google", "https://www.google.com"},
				{"to youtube", "https://www.youtube.com"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractMarkdownLinks(tt.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractMarkdownLinks() = %v, want %v", got, tt.want)
			}
		})
	}
}
