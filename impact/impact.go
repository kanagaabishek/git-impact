package impact

import (
	"strings"

	"git-impact/diff"
	"git-impact/rules"
)

type Result struct {
	Risk    string   `json:"risk"`
	Reasons []string `json:"reasons"`
}

func Analyze(diffText string) Result {
	reasons := []string{}
	riskScore := 0

	// Changed files
	files := diff.ChangedFiles(diffText)

	// Check if migration files are present
	hasMigration := false
	for _, file := range files {
		if strings.Contains(file, "/migration/") ||
			strings.Contains(file, "/migrations/") {
			hasMigration = true
		}
	}

	for _, file := range files {
		labels := rules.Classify(file)

		for _, label := range labels {
			switch label {
			case "AUTH":
				riskScore += 30
				reasons = append(reasons, "Authentication-related code modified")
			case "DATABASE":
				riskScore += 40
				reasons = append(reasons, "Database-related files changed")
			case "API":
				riskScore += 20
				reasons = append(reasons, "API layer modified")
			}
		}
	}

	//Schema-level detection
	if diff.ContainsAlterTable(diffText) {
		if hasMigration {
			riskScore += 20
			reasons = append(reasons, "Database schema changed with migration")
		} else {
			riskScore += 40
			reasons = append(reasons, "Database schema changed without migration")
		}
	}

	//Final risk calculation
	risk := "LOW"
	if riskScore >= 70 {
		risk = "HIGH"
	} else if riskScore >= 30 {
		risk = "MEDIUM"
	}

	return Result{
		Risk:    risk,
		Reasons: unique(reasons),
	}
}

func unique(input []string) []string {
	seen := make(map[string]bool)
	result := []string{}

	for _, v := range input {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}
