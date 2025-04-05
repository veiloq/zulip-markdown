package zlmd

import (
	"strings"
	"testing"
)

func TestSection_Basic(t *testing.T) {
	section := NewSection(2, "Test Section")
	result := section.Build()

	expected := "## Test Section\n\n\n"
	if result != expected {
		t.Errorf("Basic section formatting incorrect\nExpected:\n%q\nGot:\n%q", expected, result)
	}
}

func TestSection_WithText(t *testing.T) {
	section := NewSection(1, "Introduction").
		AddText("This is a paragraph.").
		AddText("This is another paragraph.")

	result := section.Build()

	if !strings.Contains(result, "# Introduction") {
		t.Errorf("Missing heading in result: %q", result)
	}

	if !strings.Contains(result, "This is a paragraph.") {
		t.Errorf("Missing first paragraph in result: %q", result)
	}

	if !strings.Contains(result, "This is another paragraph.") {
		t.Errorf("Missing second paragraph in result: %q", result)
	}
}

func TestSection_WithList(t *testing.T) {
	section := NewSection(3, "Features").
		AddBullet("Easy to use").
		AddBullet("Fast").
		AddBullet("Flexible")

	result := section.Build()

	expected := "### Features\n\n* Easy to use\n* Fast\n* Flexible\n\n"
	if result != expected {
		t.Errorf("Bullet list formatting incorrect\nExpected:\n%q\nGot:\n%q", expected, result)
	}
}

func TestSection_WithNumberedList(t *testing.T) {
	section := NewSection(2, "Steps").
		AddNumberedItem(1, "Download").
		AddNumberedItem(2, "Install").
		AddNumberedItem(3, "Run")

	result := section.Build()

	expected := "## Steps\n\n1. Download\n2. Install\n3. Run\n\n"
	if result != expected {
		t.Errorf("Numbered list formatting incorrect\nExpected:\n%q\nGot:\n%q", expected, result)
	}
}

func TestSection_WithTable(t *testing.T) {
	section := NewSection(2, "Users")

	table := NewTableBuilder().
		WithHeaders("Name", "Age").
		AddRow("Alice", "30").
		AddRow("Bob", "25")

	section.AddTable(table)

	result := section.Build()

	if !strings.Contains(result, "## Users") {
		t.Errorf("Missing heading in result: %q", result)
	}

	if !strings.Contains(result, "| Name | Age |") {
		t.Errorf("Missing table headers in result: %q", result)
	}

	if !strings.Contains(result, "| Alice | 30 |") {
		t.Errorf("Missing table row in result: %q", result)
	}

	if !strings.Contains(result, "| Bob | 25 |") {
		t.Errorf("Missing table row in result: %q", result)
	}
}

func TestSection_HeadingLevelLimits(t *testing.T) {
	// Test minimum level
	section1 := NewSection(0, "Min Level")
	if section1.Level != 1 {
		t.Errorf("Expected minimum heading level to be 1, got %d", section1.Level)
	}

	// Test maximum level
	section2 := NewSection(7, "Max Level")
	if section2.Level != 6 {
		t.Errorf("Expected maximum heading level to be 6, got %d", section2.Level)
	}
}

func TestSection_Chaining(t *testing.T) {
	section := NewSection(1, "Document")
	result := section.
		AddText("Introduction paragraph.").
		AddBullet("First point").
		AddBullet("Second point").
		AddNumberedItem(1, "Step one").
		AddNumberedItem(2, "Step two").
		Build()

	// Check that chaining works correctly
	if !strings.Contains(result, "Introduction paragraph.") ||
		!strings.Contains(result, "* First point") ||
		!strings.Contains(result, "2. Step two") {
		t.Errorf("Method chaining failed, missing expected content in:\n%s", result)
	}
}
