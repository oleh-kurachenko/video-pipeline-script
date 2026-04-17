package main

import (
	"fmt"
	"io"
	"os"
	"slices"

	"github.com/MatthiasPetermann/sysexits"
	"github.com/fatih/color"
)

func printUsage(writeTo io.Writer) {
	bold := color.New(color.Bold).SprintfFunc()
	italic := color.New(color.Italic).SprintfFunc()

	_, _ = fmt.Fprintf(writeTo, "usage: %s %s %s\n", os.Args[0],
		bold("encode|decode"), italic("directory ..."))
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printUsage(os.Stderr)
		os.Exit(sysexits.EX_USAGE)
	}
	if slices.Contains([]string{
		"help", "--help", "-h", "-?", "usage"}, args[0]) {

		printUsage(os.Stdout)
		os.Exit(sysexits.EX_OK)
	}

	switch args[0] {
	case "encode":
		_, _ = fmt.Fprintf(os.Stderr, "not yet supported command: %s\n",
			args[0])
		os.Exit(sysexits.EX_UNAVAILABLE)
	case "decode":
		err := decode(args[1:])
		if _, ok := err.(*decodingError); ok {
			_, _ = fmt.Fprintf(os.Stderr, "invalid args: %s\n", err)
			os.Exit(sysexits.EX_USAGE)
		}
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "unexpected error: %v\n", err)
			os.Exit(sysexits.EX_SOFTWARE)
		}
	default:
		_, _ = fmt.Fprintf(os.Stderr, "unknown command: %s\n", args[0])
		printUsage(os.Stderr)
		os.Exit(sysexits.EX_USAGE)
	}
}
