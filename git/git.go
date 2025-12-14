package git

import (
	"os/exec"
	"strings"
)

func RepoRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func Diff() (string, error) {
	cmd := exec.Command("git", "diff", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func ShowCommit(sha string) (string, error) {
	cmd := exec.Command("git", "show", sha)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}