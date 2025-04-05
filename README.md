# ZLMD - Zulip Markdown Library

A production-ready Go library for working with Zulip-flavored Markdown, providing utilities for formatting text, creating tables, and handling special Zulip-specific syntax.

## Installation

Requires Go 1.24 or later:

```bash
go get github.com/veiloq/zlmd
```

## Quick Example

```go
package main

import (
	"fmt"
	"time"

	"github.com/veiloq/zlmd"
)

func main() {
	// Basic formatting
	bold := zlmd.Bold("Important text")
	italic := zlmd.Italic("Emphasized text")
	fmt.Printf("Formatting: %s and %s\n", bold, italic)

	// Create a table
	table := zlmd.NewTable()
	table.AddHeader("Name", "Role", "Team")
	table.AddRow("Alice", "Engineer", "Backend")
	table.AddRow("Bob", "Designer", "UI/UX")
	
	fmt.Println("\nTable Example:")
	fmt.Println(table.String())
	
	// Format time for Zulip's time tags
	t := time.Date(2023, 5, 15, 14, 30, 0, 0, time.UTC)
	timeTag := zlmd.ZLFormatTime(t) 
	fmt.Printf("\nFormatted time: %s\n", timeTag)
}
```

## Features

### Basic Markdown Formatting

```go
// Text styling
bold := zlmd.Bold("Important")       // **Important**
italic := zlmd.Italic("Emphasized")  // *Emphasized*
strike := zlmd.Strike("Removed")     // ~~Removed~~
code := zlmd.Code("func()")          // `func()`

// Links
link := zlmd.Link("Zulip", "https://zulip.com")  // [Zulip](https://zulip.com)
```

### Code Blocks

```go
// Simple code block
codeBlock := zlmd.CodeBlock("console.log('Hello');", "javascript")

/* Output:
```javascript
console.log('Hello');
```
*/

// Code block with filename
namedCodeBlock := zlmd.CodeBlockWithFilename(
    "package main\n\nfunc main() {\n\tfmt.Println(\"Hello\")\n}",
    "go",
    "hello.go"
)

/* Output:
```go:hello.go
package main

func main() {
	fmt.Println("Hello")
}
```
*/
```

### Lists

```go
// Unordered list
list := zlmd.NewList()
list.AddItem("First item")
list.AddItem("Second item")
    .AddSubItem("Nested item 1")
    .AddSubItem("Nested item 2")
list.AddItem("Third item")

/* Output:
* First item
* Second item
  * Nested item 1
  * Nested item 2
* Third item
*/

// Ordered list
orderedList := zlmd.NewOrderedList()
orderedList.AddItem("First step")
orderedList.AddItem("Second step")
orderedList.AddItem("Third step")

/* Output:
1. First step
2. Second step
3. Third step
*/
```

### Tables

```go
// Basic table
table := zlmd.NewTable()
table.AddHeader("Item", "Price", "Quantity")
table.AddRow("Widget", "$10.00", "5")
table.AddRow("Gadget", "$25.50", "2")
fmt.Println(table.String())

/* Output:
| Item | Price | Quantity |
|------|-------|----------|
| Widget | $10.00 | 5 |
| Gadget | $25.50 | 2 |
*/

// Table with alignment
alignedTable := zlmd.NewTable()
alignedTable.AddHeader("Left", "Center", "Right")
alignedTable.SetAlignment(0, zlmd.AlignLeft)
alignedTable.SetAlignment(1, zlmd.AlignCenter)
alignedTable.SetAlignment(2, zlmd.AlignRight)
alignedTable.AddRow("Text", "Text", "Text")

/* Output:
| Left | Center | Right |
|:-----|:------:|------:|
| Text | Text | Text |
*/
```

### Zulip-Specific Features

```go
// Time formatting
timeTag := zlmd.ZLFormatTime(time.Now())  
// <time:2023-05-15T14:30:00Z>

// Emoji shortcuts
emoji := zlmd.Emoji("smile")  // :smile:
thumbsUp := zlmd.Emoji("thumbs_up")  // :thumbs_up:

// User mentions
mention := zlmd.UserMention("username")  // @**username**
groupMention := zlmd.GroupMention("developers")  // @*developers*

// Stream links
streamLink := zlmd.StreamLink("general")  // #**general**
topicLink := zlmd.TopicLink("general", "welcome")  // #**general>welcome**

// Message Quotes
quote := zlmd.Quote("This is a quoted message")
/* Output:
> This is a quoted message
*/

// Spoilers
spoiler := zlmd.Spoiler("Plot twist: Bruce Willis was dead the whole time")
/* Output:
<spoiler>Plot twist: Bruce Willis was dead the whole time</spoiler>
*/

// Math expressions
inlineMath := zlmd.InlineMath("E=mc^2")  // $$E=mc^2$$
displayMath := zlmd.DisplayMath(`\int_{a}^{b} f(x) \, dx`)
/* Output:
$$\int_{a}^{b} f(x) \, dx$$
*/
```

### Markdown Document Building

```go
// Create a document
doc := zlmd.NewDocument()

// Add a heading
doc.AddHeading("Project Overview", 1)

// Add text paragraphs
doc.AddParagraph("This is an introduction to our project.")

// Add a code block
doc.AddCodeBlock(`console.log("Hello World");`, "javascript")

// Add a table
table := zlmd.NewTable()
table.AddHeader("Feature", "Status")
table.AddRow("Authentication", "Complete")
table.AddRow("API", "In Progress")
doc.AddTable(table)

// Get the full document
markdownText := doc.String()
```

## Error Handling

The library provides standardized error types for consistent error handling:

```go
table, err := zlmd.ParseTable(markdownText)
if err != nil {
	switch {
	case errors.Is(err, zlmd.ErrInvalidTableFormat):
		log.Fatal("Invalid table format")
	case errors.Is(err, zlmd.ErrEmptyTable):
		log.Fatal("Table is empty")
	case errors.Is(err, zlmd.ErrMalformedMarkdown):
		log.Fatal("Malformed markdown syntax")
	default:
		log.Fatalf("Failed to parse table: %v", err)
	}
}
```

Common error types include:

| Error Constant | Description |
|----------------|-------------|
| `ErrInvalidTableFormat` | Table structure is malformed |
| `ErrEmptyTable` | Table has no rows or headers |
| `ErrMalformedMarkdown` | Invalid markdown syntax |
| `ErrInvalidHeadingLevel` | Heading level outside valid range (1-6) |
| `ErrInvalidListNesting` | List nesting is too deep or malformed |

## Advanced Usage

### Custom Markdown Extensions

```go
// Create a custom Zulip-compatible markdown extension
customExtension := zlmd.NewExtension("poll")
pollOptions := []string{"Option A", "Option B", "Option C"}
pollMarkdown := customExtension.Format(strings.Join(pollOptions, "\n"))

/* Output:
<poll>
Option A
Option B
Option C
</poll>
*/
```

### Markdown Parser

```go
// Parse existing markdown text
parsed, err := zlmd.Parse("**Bold** and *italic* text")
if err != nil {
    log.Fatalf("Failed to parse markdown: %v", err)
}

// Access parsed elements
for _, element := range parsed.Elements {
    switch e := element.(type) {
    case *zlmd.TextElement:
        fmt.Printf("Text: %s\n", e.Content)
    case *zlmd.EmphasisElement:
        fmt.Printf("Emphasis: %s (Level: %d)\n", e.Content, e.Level)
    // More element types...
    }
}
```

## Development

For contributing to the project:

1. Clone the repository:
   ```bash
   git clone https://github.com/veiloq/zlmd.git
   ```

2. Run tests:
   ```bash
   make test
   ```

3. Format code:
   ```bash
   make fmt
   ```

4. Run linter:
   ```bash
   make lint
   ```

## License

This project is licensed under the MIT License - see the LICENSE file for details. 