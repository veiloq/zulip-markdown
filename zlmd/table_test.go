package zlmd

import (
	"strings"
	"testing"
)

func TestTableBuilder_Empty(t *testing.T) {
	table := NewTableBuilder()
	result := table.Build()

	if result != "" {
		t.Errorf("Expected empty table to produce empty string, got %q", result)
	}
}

func TestTableBuilder_WithHeaders(t *testing.T) {
	table := NewTableBuilder().
		WithHeaders("Name", "Age", "City")

	result := table.Build()
	expected := "| Name | Age | City |\n| --- | --- | --- |\n"

	if result != expected {
		t.Errorf("Header formatting incorrect\nExpected:\n%q\nGot:\n%q", expected, result)
	}
}

func TestTableBuilder_WithBoldHeaders(t *testing.T) {
	table := NewTableBuilder().
		WithHeaders("Name", "Age", "City").
		WithBoldHeaders()

	result := table.Build()
	expected := "| **Name** | **Age** | **City** |\n| --- | --- | --- |\n"

	if result != expected {
		t.Errorf("Bold header formatting incorrect\nExpected:\n%q\nGot:\n%q", expected, result)
	}
}

func TestTableBuilder_AddRow(t *testing.T) {
	table := NewTableBuilder().
		WithHeaders("Name", "Age", "City").
		AddRow("Alice", "30", "New York").
		AddRow("Bob", "25", "San Francisco")

	result := table.Build()
	expected := "| Name | Age | City |\n| --- | --- | --- |\n| Alice | 30 | New York |\n| Bob | 25 | San Francisco |\n"

	if result != expected {
		t.Errorf("Row formatting incorrect\nExpected:\n%q\nGot:\n%q", expected, result)
	}
}

func TestTableBuilder_AddRows(t *testing.T) {
	rows := [][]string{
		{"Alice", "30", "New York"},
		{"Bob", "25", "San Francisco"},
	}

	table := NewTableBuilder().
		WithHeaders("Name", "Age", "City").
		AddRows(rows)

	result := table.Build()
	expected := "| Name | Age | City |\n| --- | --- | --- |\n| Alice | 30 | New York |\n| Bob | 25 | San Francisco |\n"

	if result != expected {
		t.Errorf("AddRows formatting incorrect\nExpected:\n%q\nGot:\n%q", expected, result)
	}
}

func TestTableBuilder_Alignments(t *testing.T) {
	table := NewTableBuilder().
		WithHeaders("Left", "Center", "Right").
		SetAlignment(0, AlignLeft).
		SetAlignment(1, AlignCenter).
		SetAlignment(2, AlignRight).
		AddRow("1", "2", "3")

	result := table.Build()
	if !strings.Contains(result, ":---") {
		t.Errorf("Left alignment missing, got: %q", result)
	}
	if !strings.Contains(result, ":---:") {
		t.Errorf("Center alignment missing, got: %q", result)
	}
	if !strings.Contains(result, "---:") {
		t.Errorf("Right alignment missing, got: %q", result)
	}
}

func TestTableBuilder_SetAlignments(t *testing.T) {
	table := NewTableBuilder().
		WithHeaders("Left", "Center", "Right").
		SetAlignments(AlignLeft, AlignCenter, AlignRight).
		AddRow("1", "2", "3")

	result := table.Build()
	expected := "| Left | Center | Right |\n| :--- | :---: | ---: |\n| 1 | 2 | 3 |\n"

	if result != expected {
		t.Errorf("Alignments incorrect\nExpected:\n%q\nGot:\n%q", expected, result)
	}
}

func TestTableBuilder_FewerCellsThanHeaders(t *testing.T) {
	table := NewTableBuilder().
		WithHeaders("Name", "Age", "City").
		AddRow("Alice", "30") // Missing city

	result := table.Build()
	expected := "| Name | Age | City |\n| --- | --- | --- |\n| Alice | 30 |  |\n"

	if result != expected {
		t.Errorf("Missing cell handling incorrect\nExpected:\n%q\nGot:\n%q", expected, result)
	}
}

func TestTableBuilder_MoreCellsThanHeaders(t *testing.T) {
	table := NewTableBuilder().
		WithHeaders("Name", "Age").
		AddRow("Alice", "30", "Extra") // Extra cell beyond headers

	result := table.Build()

	if !strings.Contains(result, "| Alice | 30 |") {
		t.Errorf("Extra cell handling incorrect\nExpected to contain:\n%q\nGot:\n%q", "| Alice | 30 |", result)
	}
}

func TestTableBuilder_CustomHeaderStyle(t *testing.T) {
	table := NewTableBuilder().
		WithHeaders("Name", "Age", "City").
		WithHeaderStyle(func(s string) string {
			return "*" + s + "*" // Italic headers
		}).
		AddRow("Alice", "30", "New York")

	result := table.Build()
	expected := "| *Name* | *Age* | *City* |\n| --- | --- | --- |\n| Alice | 30 | New York |\n"

	if result != expected {
		t.Errorf("Custom header style incorrect\nExpected:\n%q\nGot:\n%q", expected, result)
	}
}

func TestTableBuilder_Chaining(t *testing.T) {
	// Test that all methods can be chained together
	table := NewTableBuilder().
		WithHeaders("Name", "Age", "City").
		WithBoldHeaders().
		SetAlignments(AlignLeft, AlignCenter, AlignRight).
		AddRow("Alice", "30", "New York").
		AddRow("Bob", "25", "San Francisco")

	result := table.Build()

	if result == "" {
		t.Error("Method chaining failed to produce output")
	}

	// Verify the result has the expected structure
	if !strings.Contains(result, "**Name**") ||
		!strings.Contains(result, ":---") ||
		!strings.Contains(result, ":---:") ||
		!strings.Contains(result, "---:") ||
		!strings.Contains(result, "Alice") {
		t.Errorf("Method chaining produced incorrect output: %q", result)
	}
}
