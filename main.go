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
	flag.Parse()

	if showVersion {
		if version == "" {
			version = "devel"
		}

		fmt.Printf("hello-release-please %s\n", version)
		os.Exit(0)
	}

	fmt.Println("hello release-please")
}
