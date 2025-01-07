package utils

import (
	"regexp"
	"strings"
)

var (
	spaces            = regexp.MustCompile(`\s+`)
	notAllowed        = regexp.MustCompile(`[^\p{L}\p{N}\s.-]+`) // Allow spaces and hyphens
	athePrefix        = regexp.MustCompile(`^(a-|the-)`)
	consecutiveDashes = regexp.MustCompile(`-+`) // Matches multiple consecutive hyphens
)

// GenerateSlug generates a URL-friendly slug from a given string.
func GenerateSlug(input string) string {
	// Replace special characters
	result := input
	result = strings.ReplaceAll(result, "%", " percent")
	result = strings.ReplaceAll(result, "#", " sharp")

	// Remove not allowed characters but preserve spaces and hyphens
	result = notAllowed.ReplaceAllString(result, "")

	// Replace spaces with hyphens
	result = spaces.ReplaceAllString(result, "-")

	// Normalize consecutive hyphens
	result = consecutiveDashes.ReplaceAllString(result, "-")

	// Convert to lowercase
	result = strings.ToLower(result)

	// Remove specific prefixes
	result = athePrefix.ReplaceAllString(result, "")

	// Trim trailing hyphens
	result = strings.Trim(result, "-")

	return result
}
