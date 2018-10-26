# Concurrency » Pipeline

## Description

A pipeline consists of a chain of elements arranged so that the output of each
  element is the input of the next.

## Implementation

Each element of the pipeline is a function like the following function. In this
piece of code is visible the isolation of current step. It is also important to
get that inside goroutine the for treat all data received by channel input.
When all data in input is managed the goroutine ends and finally the channel is
returned. This means that each step's variable are close to current scope.

```go
type x struct { }

func step(in <-chan x) <-chan x {
	out := make(chan x, 100)
	go func() {
		for i := range in {
			// do something
		}
		close(out)
	}()
	return out
}
```

There also be the beginning function that build the pipeline. In this case
first step build a bundle of integers. Following steps sum all the integers,
append the sum with a string ":foo" and finally append the string ":bar".


```go
func StartPipeline(amount int) string {
	source := generator(amount)
	sum := sum(source)
	foo := appendFoo(sum)
	return <-appendBar(foo)
}
```

```go
func generator(max int) <-chan int {
	outChInt := make(chan int, 100)
	go func() {
		for i := 1; i <= max; i++ {
			outChInt <- i
		}
		close(outChInt)
	}()
	return outChInt
}
```

And finally main code.

```go
func main() {
	fmt.Println("pipeline")
	res := StartPipeline(4)
	fmt.Println(res)
}
```

