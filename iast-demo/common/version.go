package common

import (
	"flag"
	"fmt"
	"os"
)

var (
	Name      string
	Version   string
	BuildTime string
)

func ShowVersion() {
	if len(os.Args) == 2 {
		version := flag.Bool("version", false,
			"Shows the version information and exits")
		flag.Parse()
		if *version {
			if Version != "" {
				fmt.Printf("%s %s (%s)\n", Name, Version, BuildTime)
			} else {
				fmt.Println("No version information found")
			}
			os.Exit(0)
		}
	}
}
