# Behavioral » Strategy

## Elements

 - Context
 - Strategy
 - Concrete Strategy

## Description

The strategy pattern uses different algorithms to achieve some specific
functionality. These algorithms are hidden behind an interface and, of course,
they must be interchangeable. All algorithms achieve same functionality in a
different way.

For example, `io.Writer` interface defines a strategy ti write, and the
functionality is always the same.

## Implementation

Main feature of this pattern is that we can have different algorithms that uses
same interface. In this example we'll print some content to a different output
using different strategies:

 - FileStrategy
 - ConsoleStrategy

The strategy will be selected using a console argument.

All strategies must implement same interface:

```go
type PrintStrategy interface {
	Print() error
}
```

As said few words ago, we have two strategies. One to print contents in console
and another for files.

```go
type ConsoleStrategy struct{}

type FileStrategy struct {
	DestinationFilePath string
}
```

Each strategy must implement same `Print()` method.

```go
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
```

Code of `tplParams()` and `tplTemplate()` is omitted here. It provide a simple
template and parameters to build it.

Main program will get from console the flag "strategy". If omitted,
ConsoleStrategy will be selected and the template rendered will be visible in
console. Otherwise, if strategy flag is sent with:

> $ go run strategy.go --strategy=file

Same content is saved into a file called "bigciao".

```go
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
```
