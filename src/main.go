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
	if slices.Contains([]string{"help", "--help", "-h", "-?", "usage"},
		args[0]) {

		printUsage(os.Stdout)
		os.Exit(sysexits.EX_OK)
	}
	if len(args) == 1 {
		printUsage(os.Stderr)
		os.Exit(sysexits.EX_USAGE)
	}

	switch args[0] {
	case "encode":
		_, _ = fmt.Fprintf(os.Stderr, "not yet supported command: %s\n",
			args[0])
		os.Exit(sysexits.EX_UNAVAILABLE)
	case "decode":
		_, _ = fmt.Fprintf(os.Stderr, "not yet supported command: %s\n",
			args[0])
		os.Exit(sysexits.EX_UNAVAILABLE)
	default:
		_, _ = fmt.Fprintf(os.Stderr, "unknown command: %s\n", args[0])
		printUsage(os.Stderr)
		os.Exit(sysexits.EX_USAGE)
	}
}
