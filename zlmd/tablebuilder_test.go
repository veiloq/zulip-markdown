package zlmd

import (
	"strings"
	"testing"
)

func TestTableBuilderIntegration(t *testing.T) {
	// Test the basic integration of TableBuilder with the main package
	tableBuilder := NewTableBuilder().
		WithHeaders("Name", "Age").
		AddRow("Alice", "30").
		AddRow("Bob", "25")

	result := tableBuilder.Build()

	// Verify basic table structure
	if !strings.Contains(result, "| Name | Age |") {
		t.Error("Table doesn't contain expected headers")
	}

	if !strings.Contains(result, "| Alice | 30 |") {
		t.Error("Table doesn't contain expected row data")
	}
}
