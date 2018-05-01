package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"text/template"
)

type PrintStrategy interface {
	Print() error
}

type ConsoleStrategy struct{}

type FileStrategy struct {
	DestinationFilePath string
}

func (c *ConsoleStrategy) Print() error {
	fmt.Println("ConsoleStrategy")
	lister, _ := template.New("foo").Parse(tplTemplate())
	lister.Execute(os.Stdout, tplParams())
	return nil
}

func (c *FileStrategy) Print() error {
	fmt.Println("FileStrategy")
	var t bytes.Buffer
	foo, _ := template.New("bar").Parse(tplTemplate())
	foo.Execute(&t, tplParams())

	f, err := os.Create(c.DestinationFilePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write(t.Bytes())
	return nil
}

func tplParams() map[string]interface{} {
	items := []int{1, 1, 2, 3, 5, 8}
	return map[string]interface{}{
		"items": items,
		"last":  len(items) - 1,
	}
}

func tplTemplate() string {
	return "" +
		"{{range $i, $el := .items}}" +
		"{{$el}}" +
		"{{if eq $i $.last}}.{{else}}, {{end}}" +
		"{{end}}"
}

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
