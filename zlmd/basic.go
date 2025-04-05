package zlmd

import (
	"fmt"
	"strings"
)

// Bold formats text as bold.
//
// Parameters:
//   - text (string): The text to format as bold
//
// Returns:
//   - string: The text wrapped in bold markdown syntax
//
// Algorithm:
//  1. Wrap the input text with double asterisks
//
// Example:
//
//	boldText := Bold("important")
//	// boldText will be "**important**"
//
// Notes:
//   - Uses the standard Markdown double asterisk syntax for bold
//   - Does not check if text already contains bold formatting
func Bold(text string) string {
	return fmt.Sprintf("**%s**", text)
}

// Italic formats text as italic.
//
// Parameters:
//   - text (string): The text to format as italic
//
// Returns:
//   - string: The text wrapped in italic markdown syntax
//
// Algorithm:
//  1. Wrap the input text with single asterisks
//
// Example:
//
//	italicText := Italic("emphasized")
//	// italicText will be "*emphasized*"
//
// Notes:
//   - Uses the standard Markdown single asterisk syntax for italic
//   - Does not check if text already contains italic formatting
func Italic(text string) string {
	return fmt.Sprintf("*%s*", text)
}

// Code formats text as inline code.
//
// Parameters:
//   - text (string): The text to format as inline code
//
// Returns:
//   - string: The text wrapped in inline code markdown syntax
//
// Algorithm:
//  1. Wrap the input text with backticks
//
// Example:
//
//	codeText := Code("var x = 10")
//	// codeText will be "`var x = 10`"
//
// Notes:
//   - Uses the standard Markdown backtick syntax for inline code
//   - Does not escape backticks within the text which could break formatting
func Code(text string) string {
	return fmt.Sprintf("`%s`", text)
}

// Link creates a markdown link.
//
// Parameters:
//   - text (string): The display text for the link
//   - url (string): The URL the link points to
//
// Returns:
//   - string: The formatted markdown link
//
// Algorithm:
//  1. Format the text and URL using Markdown link syntax
//
// Example:
//
//	linkText := Link("GitHub", "https://github.com")
//	// linkText will be "[GitHub](https://github.com)"
//
// Notes:
//   - Uses the standard Markdown link syntax [text](url)
//   - Does not validate or escape special characters in the URL
func Link(text, url string) string {
	return fmt.Sprintf("[%s](%s)", text, url)
}

// Image creates a markdown image.
//
// Parameters:
//   - altText (string): The alternative text for the image
//   - url (string): The URL of the image
//
// Returns:
//   - string: The formatted markdown image
//
// Algorithm:
//  1. Format the alt text and URL using Markdown image syntax
//
// Example:
//
//	imgText := Image("Logo", "https://example.com/logo.png")
//	// imgText will be "![Logo](https://example.com/logo.png)"
//
// Notes:
//   - Uses the standard Markdown image syntax ![alt](url)
//   - Does not validate or escape special characters in the URL
func Image(altText, url string) string {
	return fmt.Sprintf("![%s](%s)", altText, url)
}

// HorizontalRule returns a horizontal rule.
//
// Parameters:
//   - None
//
// Returns:
//   - string: The markdown horizontal rule
//
// Algorithm:
//  1. Return the Markdown horizontal rule syntax
//
// Example:
//
//	hr := HorizontalRule()
//	// hr will be "---"
//
// Notes:
//   - Uses the standard Markdown three-dash syntax for horizontal rules
//   - Can be used to visually separate sections of a document
func HorizontalRule() string {
	return "---"
}

// Heading creates a markdown heading with the specified level (1-6).
//
// Parameters:
//   - level (int): The heading level from 1 to 6
//   - text (string): The heading text
//
// Returns:
//   - string: The formatted markdown heading
//
// Algorithm:
//  1. Validate the heading level (ensure it's between 1 and 6)
//  2. Create a string with the appropriate number of hash symbols followed by a space and the text
//
// Example:
//
//	h2 := Heading(2, "Introduction")
//	// h2 will be "## Introduction"
//
// Notes:
//   - Automatically clamps level to range 1-6 if outside that range
//   - Level 1 is the highest level heading (# Title)
//   - Level 6 is the lowest level heading (###### Title)
func Heading(level int, text string) string {
	if level < 1 || level > 6 {
		level = 1
	}
	return fmt.Sprintf("%s %s", strings.Repeat("#", level), text)
}

// H1 creates a level 1 markdown heading.
//
// Parameters:
//   - text (string): The heading text
//
// Returns:
//   - string: The formatted level 1 markdown heading
//
// Algorithm:
//  1. Call Heading function with level 1 and the provided text
//
// Example:
//
//	heading := H1("Document Title")
//	// heading will be "# Document Title"
//
// Notes:
//   - Convenience function for creating top-level headings
//   - Equivalent to calling Heading(1, text)
func H1(text string) string {
	return Heading(1, text)
}

// H2 creates a level 2 markdown heading.
//
// Parameters:
//   - text (string): The heading text
//
// Returns:
//   - string: The formatted level 2 markdown heading
//
// Algorithm:
//  1. Call Heading function with level 2 and the provided text
//
// Example:
//
//	heading := H2("Section Title")
//	// heading will be "## Section Title"
//
// Notes:
//   - Convenience function for creating section headings
//   - Equivalent to calling Heading(2, text)
func H2(text string) string {
	return Heading(2, text)
}

// H3 creates a level 3 markdown heading.
//
// Parameters:
//   - text (string): The heading text
//
// Returns:
//   - string: The formatted level 3 markdown heading
//
// Algorithm:
//  1. Call Heading function with level 3 and the provided text
//
// Example:
//
//	heading := H3("Subsection Title")
//	// heading will be "### Subsection Title"
//
// Notes:
//   - Convenience function for creating subsection headings
//   - Equivalent to calling Heading(3, text)
func H3(text string) string {
	return Heading(3, text)
}

// H4 creates a level 4 markdown heading.
//
// Parameters:
//   - text (string): The heading text
//
// Returns:
//   - string: The formatted level 4 markdown heading
//
// Algorithm:
//  1. Call Heading function with level 4 and the provided text
//
// Example:
//
//	heading := H4("Sub-subsection Title")
//	// heading will be "#### Sub-subsection Title"
//
// Notes:
//   - Convenience function for creating lower-level headings
//   - Equivalent to calling Heading(4, text)
func H4(text string) string {
	return Heading(4, text)
}

// H5 creates a level 5 markdown heading.
//
// Parameters:
//   - text (string): The heading text
//
// Returns:
//   - string: The formatted level 5 markdown heading
//
// Algorithm:
//  1. Call Heading function with level 5 and the provided text
//
// Example:
//
//	heading := H5("Minor Section Title")
//	// heading will be "##### Minor Section Title"
//
// Notes:
//   - Convenience function for creating minor headings
//   - Equivalent to calling Heading(5, text)
func H5(text string) string {
	return Heading(5, text)
}

// H6 creates a level 6 markdown heading.
//
// Parameters:
//   - text (string): The heading text
//
// Returns:
//   - string: The formatted level 6 markdown heading
//
// Algorithm:
//  1. Call Heading function with level 6 and the provided text
//
// Example:
//
//	heading := H6("Smallest Section Title")
//	// heading will be "###### Smallest Section Title"
//
// Notes:
//   - Convenience function for creating the lowest level headings
//   - Equivalent to calling Heading(6, text)
func H6(text string) string {
	return Heading(6, text)
}

// QuoteBlock formats text as a blockquote.
//
// Parameters:
//   - text (string): The text to format as a blockquote
//
// Returns:
//   - string: The text formatted as a markdown blockquote
//
// Algorithm:
//  1. Split the input text into lines
//  2. Prepend each line with "> "
//  3. Join the lines back together with newlines
//
// Example:
//
//	quote := QuoteBlock("This is\na quote")
//	// quote will be:
//	// "> This is
//	// > a quote
//	// "
//
// Notes:
//   - Properly handles multi-line quotes by adding the quote prefix to each line
//   - Adds a trailing newline to the result
func QuoteBlock(text string) string {
	lines := strings.Split(text, "\n")
	quotedLines := make([]string, len(lines))
	for i, line := range lines {
		quotedLines[i] = fmt.Sprintf("> %s", line)
	}
	return strings.Join(quotedLines, "\n") + "\n"
}

// QuoteBlocknl formats text as a blockquote with an extra newline.
//
// Parameters:
//   - text (string): The text to format as a blockquote
//
// Returns:
//   - string: The text formatted as a markdown blockquote with an extra newline
//
// Algorithm:
//  1. Call QuoteBlock with the provided text
//  2. Append an additional newline
//
// Example:
//
//	quote := QuoteBlocknl("This is a quote")
//	// quote will be:
//	// "> This is a quote
//	//
//	// "
//
// Notes:
//   - Convenience function for adding a blockquote with extra spacing
//   - Useful for separating blockquotes from subsequent content
func QuoteBlocknl(text string) string {
	return QuoteBlock(text) + "\n"
}

// Paragraph formats text as a paragraph with trailing newlines.
//
// Parameters:
//   - text (string): The paragraph text
//
// Returns:
//   - string: The text with two trailing newlines
//
// Algorithm:
//  1. Append two newlines to the input text
//
// Example:
//
//	para := Paragraph("This is a paragraph.")
//	// para will be "This is a paragraph.\n\n"
//
// Notes:
//   - The double newlines ensure proper paragraph separation in Markdown
//   - Text is not otherwise modified
func Paragraph(text string) string {
	return text + "\n\n"
}

// P is a shortcut for Paragraph.
//
// Parameters:
//   - text (string): The paragraph text
//
// Returns:
//   - string: The text with two trailing newlines
//
// Algorithm:
//  1. Call Paragraph with the provided text
//
// Example:
//
//	para := P("This is a paragraph.")
//	// para will be "This is a paragraph.\n\n"
//
// Notes:
//   - Convenience alias for the Paragraph function
//   - Provides a shorter name for frequently used function
func P(text string) string {
	return Paragraph(text)
}

// BR is a shortcut for Break, adding a single newline after text.
//
// Parameters:
//   - text (string): The text to append a line break to
//
// Returns:
//   - string: The text with a single trailing newline
//
// Algorithm:
//  1. Append a single newline to the input text
//
// Example:
//
//	line := BR("This is a line.")
//	// line will be "This is a line.\n"
//
// Notes:
//   - Useful for adding line breaks without paragraph spacing
//   - Unlike Paragraph, adds only a single newline
func BR(text string) string {
	return text + "\n"
}

// WriteHeading writes a heading to the provided StringBuilder.
//
// Parameters:
//   - sb (*strings.Builder): The string builder to write to
//   - level (int): The heading level from 1 to 6
//   - text (string): The heading text
//
// Returns:
//   - None
//
// Algorithm:
//  1. Generate heading with the Heading function
//  2. Write the heading to the string builder
//  3. Add a newline after the heading
//
// Example:
//
//	var sb strings.Builder
//	WriteHeading(&sb, 2, "Introduction")
//	// sb now contains "## Introduction\n"
//
// Notes:
//   - Helper function for writing headings directly to a string builder
//   - Automatically adds a trailing newline
func WriteHeading(sb *strings.Builder, level int, text string) {
	sb.WriteString(Heading(level, text))
	sb.WriteString("\n")
}

// WriteBold writes bold text to the provided StringBuilder.
//
// Parameters:
//   - sb (*strings.Builder): The string builder to write to
//   - text (string): The text to format as bold
//
// Returns:
//   - None
//
// Algorithm:
//  1. Format text with the Bold function
//  2. Write the formatted text to the string builder
//
// Example:
//
//	var sb strings.Builder
//	WriteBold(&sb, "important")
//	// sb now contains "**important**"
//
// Notes:
//   - Helper function for writing bold text directly to a string builder
//   - Does not add a trailing newline
func WriteBold(sb *strings.Builder, text string) {
	sb.WriteString(Bold(text))
}

// WB is a shortcut for WriteBold.
//
// Parameters:
//   - sb (*strings.Builder): The string builder to write to
//   - text (string): The text to format as bold
//
// Returns:
//   - None
//
// Algorithm:
//  1. Call WriteBold with the provided string builder and text
//
// Example:
//
//	var sb strings.Builder
//	WB(&sb, "important")
//	// sb now contains "**important**"
//
// Notes:
//   - Convenience alias for the WriteBold function
//   - Provides a shorter name for frequently used function
func WB(sb *strings.Builder, text string) {
	WriteBold(sb, text)
}

// WriteItalic writes italic text to the provided StringBuilder.
//
// Parameters:
//   - sb (*strings.Builder): The string builder to write to
//   - text (string): The text to format as italic
//
// Returns:
//   - None
//
// Algorithm:
//  1. Format text with the Italic function
//  2. Write the formatted text to the string builder
//
// Example:
//
//	var sb strings.Builder
//	WriteItalic(&sb, "emphasized")
//	// sb now contains "*emphasized*"
//
// Notes:
//   - Helper function for writing italic text directly to a string builder
//   - Does not add a trailing newline
func WriteItalic(sb *strings.Builder, text string) {
	sb.WriteString(Italic(text))
}

// WI is a shortcut for WriteItalic.
//
// Parameters:
//   - sb (*strings.Builder): The string builder to write to
//   - text (string): The text to format as italic
//
// Returns:
//   - None
//
// Algorithm:
//  1. Call WriteItalic with the provided string builder and text
//
// Example:
//
//	var sb strings.Builder
//	WI(&sb, "emphasized")
//	// sb now contains "*emphasized*"
//
// Notes:
//   - Convenience alias for the WriteItalic function
//   - Provides a shorter name for frequently used function
func WI(sb *strings.Builder, text string) {
	WriteItalic(sb, text)
}

// WriteCode writes inline code to the provided StringBuilder.
//
// Parameters:
//   - sb (*strings.Builder): The string builder to write to
//   - text (string): The text to format as inline code
//
// Returns:
//   - None
//
// Algorithm:
//  1. Format text with the Code function
//  2. Write the formatted text to the string builder
//
// Example:
//
//	var sb strings.Builder
//	WriteCode(&sb, "var x = 10")
//	// sb now contains "`var x = 10`"
//
// Notes:
//   - Helper function for writing inline code directly to a string builder
//   - Does not add a trailing newline
func WriteCode(sb *strings.Builder, text string) {
	sb.WriteString(Code(text))
}

// WC is a shortcut for WriteCode.
//
// Parameters:
//   - sb (*strings.Builder): The string builder to write to
//   - text (string): The text to format as inline code
//
// Returns:
//   - None
//
// Algorithm:
//  1. Call WriteCode with the provided string builder and text
//
// Example:
//
//	var sb strings.Builder
//	WC(&sb, "var x = 10")
//	// sb now contains "`var x = 10`"
//
// Notes:
//   - Convenience alias for the WriteCode function
//   - Provides a shorter name for frequently used function
func WC(sb *strings.Builder, text string) {
	WriteCode(sb, text)
}

// WriteLink writes a link to the provided StringBuilder.
//
// Parameters:
//   - sb (*strings.Builder): The string builder to write to
//   - text (string): The display text for the link
//   - url (string): The URL the link points to
//
// Returns:
//   - None
//
// Algorithm:
//  1. Format the link with the Link function
//  2. Write the formatted link to the string builder
//
// Example:
//
//	var sb strings.Builder
//	WriteLink(&sb, "GitHub", "https://github.com")
//	// sb now contains "[GitHub](https://github.com)"
//
// Notes:
//   - Helper function for writing links directly to a string builder
//   - Does not add a trailing newline
func WriteLink(sb *strings.Builder, text, url string) {
	sb.WriteString(Link(text, url))
}

// WL is a shortcut for WriteLink.
//
// Parameters:
//   - sb (*strings.Builder): The string builder to write to
//   - text (string): The display text for the link
//   - url (string): The URL the link points to
//
// Returns:
//   - None
//
// Algorithm:
//  1. Call WriteLink with the provided string builder, text, and URL
//
// Example:
//
//	var sb strings.Builder
//	WL(&sb, "GitHub", "https://github.com")
//	// sb now contains "[GitHub](https://github.com)"
//
// Notes:
//   - Convenience alias for the WriteLink function
//   - Provides a shorter name for frequently used function
func WL(sb *strings.Builder, text, url string) {
	WriteLink(sb, text, url)
}

// WriteImage writes an image to the provided StringBuilder.
//
// Parameters:
//   - sb (*strings.Builder): The string builder to write to
//   - altText (string): The alternative text for the image
//   - url (string): The URL of the image
//
// Returns:
//   - None
//
// Algorithm:
//  1. Format the image with the Image function
//  2. Write the formatted image to the string builder
//
// Example:
//
//	var sb strings.Builder
//	WriteImage(&sb, "Logo", "https://example.com/logo.png")
//	// sb now contains "![Logo](https://example.com/logo.png)"
//
// Notes:
//   - Helper function for writing images directly to a string builder
//   - Does not add a trailing newline
func WriteImage(sb *strings.Builder, altText, url string) {
	sb.WriteString(Image(altText, url))
}

// WIMG is a shortcut for WriteImage.
//
// Parameters:
//   - sb (*strings.Builder): The string builder to write to
//   - altText (string): The alternative text for the image
//   - url (string): The URL of the image
//
// Returns:
//   - None
//
// Algorithm:
//  1. Call WriteImage with the provided string builder, alt text, and URL
//
// Example:
//
//	var sb strings.Builder
//	WIMG(&sb, "Logo", "https://example.com/logo.png")
//	// sb now contains "![Logo](https://example.com/logo.png)"
//
// Notes:
//   - Convenience alias for the WriteImage function
//   - Provides a shorter name for frequently used function
func WIMG(sb *strings.Builder, altText, url string) {
	WriteImage(sb, altText, url)
}

// WriteHorizontalRule writes a horizontal rule to the provided StringBuilder.
//
// Parameters:
//   - sb (*strings.Builder): The string builder to write to
//
// Returns:
//   - None
//
// Algorithm:
//  1. Write a horizontal rule using the HorizontalRule function
//  2. Add a newline after the horizontal rule
//
// Example:
//
//	var sb strings.Builder
//	WriteHorizontalRule(&sb)
//	// sb now contains "---\n"
//
// Notes:
//   - Helper function for writing horizontal rules directly to a string builder
//   - Automatically adds a trailing newline
func WriteHorizontalRule(sb *strings.Builder) {
	sb.WriteString(HorizontalRule())
	sb.WriteString("\n")
}

// WHR is a shortcut for WriteHorizontalRule.
//
// Parameters:
//   - sb (*strings.Builder): The string builder to write to
//
// Returns:
//   - None
//
// Algorithm:
//  1. Call WriteHorizontalRule with the provided string builder
//
// Example:
//
//	var sb strings.Builder
//	WHR(&sb)
//	// sb now contains "---\n"
//
// Notes:
//   - Convenience alias for the WriteHorizontalRule function
//   - Provides a shorter name for frequently used function
func WHR(sb *strings.Builder) {
	WriteHorizontalRule(sb)
}

// WriteQuoteBlock writes a block quote to the provided StringBuilder.
//
// Parameters:
//   - sb (*strings.Builder): The string builder to write to
//   - text (string): The text to format as a blockquote
//
// Returns:
//   - None
//
// Algorithm:
//  1. Format the text as a blockquote using the QuoteBlock function
//  2. Write the formatted blockquote to the string builder
//
// Example:
//
//	var sb strings.Builder
//	WriteQuoteBlock(&sb, "This is a quote")
//	// sb now contains "> This is a quote\n"
//
// Notes:
//   - Helper function for writing blockquotes directly to a string builder
//   - The trailing newline comes from the QuoteBlock function
func WriteQuoteBlock(sb *strings.Builder, text string) {
	sb.WriteString(QuoteBlock(text))
}

// WQB is a shortcut for WriteQuoteBlock.
//
// Parameters:
//   - sb (*strings.Builder): The string builder to write to
//   - text (string): The text to format as a blockquote
//
// Returns:
//   - None
//
// Algorithm:
//  1. Call WriteQuoteBlock with the provided string builder and text
//
// Example:
//
//	var sb strings.Builder
//	WQB(&sb, "This is a quote")
//	// sb now contains "> This is a quote\n"
//
// Notes:
//   - Convenience alias for the WriteQuoteBlock function
//   - Provides a shorter name for frequently used function
func WQB(sb *strings.Builder, text string) {
	WriteQuoteBlock(sb, text)
}

// ListItem creates a markdown list item with proper indentation.
//
// Parameters:
//   - text (string): The text content for the list item
//   - level (int): The indentation level (0 for top level, 1+ for nested levels)
//
// Returns:
//   - string: The formatted markdown list item
//
// Example:
//
//	item := ListItem("First item", 0)
//	// item will be "- First item"
//
//	nestedItem := ListItem("Nested item", 1)
//	// nestedItem will be "  - Nested item"
//
// Notes:
//   - Uses standard Markdown dash syntax for unordered lists
//   - Indentation uses 2 spaces per level for proper nesting
func ListItem(text string, level int) string {
	indent := strings.Repeat("  ", level)
	return fmt.Sprintf("%s- %s", indent, text)
}

// LI is a shortcut for calling WriteListItem.
//
// Parameters:
//   - sb (*strings.Builder): The string builder to write to
//   - text (string): The text content for the list item
//   - level (int): The indentation level (0 for top level, 1+ for nested levels)
//
// Returns:
//   - None
//
// Algorithm:
//  1. Call WriteListItem with the provided string builder, text, and level
//
// Example:
//
//	var sb strings.Builder
//	LI(&sb, "First item", 0)
//	// sb now contains "- First item\n"
//
// Notes:
//   - Convenience alias for the WriteListItem function
//   - Provides a shorter name for frequently used function
func LI(sb *strings.Builder, text string, level int) {
	WriteListItem(sb, text, level)
}

// WriteListItem writes a list item to the provided StringBuilder.
//
// Parameters:
//   - sb (*strings.Builder): The string builder to write to
//   - text (string): The text content for the list item
//   - level (int): The indentation level (0 for top level, 1+ for nested levels)
//
// Returns:
//   - None
//
// Algorithm:
//  1. Format the list item using the ListItem function
//  2. Write the formatted list item to the string builder
//  3. Add a newline after the list item
//
// Example:
//
//	var sb strings.Builder
//	WriteListItem(&sb, "First item", 0)
//	// sb now contains "- First item\n"
//
// Notes:
//   - Helper function for writing list items directly to a string builder
//   - Automatically adds a trailing newline
func WriteListItem(sb *strings.Builder, text string, level int) {
	sb.WriteString(ListItem(text, level))
	sb.WriteString("\n")
}

// WLI is a shortcut for WriteListItem.
//
// Parameters:
//   - sb (*strings.Builder): The string builder to write to
//   - text (string): The text content for the list item
//   - level (int): The indentation level (0 for top level, 1+ for nested levels)
//
// Returns:
//   - None
//
// Algorithm:
//  1. Call WriteListItem with the provided string builder, text, and level
//
// Example:
//
//	var sb strings.Builder
//	WLI(&sb, "First item", 0)
//	// sb now contains "- First item\n"
//
// Notes:
//   - Convenience alias for the WriteListItem function
//   - Provides a shorter name for frequently used function
func WLI(sb *strings.Builder, text string, level int) {
	WriteListItem(sb, text, level)
}

// ChecklistItem creates a markdown checklist item with proper indentation.
//
// Parameters:
//   - text (string): The text content for the checklist item
//   - checked (bool): Whether the checkbox should be checked or not
//   - level (int): The indentation level (0 for top level, 1+ for nested levels)
//
// Returns:
//   - string: The formatted markdown checklist item
//
// Algorithm:
//  1. Create indentation string based on level (2 spaces per level)
//  2. Determine the checkbox state based on the checked parameter
//  3. Format the string with proper indentation, checkbox state, and text
//
// Example:
//
//	item := ChecklistItem("Task to complete", false, 0)
//	// item will be "- [ ] Task to complete"
//
//	checkedItem := ChecklistItem("Completed task", true, 1)
//	// checkedItem will be "  - [x] Completed task"
//
// Notes:
//   - Uses standard Markdown task list syntax with brackets
//   - Unchecked items use "[ ]", checked items use "[x]"
//   - Indentation uses 2 spaces per level for proper nesting
//   - Compatible with most Markdown renderers that support task lists
func ChecklistItem(text string, checked bool, level int) string {
	indent := strings.Repeat("  ", level)
	checkmark := "[ ]"
	if checked {
		checkmark = "[x]"
	}
	return fmt.Sprintf("%s- %s %s", indent, checkmark, text)
}

func CLI(text string, checked bool, level int) string {
	return ChecklistItem(text, checked, level)
}

// WriteChecklistItem writes a checklist item to the provided StringBuilder.
func WriteChecklistItem(sb *strings.Builder, text string, checked bool, level int) {
	sb.WriteString(ChecklistItem(text, checked, level))
	sb.WriteString("\n")
}

func WCLI(sb *strings.Builder, text string, checked bool, level int) {
	WriteChecklistItem(sb, text, checked, level)
}

// KeyValue formats a key-value pair with the key in bold
func KeyValue(key string, value string) string {
	return fmt.Sprintf("%s: %s", Bold(key), value)
}

// KV is a shortcut for KeyValue.
//
// Parameters:
//   - key (string): The key to format
//   - value (string): The value to format
//
// Returns:
//   - string: The formatted key-value pair
//
// Example:
//
//	kv := KV("Name", "John Doe")
//	// kv will be "**Name**: John Doe"
//
// Notes:
//   - Convenience alias for the KeyValue function
//   - Provides a shorter name for frequently used function
func KV(key string, value string) string {
	return KeyValue(key, value)
}

// WriteKeyValue writes a key-value pair to the provided StringBuilder.
func WriteKeyValue(sb *strings.Builder, key string, value string) {
	sb.WriteString(KeyValue(key, value))
	sb.WriteString("\n")
}

func WKV(sb *strings.Builder, key string, value string) {
	WriteKeyValue(sb, key, value)
}
