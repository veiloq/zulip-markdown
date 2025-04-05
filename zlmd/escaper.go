package zlmd

import (
	"strings"
)

// EscapeMarkdown processes a Markdown string to escape nested code fences
// within spoiler blocks (` ```spoiler ... ``` `).
// It replaces the inner ` ``` ` fences with ` ~~~ `.
//
// Parameters:
//   - markdown: The input Markdown string.
//
// Returns:
//   - string: The processed string with nested fences escaped, or the original
//     string if the input structure is invalid (e.g., mismatched fences).
//   - bool: True if the processing was successful (input was valid or required no changes),
//     False if the structure was invalid (e.g., unbalanced fences).
func _EscapeMarkdown(markdown string) (string, bool) {
	var result strings.Builder
	var spoilerContent strings.Builder // Temporarily holds content within a spoiler block
	lines := strings.Split(markdown, "\n")
	inSpoiler := false
	inTopLevelCode := false // Tracks if currently inside a ``` block outside a spoiler

	for i, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		if inSpoiler {
			// We are inside a spoiler block, keep collecting lines
			spoilerContent.WriteString(line)
			// Add newline unless it's the very last line of the input
			if i < len(lines)-1 {
				spoilerContent.WriteByte('\n')
			}

			// Check if this line closes the spoiler block
			if trimmedLine == "```" {
				inSpoiler = false
				// Process the collected spoiler content
				processedSpoiler, ok := processSpoilerContent(spoilerContent.String())
				if !ok {
					// Invalid nesting found within the spoiler block
					return markdown, false
				}
				result.WriteString(processedSpoiler)
				// processSpoilerContent should handle its own trailing newline if needed
				// Based on tests, the output might not need an extra newline here if processSpoilerContent includes the final ```
				// Let's remove the automatic newline addition after processing spoiler
			}
		} else if inTopLevelCode {
			// Inside a regular code block (outside spoiler)
			result.WriteString(line)
			if i < len(lines)-1 {
				result.WriteByte('\n')
			}
			if trimmedLine == "```" {
				inTopLevelCode = false // Closing the top-level code block
			}
		} else {
			// Not in spoiler, not in top-level code block
			if strings.HasPrefix(trimmedLine, "```spoiler") {
				// Start of a spoiler block
				inSpoiler = true
				spoilerContent.Reset()
				spoilerContent.WriteString(line)
				// Add newline unless it's the very last line of the input
				if i < len(lines)-1 {
					spoilerContent.WriteByte('\n')
				}
			} else if strings.HasPrefix(trimmedLine, "```") {
				// Start of a top-level code block
				inTopLevelCode = true
				result.WriteString(line)
				if i < len(lines)-1 {
					result.WriteByte('\n')
				}
				// Handle immediate close like ``` ``` case - though unlikely markdown
				if trimmedLine == "```" && len(trimmedLine) == 3 {
					// This assumes ``` alone starts AND immediately ends.
					// If ``` has lang specifier, inTopLevelCode remains true.
					// If it's just ```, it should close on next ``` line.
					// Let's refine: Only check for closing ``` when already inTopLevelCode
					// So, this immediate close check might be wrong/unnecessary.
					// Let's stick to the state logic: It starts here, closes later.
				}

			} else {
				// Plain text line
				result.WriteString(line)
				// Add newline unless it's the very last line of the input
				if i < len(lines)-1 {
					result.WriteByte('\n')
				}
			}
		}
	}

	// After processing all lines, check for unclosed blocks
	if inSpoiler || inTopLevelCode {
		return markdown, false // Unclosed spoiler or top-level code block
	}

	// Handle potential trailing newline discrepancies if needed, though the line-by-line
	// approach with checks for the last line aims to prevent this.
	finalResult := result.String()
	// Example check (might need refinement based on desired exact behavior):
	// if !strings.HasSuffix(markdown, "\n") && strings.HasSuffix(finalResult, "\n") {
	// 	finalResult = strings.TrimSuffix(finalResult, "\n")
	// }

	return finalResult, true
}

// processSpoilerContent handles the transformation of nested fences within a spoiler block's content.
// Input spoilerBlock includes the opening ```spoiler... and closing ``` lines.
func processSpoilerContent(spoilerBlock string) (string, bool) {
	var result strings.Builder
	// Use TrimSpace to handle potential empty spoilerBlock input gracefully,
	// although EscapeMarkdown should generally not pass empty content here.
	trimmedBlock := strings.TrimSpace(spoilerBlock)
	if trimmedBlock == "" {
		return "", true // Or false if empty spoiler is invalid
	}
	lines := strings.Split(trimmedBlock, "\n") // Split the content part
	numLines := len(lines)

	if numLines == 0 || !strings.HasPrefix(lines[0], "```spoiler") {
		return spoilerBlock, false // Should start with spoiler fence
	}
	if numLines < 2 || lines[numLines-1] != "```" {
		return spoilerBlock, false // Should have at least 2 lines and end with ```
	}

	result.WriteString(lines[0]) // Write the opening ```spoiler line
	result.WriteByte('\n')

	inNestedCode := false
	fenceBalance := 0 // 0: outside nested, 1: inside nested

	// Process lines between the spoiler start and end fences
	for i := 1; i < numLines-1; i++ {
		line := lines[i]
		trimmedLine := strings.TrimSpace(line)

		if !inNestedCode {
			if strings.HasPrefix(trimmedLine, "```") { // Start of a nested block
				if fenceBalance != 0 {
					return spoilerBlock, false
				} // Cannot start nested if already started
				fenceBalance = 1
				inNestedCode = true
				result.WriteString("~~~" + strings.TrimPrefix(line, "```")) // Replace ``` with ~~~
			} else {
				// Just text inside spoiler, outside nested code
				result.WriteString(line)
			}
		} else { // Inside a nested code block
			if trimmedLine == "```" { // End of the nested block
				if fenceBalance != 1 {
					return spoilerBlock, false
				} // Cannot end if not started
				fenceBalance = 0
				inNestedCode = false
				result.WriteString("~~~") // Replace closing ``` with ~~~
			} else {
				// Code content inside the nested block
				result.WriteString(line)
			}
		}
		result.WriteByte('\n') // Add newline after processing each inner line
	}

	// After loop, check if nested code block was closed
	if inNestedCode || fenceBalance != 0 {
		return spoilerBlock, false // Unclosed nested fence
	}

	result.WriteString(lines[numLines-1]) // Write the final closing ``` fence for the spoiler

	// Determine if the original block ended with a newline *after* the final ```
	// This is tricky because Split removes the final newline if present.
	// Let's assume processSpoilerContent output should NOT end in \n unless the input block did.
	// The main loop handles newlines between blocks/lines.

	finalSpoilerResult := result.String()
	// If the original multi-line spoilerBlock string ended with \n, our result should too.
	// However, the current logic adds \n after each inner line. The final ``` is added without \n.
	// This seems correct for block-level processing.

	return finalSpoilerResult, true
}
