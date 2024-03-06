package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		usage()
		return
	}

	switch os.Args[1] {
	case "export":
		fmt.Println(PlaygroundExport())
	default:
		usage()
	}
}

func usage() {
	fmt.Println("Usage: policycmd [command]")
	fmt.Println("Commands:")
	fmt.Println("  export")
	// fmt.Println("  import")
	//fmt.Println("  playground")
	//fmt.Println("  run")
	//fmt.Println("  test")
}
