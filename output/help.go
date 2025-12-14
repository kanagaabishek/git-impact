package output

import "fmt"

func Help() {
	Banner()

	fmt.Println(`Usage:
  git impact [options] [<ref>]

Options:
  --commit <sha>     analyze a specific commit
  --json             output machine-readable result
  --help             show this help

Examples:
  git impact
  git impact main..feature/auth
`)
}
