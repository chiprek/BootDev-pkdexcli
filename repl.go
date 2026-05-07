package main

import "strings"

func cleanInput(text string) []string {
	sanitized := strings.ToLower(text)
	words := strings.Fields(sanitized)
	return words
}
