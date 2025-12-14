package rules

import "strings"


//simple classification based on file path keywords
func Classify(file string) []string {
	labels := []string{}

	if strings.Contains(file, "auth") || strings.Contains(file, "security") {
		labels = append(labels, "AUTH")
	}
	if strings.Contains(file, "db") || strings.Contains(file, "schema") {
		labels = append(labels, "DATABASE")
	}
	if strings.Contains(file, "controller") || strings.Contains(file, "api") || strings.Contains(file, "routes") {
		labels = append(labels, "API")
	}
	return labels
}

