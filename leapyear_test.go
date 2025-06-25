package main

import (
    "testing"
)

func TestLeapYear01(t *testing.T) {
	tests := []struct {
    year     int
    expected bool
    }{
        {2000, true}, 
        {1900, false}, 
        {2024, true},  
        {2023, false}, 
        {1600, true}, 
        {2100, false}, 
    }
	for _, test := range tests {
        expected := leapyear(test.year)
        if expected != test.expected {
            t.Errorf("leapyear(%d) = %v; want %v", test.year, expected, test.expected)
        }
    }
}