package zlmd

import (
	"strings"
)

// Alignment represents the text alignment in a table column.
type Alignment string

const (
	// AlignDefault represents default alignment (typically left)
	AlignDefault Alignment = ""
	// AlignLeft represents left alignment
	AlignLeft Alignment = "left"
	// AlignCenter represents center alignment
	AlignCenter Alignment = "center"
	// AlignRight represents right alignment
	AlignRight Alignment = "right"
)

// TableBuilder represents a markdown table builder.
type TableBuilder struct {
	headers       []string
	rows          [][]string
	alignments    []Alignment
	headerBuilder func(string) string
}

// NewTableBuilder creates a new markdown table builder.
//
// Returns:
//   - *TableBuilder: A new initialized TableBuilder instance with empty headers, rows, and default alignment
//
// Example:
//
//	table := NewTableBuilder()
//	// Creates an empty table builder
func NewTableBuilder() *TableBuilder {
	return &TableBuilder{
		headers:    []string{},
		rows:       [][]string{},
		alignments: []Alignment{},
		headerBuilder: func(s string) string {
			return s
		},
	}
}

// WithHeaders adds headers to the table.
//
// Parameters:
//   - headers (...string): Variable number of header strings to add
//
// Returns:
//   - *TableBuilder: The same TableBuilder instance (for method chaining)
//
// Example:
//
//	table.WithHeaders("Name", "Age", "Email")
//	// Adds three headers to the table
func (t *TableBuilder) WithHeaders(headers ...string) *TableBuilder {
	t.headers = append(t.headers, headers...)
	// Initialize alignments for new headers
	for i := len(t.alignments); i < len(t.headers); i++ {
		t.alignments = append(t.alignments, AlignDefault)
	}
	return t
}

// WithHeaderStyle sets a function to style header cells.
//
// Parameters:
//   - styleFunc (func(string) string): A function that takes a header string and returns a styled version
//
// Returns:
//   - *TableBuilder: The same TableBuilder instance (for method chaining)
//
// Example:
//
//	table.WithHeaderStyle(func(s string) string {
//	  return "*" + s + "*" // Makes headers italic
//	})
func (t *TableBuilder) WithHeaderStyle(styleFunc func(string) string) *TableBuilder {
	t.headerBuilder = styleFunc
	return t
}

// WithBoldHeaders applies bold formatting to headers.
//
// Returns:
//   - *TableBuilder: The same TableBuilder instance (for method chaining)
//
// Example:
//
//	table.WithBoldHeaders()
//	// Makes all headers bold with ** markers
func (t *TableBuilder) WithBoldHeaders() *TableBuilder {
	return t.WithHeaderStyle(func(s string) string {
		return "**" + s + "**"
	})
}

// AddRow adds a row to the table.
//
// Parameters:
//   - cells (...string): Variable number of cell values to add as a row
//
// Returns:
//   - *TableBuilder: The same TableBuilder instance (for method chaining)
//
// Example:
//
//	table.AddRow("John Doe", "30", "john@example.com")
//	// Adds a row with three cells
func (t *TableBuilder) AddRow(cells ...string) *TableBuilder {
	// Make a copy of the cells to avoid external modification
	row := make([]string, len(cells))
	copy(row, cells)
	t.rows = append(t.rows, row)
	return t
}

// AddRows adds multiple rows to the table.
//
// Parameters:
//   - rows ([][]string): A slice of rows, where each row is a slice of cell values
//
// Returns:
//   - *TableBuilder: The same TableBuilder instance (for method chaining)
//
// Example:
//
//	data := [][]string{
//	  {"John", "30", "john@example.com"},
//	  {"Jane", "25", "jane@example.com"},
//	}
//	table.AddRows(data)
//	// Adds two rows to the table
func (t *TableBuilder) AddRows(rows [][]string) *TableBuilder {
	for _, row := range rows {
		t.AddRow(row...)
	}
	return t
}

// SetAlignment sets the alignment for a specific column.
//
// Parameters:
//   - column (int): Zero-based index of the column to set alignment for
//   - alignment (Alignment): The alignment type (AlignDefault, AlignLeft, AlignCenter, or AlignRight)
//
// Returns:
//   - *TableBuilder: The same TableBuilder instance (for method chaining)
//
// Example:
//
//	table.SetAlignment(1, AlignCenter)
//	// Sets the second column to be center-aligned
func (t *TableBuilder) SetAlignment(column int, alignment Alignment) *TableBuilder {
	if column >= 0 && column < len(t.alignments) {
		t.alignments[column] = alignment
	}
	return t
}

// SetAlignments sets the alignment for all columns.
//
// Parameters:
//   - alignments (...Alignment): Variable number of alignment values to apply to columns
//
// Returns:
//   - *TableBuilder: The same TableBuilder instance (for method chaining)
//
// Example:
//
//	table.SetAlignments(AlignLeft, AlignCenter, AlignRight)
//	// Sets the first column left-aligned, second centered, and third right-aligned
func (t *TableBuilder) SetAlignments(alignments ...Alignment) *TableBuilder {
	for i, alignment := range alignments {
		if i < len(t.alignments) {
			t.alignments[i] = alignment
		}
	}
	return t
}

// Build generates the markdown table string.
//
// Returns:
//   - string: The formatted markdown table as a string
//
// Example:
//
//	tableStr := table.Build()
//	// Generates a formatted markdown table
func (t *TableBuilder) Build() string {
	if len(t.headers) == 0 {
		return ""
	}

	var sb strings.Builder

	// Write header row
	sb.WriteString("| ")
	for i, header := range t.headers {
		if i > 0 {
			sb.WriteString(" | ")
		}
		sb.WriteString(t.headerBuilder(header))
	}
	sb.WriteString(" |\n")

	// Write separator row with alignment markers
	sb.WriteString("| ")
	for i, alignment := range t.alignments {
		if i > 0 {
			sb.WriteString(" | ")
		}

		switch alignment {
		case AlignLeft:
			sb.WriteString(":---")
		case AlignCenter:
			sb.WriteString(":---:")
		case AlignRight:
			sb.WriteString("---:")
		default:
			sb.WriteString("---")
		}
	}
	sb.WriteString(" |\n")

	// Write data rows
	for _, row := range t.rows {
		sb.WriteString("| ")
		for i, cell := range row {
			if i > 0 {
				sb.WriteString(" | ")
			}
			if i < len(t.headers) {
				sb.WriteString(cell)
			}
		}

		// Add empty cells if row has fewer cells than headers
		for i := len(row); i < len(t.headers); i++ {
			sb.WriteString(" | ")
		}

		sb.WriteString(" |\n")
	}

	return sb.String()
}
