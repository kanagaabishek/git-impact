package main

import (
	"fmt"
	"os"

	"impact/cmd"
	"impact/diff"
	"impact/git"
	"impact/output"
	"impact/rules"
)

func main() {

	opts := cmd.ParseFlags()
	
	if opts.Help {
		output.Help()
		return
	}

	//check if git repo
	root, err := git.RepoRoot()
	if err != nil {
		fmt.Println("Not a git repository")
		return
	}
	fmt.Println("Git repo detected at:", root)
	
	// get git diff (or show a specific commit if requested)
	var diffText string
	if opts.Commit != "" {
		diffText, err = git.ShowCommit(opts.Commit)
	} else {
		diffText, err = git.Diff()
	}

	if err != nil {
		fmt.Println("Failed to read git diff:", err)
		os.Exit(1)
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
