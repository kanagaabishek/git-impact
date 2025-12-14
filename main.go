package main

import (
	"fmt"
	"git-impact/git"
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
	diff, err := git.Diff()
	if err != nil {
		fmt.Println("Failed to get git diff")
		return
	}
	fmt.Println(diff)
}
