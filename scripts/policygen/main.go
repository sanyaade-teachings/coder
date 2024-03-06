package main

import (
	"fmt"
	"os"

	"github.com/coder/coder/v2/coderd/database/spice/policy"
	"github.com/coder/coder/v2/scripts/policygen/gen"
)

func main() {
	out, err := gen.Generate(policy.Schema)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Println(out)
}
