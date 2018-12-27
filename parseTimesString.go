package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// parseTimeFilterString parses a string and returns a string that is usable in
// GetTransactionsWithFilters
func parseTimeFilterString(filter string) string {
	fmt.Printf("*%s*", filter)

	filter = strings.TrimSpace(filter)
	fmt.Printf("*%s*", filter)
	months := []string{
		"jan", "january",
		"feb", "february",
		"mar", "march",
		"apr", "april",
		"may", "may",
		"jun", "june",
		"jul", "july",
		"aug", "august",
		"sep", "september",
		"oct", "october",
		"nov", "november",
		"dec", "december",
	}
	now := time.Now()

	// If time is just a number its an ISO week -> return data from that week
	if weekOfYear, err := strconv.Atoi(filter); err == nil {
		_, currentWeek := now.ISOWeek()
		if weekOfYear <= currentWeek {
			return fmt.Sprintf("strftime('%%W', date) = '%d'", weekOfYear)
		}
		// Display last years data
		return fmt.Sprintf("strftime('%%W', date, '-1 year') = '%d'", weekOfYear)

	}

	// Check if user input is a month
	if inSlice(filter, months) {
		fmt.Println("they've entered a months name O_O")
	}

	// Default -> return date from current month
	return "date(date) >= date('now', 'start of month') AND " +
		"date(date) <= date('now', '+1 month', 'start of month', '-1 day')"
}
