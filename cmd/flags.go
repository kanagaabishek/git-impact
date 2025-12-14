package cmd

import "flag"

type Options struct {
	Help   bool
	JSON   bool
	Commit string
}

func ParseFlags() Options {
	var opts Options

	flag.BoolVar(&opts.Help, "help", false, "show help")
	flag.BoolVar(&opts.JSON, "json", false, "output json")
	flag.StringVar(&opts.Commit, "commit", "", "analyze specific commit")

	flag.Parse()
	return opts
}
