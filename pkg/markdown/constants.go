package markdown

type MarkdownBlockType string

const (
	MARKDOWN_BLOCK_TYPE_PARAGRAPH      MarkdownBlockType = "paragraph"
	MARKDOWN_BLOCK_TYPE_HEADING        MarkdownBlockType = "heading"
	MARKDOWN_BLOCK_TYPE_CODE           MarkdownBlockType = "code"
	MARKDOWN_BLOCK_TYPE_QUOTE          MarkdownBlockType = "quote"
	MARKDOWN_BLOCK_TYPE_UNORDERED_LIST MarkdownBlockType = "unordered_list"
	MARKDOWN_BLOCK_TYPE_ORDERED_LIST   MarkdownBlockType = "ordered_list"
)
