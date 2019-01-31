package main

import (
	"flag"
)

func main() {
	strategy := flag.String("strategy", "console", "selected strategy")
	flag.Parse()

	var printStrategy PrintStrategy

	switch *strategy {
	case "console":
		printStrategy = &ConsoleStrategy{}
	case "file":
		printStrategy = &FileStrategy{"bigciao"}
	default:
		printStrategy = &ConsoleStrategy{}
	}

	printStrategy.Print()
}
