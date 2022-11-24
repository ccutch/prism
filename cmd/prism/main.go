package main

import (
	"flag"
	"fmt"
	"os"

	"prism.made-by-connor.com/prism"
)

func main() {
	deployCmd := flag.NewFlagSet("deploy", flag.ExitOnError)
	deploySource := deployCmd.String("src", ".", "Source directory.")
	deployDryRun := deployCmd.Bool("dry-run", false, "Dry run option if dry run is desired")

	if len(os.Args) < 2 {
		fmt.Println("No sub command provided from list: deploy")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "deploy":
		deployCmd.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if deployCmd.Parsed() {
		if *deployDryRun {
			fmt.Println("Dry run of deploy....")
			os.Exit(0)
		}
		if err := prism.Deploy(*deploySource); err != nil {
			panic(err)
		}
	}
}
