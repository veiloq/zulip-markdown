package zlmd

import (
	"testing"
	"time"
)

func TestZLFormatTime(t *testing.T) {
	// Test with a fixed known time.
	testTime := time.Date(2023, time.January, 2, 15, 4, 5, 0, time.UTC)
	expected := "<time:" + testTime.Format(time.RFC3339) + ">"
	result := ZLFormatTime(testTime)
	if result != expected {
		t.Errorf("ZLFormatTime(%v) = %q; want %q", testTime, result, expected)
	}
}

func TestZLFormatTime_ZeroTime(t *testing.T) {
	// Test with the zero time value.
	var zeroTime time.Time
	expected := "<time:" + zeroTime.Format(time.RFC3339) + ">"
	result := ZLFormatTime(zeroTime)
	if result != expected {
		t.Errorf("ZLFormatTime(%v) = %q; want %q", zeroTime, result, expected)
	}
}
