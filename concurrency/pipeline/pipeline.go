package main

import "fmt"

func StartPipeline(amount int) int {
	return <-power(
		generator(amount),
	)
}

func power(in <-chan int) <-chan int {
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
