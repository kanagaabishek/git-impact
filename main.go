package main

import (
	"encoding/json"
	"fmt"
	"os"

	"git-impact/cmd"
	"git-impact/git"
	"git-impact/impact"
	"git-impact/output"
)

func main() {

	opts := cmd.ParseFlags()

	if opts.Help {
		output.Help()
		return
	}

	output.Banner()
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

	result := impact.Analyze(diffText)

	if opts.JSON {
		b, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			fmt.Fprintln(os.Stderr, "failed to marshal result:", err)
			os.Exit(1)
		}
		fmt.Println(string(b))
		return
	}

	fmt.Println("Overall Risk:", result.Risk)
	for _, r := range result.Reasons {
		fmt.Println("•", r)
	}

}
