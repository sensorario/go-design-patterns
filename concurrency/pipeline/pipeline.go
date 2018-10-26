package main

import "fmt"
import "strings"
import "strconv"

func StartPipeline(amount int) string {
	source := generator(amount)
	sum := sum(source)
	foo := appendFoo(sum)
	return <-appendBar(foo)
}

func appendBar(in <-chan string) <-chan string {
	out := make(chan string, 100)
	go func() {
		bar := <-in
		out <- string(strings.Join([]string{bar, "bar"}, ":"))
		close(out)
	}()
	return out
}

func appendFoo(in <-chan int) <-chan string {
	out := make(chan string, 100)
	go func() {
		foo := <-in
		out <- string(strings.Join([]string{strconv.Itoa(foo), "foo"}, ":"))
		close(out)
	}()
	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int, 100)
	go func() {
		var sum int
		for v := range in {
			sum += v
		}
		out <- sum
		close(out)
	}()
	return out
}

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

func main() {
	fmt.Println("pipeline")
	res := StartPipeline(4)
	fmt.Println(res)
}
