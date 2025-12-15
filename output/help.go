package output

import "fmt"

func Help() {
	Banner()

	fmt.Println(`Usage:
  git impact [options] [<ref>]

<ref> can be:
  main..feature
  v1.0..HEAD
  abc123

Options:
  --commit <sha>     analyze a specific commit
  --json             output machine-readable result
  --help             show this help
`)
}
