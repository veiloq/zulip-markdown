// Package zlmd provides utility functions for formatting text output.
package zlmd

import (
	"fmt"
	"strings"
)

// NoErr formats an error message with a red X emoji if the error is not nil
// and appends it to the provided StringBuilder. It returns the original error
// to allow for error propagation.
func NoErr(info *strings.Builder, err error) error {
	if err != nil {
		info.WriteString(fmt.Sprintf("❌ %s\n", err))
		return err
	}
	return nil
}

// NoErrWarn formats a warning message with a warning emoji if the error is not nil
// and appends it to the provided StringBuilder. It returns the original error
// to allow for error propagation.
func NoErrWarn(info *strings.Builder, err error) error {
	if err != nil {
		info.WriteString(fmt.Sprintf("⚠️ %s\n", err))
		return err
	}
	return nil
}

// Warnf formats a warning message with a warning emoji and appends it to
// the provided StringBuilder. It supports printf-style formatting.
func Warnf(info *strings.Builder, format string, args ...interface{}) {
	info.WriteString(fmt.Sprintf("⚠️ %s\n", fmt.Sprintf(format, args...)))
}

// WarnUsage formats a usage warning message with a warning emoji and appends it
// to the provided StringBuilder. It's specifically for showing command usage help.
func WarnUsage(info *strings.Builder, format string, args ...interface{}) {
	Warnf(info, "Usage: `%s`.", fmt.Sprintf(format, args...))
}

// Infof formats an informational message with an info emoji and appends it to
// the provided StringBuilder. It supports printf-style formatting.
func Infof(info *strings.Builder, format string, args ...interface{}) {
	info.WriteString(fmt.Sprintf("ℹ️ %s\n", fmt.Sprintf(format, args...)))
}

// Successf formats a success message with a checkmark emoji and appends it to
// the provided StringBuilder. It supports printf-style formatting.
func Successf(info *strings.Builder, format string, args ...interface{}) {
	info.WriteString(fmt.Sprintf("✅ %s\n", fmt.Sprintf(format, args...)))
}

// Errorf formats an error message with a red X emoji and appends it to
// the provided StringBuilder. It supports printf-style formatting.
func Errorf(info *strings.Builder, format string, args ...interface{}) {
	info.WriteString(fmt.Sprintf("❌ %s\n", fmt.Sprintf(format, args...)))
}

// Debugf formats a debug message with a magnifying glass emoji and appends it to
// the provided StringBuilder. It supports printf-style formatting.
func Debugf(info *strings.Builder, format string, args ...interface{}) {
	info.WriteString(fmt.Sprintf("🔍 %s\n", fmt.Sprintf(format, args...)))
}

type ArrowType int

const (
	ArrowLeft        = "←"
	ArrowRight       = "→"
	ArrowLeftRight   = "↔"
	ArrowRightDotted = "⤑"
	ArrowLeftDotted  = "⬸"
)

// Point formats a text with a right arrow emoji and appends it to the provided StringBuilder.
// It also appends a newline if newline is true.
//
// Example:
// Point(info, "Hello, world!") -> " → Hello, world!"
func Point(info *strings.Builder, text string) {
	info.WriteString(fmt.Sprintf(" %s %s", ArrowRight, text))
}

// Pointnl formats a text with a right arrow emoji and appends it to the provided StringBuilder.
// It also appends a newline.
//
// Example:
// Pointnl(info, "Hello, world!") -> " → Hello, world!\n"
func Pointnl(info *strings.Builder, text string) {
	Point(info, fmt.Sprintf("%s\n", text))
}

// Right formats a text with a right arrow emoji and appends it to the provided StringBuilder.
// It also appends a newline if newline is true.
//
// Example:
// Right(info, "A", "B", "C") -> "A → B → C"
func Right(info *strings.Builder, text ...string) {
	arrow(info, ArrowRight, true, text...)
}

// Left formats a text with a left arrow emoji and appends it to the provided StringBuilder.
// It also appends a newline if newline is true.
//
// Example:
// Left(info, "A", "B", "C") -> "A ← B ← C"
func Left(info *strings.Builder, text ...string) {
	arrow(info, ArrowLeft, true, text...)
}

// LeftRight formats a text with a left and right arrow emoji and appends it to the provided StringBuilder.
// It also appends a newline if newline is true.
//
// Example:
// LeftRight(info, "A", "B", "C") -> "A ↔ B ↔ C"
func LeftRight(info *strings.Builder, text ...string) {
	arrow(info, ArrowLeftRight, true, text...)
}

// RightDotted formats a text with a right dotted arrow emoji and appends it to the provided StringBuilder.
// It also appends a newline if newline is true.
//
// Example:
// RightDotted(info, "A", "B", "C") -> "A ⤑ B ⤑ C"

func RightDotted(info *strings.Builder, text ...string) {
	arrow(info, ArrowRightDotted, true, text...)
}

// LeftDotted formats a text with a left dotted arrow emoji and appends it to the provided StringBuilder.
// It also appends a newline if newline is true.
//
// Example:
// LeftDotted(info, "A", "B", "C") -> "A ⬸ B ⬸ C"
func LeftDotted(info *strings.Builder, text ...string) {
	arrow(info, ArrowLeftDotted, true, text...)
}

// Rightnl formats a text with a right arrow emoji and appends it to the provided StringBuilder.
// It also appends a newline.
//
// Example:
// Rightnl(info, "A", "B", "C") -> "A → B → C\n"
func Rightnl(info *strings.Builder, text string) {
	arrow(info, ArrowRight, true, text)
}

// Leftnl formats a text with a left arrow emoji and appends it to the provided StringBuilder.
// It also appends a newline.
func Leftnl(info *strings.Builder, text string) {
	arrow(info, ArrowLeft, true, text)
}

// RightDottednl formats a text with a right dotted arrow emoji and appends it to the provided StringBuilder.
// It also appends a newline.
//
// Example:
// RightDottednl(info, "A", "B", "C") -> "A ⤑ B ⤑ C\n"
func RightDottednl(info *strings.Builder, text string) {
	arrow(info, ArrowRightDotted, true, text)
}

// LeftDottednl formats a text with a left dotted arrow emoji and appends it to the provided StringBuilder.
// It also appends a newline.
//
// Example:
// LeftDottednl(info, "A", "B", "C") -> "A ⬸ B ⬸ C\n"
func LeftDottednl(info *strings.Builder, text string) {
	arrow(info, ArrowLeftDotted, true, text)
}

// arrow formats a text with an arrow emoji and appends it to the provided StringBuilder.
// It also appends a newline if newline is true.
//
// Example:
// arrow(info, ArrowRight, true, "A", "B", "C") -> "A → B → C\n"
func arrow(info *strings.Builder, arrowType string, newline bool, text ...string) {
	if len(text) == 1 {
		info.WriteString(fmt.Sprintf("%s %s", arrowType, text[0]))
	} else {
		info.WriteString(strings.Join(text, fmt.Sprintf(" %s ", arrowType)))
	}
	if newline {
		info.WriteString("\n")
	}
}

// Badge creates a colored badge/tag for important information.
// style can be: "primary", "success", "warning", "danger", "info"
func Badge(info *strings.Builder, text string, style string) {
	var emoji string

	switch style {
	case "primary":
		emoji = "🔵"
	case "success":
		emoji = "✅"
	case "warning":
		emoji = "⚠️"
	case "danger":
		emoji = "❌"
	case "info":
		emoji = "ℹ️"
	case "rejected":
		emoji = "✴️"
	default:
		emoji = "🧷"
	}

	info.WriteString(fmt.Sprintf("%s `%s`\n", emoji, text))
}

// Usage formats a usage message with a warning emoji and appends it to
// the provided StringBuilder. It supports printf-style formatting.
//
// Example:
// Usage(info, "help", "arg1", "arg2") -> "Usage: `/help [arg1] [arg2]`."
func Usage(info *strings.Builder, alias string, opts ...string) {
	// Format options with brackets if provided
	optStr := ""
	if len(opts) > 0 {
		formattedOpts := make([]string, len(opts))
		for i, opt := range opts {
			formattedOpts[i] = "[" + opt + "]"
		}
		optStr = " " + strings.Join(formattedOpts, " ")
	}

	Warnf(info, "Usage: `%s%s`.", alias, optStr)
}

// CommandInfo formats a command info message with a command alias and
// description and appends it to the provided StringBuilder.
//
// Example:
// CommandInfo(info, "help", "Show help information") -> "- **help** - Show help information\n"
func CommandInfo(info *strings.Builder, alias string, description string) {
	info.WriteString(fmt.Sprintf("- **%s** - %s\n", alias, description))
}
