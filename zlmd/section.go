package zlmd

import (
	"strconv"
	"strings"
)

// Section represents a markdown document section with a title and content.
type Section struct {
	Level   int
	Title   string
	Content []string
}

// NewSection creates a new markdown section with the specified heading level and title.
//
// Parameters:
//   - level (int): The heading level (1-6)
//   - title (string): The section title
//
// Returns:
//   - *Section: A new initialized Section instance
//
// Example:
//
//	section := NewSection(2, "Introduction")
//	// Creates a section with ## heading
func NewSection(level int, title string) *Section {
	// Limit level to valid markdown heading levels (1-6)
	if level < 1 {
		level = 1
	} else if level > 6 {
		level = 6
	}

	return &Section{
		Level:   level,
		Title:   title,
		Content: []string{},
	}
}

// AddBullet adds a bullet point to the section.
//
// Parameters:
//   - text (string): The bullet point text
//
// Returns:
//   - *Section: The same Section instance (for method chaining)
//
// Example:
//
//	section.AddBullet("This is a bullet point")
//	// Adds "* This is a bullet point" to the section content
func (s *Section) AddBullet(text string) *Section {
	s.Content = append(s.Content, "* "+text)
	return s
}

// AddNumberedItem adds a numbered item to the section.
//
// Parameters:
//   - number (int): The item number
//   - text (string): The item text
//
// Returns:
//   - *Section: The same Section instance (for method chaining)
//
// Example:
//
//	section.AddNumberedItem(1, "First step")
//	// Adds "1. First step" to the section content
func (s *Section) AddNumberedItem(number int, text string) *Section {
	s.Content = append(s.Content, strconv.Itoa(number)+". "+text)
	return s
}

// AddText adds a text paragraph to the section.
//
// Parameters:
//   - text (string): The text to add
//
// Returns:
//   - *Section: The same Section instance (for method chaining)
//
// Example:
//
//	section.AddText("This is a paragraph of explanatory text.")
//	// Adds the text to the section content
func (s *Section) AddText(text string) *Section {
	s.Content = append(s.Content, text)
	return s
}

// AddTable adds a markdown table to the section.
//
// Parameters:
//   - table (*TableBuilder): The table builder to add
//
// Returns:
//   - *Section: The same Section instance (for method chaining)
//
// Example:
//
//	table := NewTableBuilder().WithHeaders("Name", "Age").AddRow("John", "30")
//	section.AddTable(table)
//	// Adds the built table to the section content
func (s *Section) AddTable(table *TableBuilder) *Section {
	s.Content = append(s.Content, table.Build())
	return s
}

// Build generates the markdown section string.
//
// Returns:
//   - string: The complete formatted markdown section as a string
//
// Example:
//
//	sectionStr := section.Build()
//	// Generates a formatted markdown section with heading and content
func (s *Section) Build() string {
	var sb strings.Builder

	// Add heading
	for i := 0; i < s.Level; i++ {
		sb.WriteString("#")
	}
	sb.WriteString(" ")
	sb.WriteString(s.Title)
	sb.WriteString("\n\n")

	// Add content
	for _, item := range s.Content {
		sb.WriteString(item)
		sb.WriteString("\n")
	}

	// Add final newline
	sb.WriteString("\n")

	return sb.String()
}
