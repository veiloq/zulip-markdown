package zlmd

import (
	"strings"
	"testing"
)

func TestSpoiler(t *testing.T) {
	tests := []struct {
		name     string
		heading  string
		content  string
		expected string
	}{
		{
			name:     "Basic spoiler",
			heading:  "Warning",
			content:  "This is hidden content",
			expected: "```spoiler Warning\nThis is hidden content\n```",
		},
		{
			name:     "Spoiler with multiline content",
			heading:  "Details",
			content:  "Line 1\nLine 2\nLine 3",
			expected: "```spoiler Details\nLine 1\nLine 2\nLine 3\n```",
		},
		{
			name:     "Spoiler with nested code block",
			heading:  "Code Example",
			content:  "```go\nfmt.Println(\"Hello\")\n```",
			expected: "```spoiler Code Example\n~~~go\nfmt.Println(\"Hello\")\n~~~\n```",
		},
		{
			name:     "Empty spoiler",
			heading:  "Empty",
			content:  "",
			expected: "```spoiler Empty\n\n```",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Spoiler(tt.heading, tt.content)
			if got != tt.expected {
				t.Errorf("Spoiler() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestSpoilerEscaped(t *testing.T) {
	tests := []struct {
		name     string
		heading  string
		content  string
		fence    string
		expected string
	}{
		{
			name:     "With backticks fence",
			heading:  "Notice",
			content:  "Important information",
			fence:    "```",
			expected: "```spoiler Notice\nImportant information\n```",
		},
		{
			name:     "With tildes fence",
			heading:  "Warning",
			content:  "Be careful",
			fence:    "~~~",
			expected: "~~~spoiler Warning\nBe careful\n~~~",
		},
		{
			name:     "With custom fence",
			heading:  "Custom",
			content:  "Custom fence",
			fence:    "+++",
			expected: "+++spoiler Custom\nCustom fence\n+++",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SpoilerEscaped(tt.heading, tt.content, tt.fence)
			if got != tt.expected {
				t.Errorf("SpoilerEscaped() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestSpoilerEscapedTilde(t *testing.T) {
	got := SpoilerEscapedTilde("Title", "Content")
	expected := "~~~spoiler Title\nContent\n~~~"

	if got != expected {
		t.Errorf("SpoilerEscapedTilde() = %q, want %q", got, expected)
	}
}

func TestSpoilerEscapedFence(t *testing.T) {
	got := SpoilerEscapedFence("Title", "Content")
	expected := "```spoiler Title\nContent\n```"

	if got != expected {
		t.Errorf("SpoilerEscapedFence() = %q, want %q", got, expected)
	}
}

func TestCodeBlock(t *testing.T) {
	tests := []struct {
		name     string
		language string
		content  string
		expected string
	}{
		{
			name:     "Go code",
			language: "go",
			content:  "fmt.Println(\"Hello, World!\")",
			expected: "```go\nfmt.Println(\"Hello, World!\")\n```",
		},
		{
			name:     "JavaScript code",
			language: "javascript",
			content:  "console.log('Hello, World!');",
			expected: "```javascript\nconsole.log('Hello, World!');\n```",
		},
		{
			name:     "No language specified",
			language: "",
			content:  "Some code",
			expected: "```\nSome code\n```",
		},
		{
			name:     "Multiline code",
			language: "python",
			content:  "def hello():\n    print('Hello, World!')",
			expected: "```python\ndef hello():\n    print('Hello, World!')\n```",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CodeBlock(tt.language, tt.content)
			if got != tt.expected {
				t.Errorf("CodeBlock() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestMarkdownBlock(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected string
	}{
		{
			name:     "Basic markdown",
			content:  "# Heading\n\nParagraph",
			expected: "```markdown\n# Heading\n\nParagraph\n```",
		},
		{
			name:     "Markdown with formatting",
			content:  "**Bold** and *Italic*",
			expected: "```markdown\n**Bold** and *Italic*\n```",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MarkdownBlock(tt.content)
			if got != tt.expected {
				t.Errorf("MarkdownBlock() = %q, want %q", got, tt.expected)
			}
		})
	}
}

// Test the integration of multiple blocks
func TestNestedBlocks(t *testing.T) {
	// Create a spoiler that contains a code block
	spoilerWithCode := Spoiler("Code Example", CodeBlock("go", "fmt.Println(\"Hello\")"))

	// The inner code block should have tildes, not backticks
	if !strings.Contains(spoilerWithCode, "~~~go") {
		t.Errorf("Inner code block fence not converted to tildes: %s", spoilerWithCode)
	}

	// The code block inside should be properly formatted
	expected := "```spoiler Code Example\n~~~go\nfmt.Println(\"Hello\")\n~~~\n```"
	if spoilerWithCode != expected {
		t.Errorf("Nested blocks incorrect\nGot: %q\nWant: %q", spoilerWithCode, expected)
	}
}
