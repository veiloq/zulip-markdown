package zlmd

import (
	"fmt"
	"time"
)

// ZLFormatTime formats a time.Time for Zulip's time formatting syntax.
//
// It converts a Go time.Time to Zulip's special time tag format, which allows
// Zulip to display the time according to the viewer's timezone settings.
//
// Example:
//
//	t := time.Date(2023, 5, 15, 14, 30, 0, 0, time.UTC)
//	result := ZLFormatTime(t)
//	// result will be something like:
//	// <time:2023-05-15T14:30:00Z>
func ZLFormatTime(t time.Time) string {
	return fmt.Sprintf("<time:%s>", t.Format(time.RFC3339))
}
