package main

import (
	"fmt"
	"git-impact/git"
	"git-impact/diff"
	"git-impact/rules"
)

func main() {

	//check if git repo
	root, err := git.RepoRoot()
	if err != nil {
		fmt.Println("Not a git repository")
		return
	}
	fmt.Println("Git repo detected at:", root)
	
	//get git diff
	diffText, err := git.Diff()
	if err != nil {
		fmt.Println("Failed to get git diffText")
		return
	}
	fmt.Println(diffText)

	//get changed files from diffText
	files := diff.ChangedFiles(diffText)
	fmt.Println("Changed files:", files)

	//classify changed files
	for _, file := range files {
		labels := rules.Classify(file)
		fmt.Println(file, "→", labels)
	}

	//check for database schema changes
	if diff.ContainsAlterTable(diffText) {
	fmt.Println("⚠️ Database schema change detected")
}

}
