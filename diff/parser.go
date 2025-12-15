package diff

import "strings"

func ChangedFiles(diff string) []string {
	lines := strings.Split(diff, "\n")
	files := []string{}

	for _, line := range lines {
		if strings.HasPrefix(line, "diff --git") {
			parts := strings.Split(line, " ")
			if len(parts) >= 4 {
				file := strings.TrimPrefix(parts[3], "b/")
				files = append(files, file)
			}
		}
	}
	return files
}

func ContainsAlterTable(diff string) bool {
	lines := strings.Split(diff, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Only added SQL lines
		if strings.HasPrefix(line, "+") &&
			!strings.HasPrefix(line, "+++") {

			lower := strings.ToLower(line)
			if strings.HasPrefix(lower, "+alter table") {
				return true
			}
		}
	}
	return false
}
