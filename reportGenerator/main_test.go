package main

import (
	"testing"
	"time"
)

func TestMatchesCronPart(t *testing.T) {
	tests := []struct {
		value    int
		cronPart string
		expect   bool
	}{
		{2, "*", true},
		{2, "*/2", true},
		{2, "*/3", false},
		{5, "*/3", false},
		{6, "*/3", true},
		{1, "1,2,3", true},
		{4, "1,2,3", false},
		{12, "*/5,10", false},
		{15, "*/5,10", true},
		{10, "*/5,10", true},
		{15, "7,8,19", false},
		{8, "7,8,19", true},
		{19, "7,8,19", true},
	}

	for _, test := range tests {
		result := matchesCronPart(test.value, test.cronPart)
		if result != test.expect {
			t.Errorf("For value %d and cronPart %s, expected %t, but got %t", test.value, test.cronPart, test.expect, result)
		}
	}
}

func TestMatchCronExpression(t *testing.T) {
	tests := []struct {
		cronExpression string
		date           time.Time
		expect         bool
	}{
		{"*/2,3,5,17 1,2,8 */2,5", time.Date(2023, time.August, 17, 0, 0, 0, 0, time.UTC), true},
		{"*/2,3,5,17 1,2,8 */2,5", time.Date(2023, time.August, 18, 0, 0, 0, 0, time.UTC), true},
		{"*/2,3,5,17 1,2,8 */2,5", time.Date(2023, time.March, 18, 0, 0, 0, 0, time.UTC), false},
		{"*/2,3,5,17 1,2,8 */2,5", time.Date(2023, time.August, 16, 0, 0, 0, 0, time.UTC), false},
		{"*/2,3,5,17 1,2,8 3", time.Date(2023, time.August, 16, 0, 0, 0, 0, time.UTC), true},
		{"1 * *", time.Date(2023, time.August, 1, 0, 0, 0, 0, time.UTC), true},
		{"1 * *", time.Date(2023, time.August, 2, 0, 0, 0, 0, time.UTC), false},
		{"* * 3", time.Date(2023, time.August, 16, 0, 0, 0, 0, time.UTC), true},
		{"* * 4", time.Date(2023, time.August, 16, 0, 0, 0, 0, time.UTC), false},
	}

	for _, test := range tests {
		result := matchCronExpression(test.date, test.cronExpression)
		if result != test.expect {
			t.Errorf("For cronExpression %s and date %v, expected %t, but got %t", test.cronExpression, test.date, test.expect, result)
		}
	}
}
