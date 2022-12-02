package main

import (
	"flag"
	"fmt"
	"os"
)

// Populated during build.
var version string
var showVersion bool

func main() {
	flag.BoolVar(&showVersion, "version", false, "Display the version and exit")
	flag.Usage = func() {
		fmt.Println("Usage: hello-release-please <name> [flags...]\n\nFlags:")
		flag.PrintDefaults()
	}
	flag.Parse()

	if showVersion {
		if version == "" {
			version = "devel"
		}

		fmt.Printf("hello-release-please %s\n", version)
		os.Exit(0)
	}

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("hello %s\n", flag.Arg(0))
}
