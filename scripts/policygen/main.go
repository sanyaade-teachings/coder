package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/coder/coder/v2/coderd/database/spice/policy"
	"github.com/coder/coder/v2/scripts/policygen/gen"
)

var destination = flag.String("destination", "", "destination file")

func main() {
	flag.Parse()

	out, err := gen.Generate(policy.Schema)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

	if *destination == "" {
		fmt.Println(out)
		return
	}

	fmt.Println("Writing to", *destination)
	//err = os.WriteFile(*destination, []byte(out), 0644)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}
