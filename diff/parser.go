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
		if strings.HasPrefix(line, "+") &&
			strings.Contains(strings.ToLower(line), "alter table") {
			return true
		}
	}
	return false
}