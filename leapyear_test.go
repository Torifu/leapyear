package main

import (
    "testing"
)

func TestLeapYear01(t *testing.T) {
	tests := []struct {
    year     int
    expected bool
    }{
        {2000, true},   // Return true for the year is divisible by both 4 and 100 but it is divisible by 400 
        {2024, true},   // Return true for the year is divisible by 4
        {2023, false},  // Return false for any year 
        {2100, false},  // Return false for the year is divisible by 4 but it is divisible by 100
        {1, false},  // Return false for any year
    }
	for _, test := range tests {
        expected := leapyear(test.year)
        if expected != test.expected {
            t.Errorf("leapyear(%d) = %v; want %v", test.year, expected, test.expected)
        }
    }
}