package rules

import "strings"

func Classify(file string) []string {
	labels := []string{}
	f := strings.ToLower(file)

	if isAuthFile(f) {
		labels = append(labels, "AUTH")
	}
	if isDatabaseFile(f) {
		labels = append(labels, "DATABASE")
	}
	if isAPIFile(f) {
		labels = append(labels, "API")
	}
	return labels
}

func isAuthFile(file string) bool {
	return strings.Contains(file, "/auth/") ||
		strings.Contains(file, "/security/") ||
		strings.HasSuffix(file, "Config") ||
		strings.HasSuffix(file, "Auth")
}

func isDatabaseFile(file string) bool {
	return strings.Contains(file, "/db/") ||
		strings.Contains(file, "/migration/") ||
		strings.HasSuffix(file, ".sql")
}

func isAPIFile(file string) bool {
	return strings.Contains(file, "/controller/") ||
		strings.Contains(file, "/routes/")
}
