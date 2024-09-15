package markdown

import (
	"errors"
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

func TestExtractMarkdownTitle(t *testing.T) {
	tests := []struct {
		name      string
		markdown  string
		wantTitle string
		wantErr   error
	}{
		{
			name:      "shouldExtractTitleFromSingleLineMarkdown",
			markdown:  "# Hello world!",
			wantTitle: "Hello world!",
			wantErr:   nil,
		},
		{
			name:      "shouldExtractTitleFromMultilineMarkdown",
			markdown:  "## This\n#should be extracted\n# Hello world!\n> from multiline",
			wantTitle: "Hello world!",
			wantErr:   nil,
		},
		{
			name:      "shouldExtractAndTrimTitle",
			markdown:  "## This\n#should be trimmed\n#  \tHello world!\t\t \n# Second title\n\n",
			wantTitle: "Hello world!",
			wantErr:   nil,
		},
		{
			name:      "shouldReturnErrMissingMarkdownTitle",
			markdown:  "## This text\n#doesn't have\n#\tany valid\ntitles\n#\n# \n\n## \n",
			wantTitle: "",
			wantErr:   ErrMissingMarkdownTitle,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTitle, gotErr := ExtractMarkdownTitle(tt.markdown)

			if gotTitle != tt.wantTitle || !errors.Is(gotErr, tt.wantErr) {
				t.Errorf(
					"ExtractMarkdownTitle() = (%s, %v), want (%s, %v)",
					gotTitle, gotErr, tt.wantTitle, tt.wantErr,
				)
			}
		})
	}
}
