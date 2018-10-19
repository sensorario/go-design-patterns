# Concurrency » Barrier

## Description

The purpose of this pattern is to collect all the results from different
functions and goroutines before pass.

## Implementation

In this implementation we'll GET some urls (often used for testing purpose).
According to pattern description we will output somethine just after both
request have been received a response.

```go
func main() {
	barrier(
		"http://httpbin.org/headers",
		"http://httpbin.org/user-agent",
	)
}
```

This implementation is very idiomatic and uses `go` keyword. This keyword will
start new gorouting with function makeRequest. Also, this function contains a
channel called `in`.


```go
func barrier(urls ...string) {
	numOfRequsts := len(urls)

	in := make(chan Response, numOfRequsts)
	defer close(in)

	responses := make([]Response, numOfRequsts)

	for _, uri := range urls {
		go makeRequest(
			in,
			uri,
		)
	}

  // …
}
```

Instead, makeRequest function get the content, makes a covertion to string and
return to channel everythin.

```go
func makeRequest(out chan<- Response, url string) {
	// …

	resp, err := client.Get(url)
	byt, err := ioutil.ReadAll(resp.Body)
	res.Resp = string(byt)
	out <- res

	// …
}
```

Here the complete solution:

```go
var timeout int = 5000

type Response struct {
	Err  error
	Resp string
}

func barrier(urls ...string) {
	numOfRequsts := len(urls)

	in := make(chan Response, numOfRequsts)
	defer close(in)

	responses := make([]Response, numOfRequsts)

	for _, uri := range urls {
		go makeRequest(
			in,
			uri,
		)
	}

	var hasError bool
	for i := 0; i < numOfRequsts; i++ {
		resp := <-in
		if resp.Err != nil {
			fmt.Println("ERROR: ", resp.Err)
			hasError = true
		}
		responses[i] = resp
	}

	if !hasError {
		for _, resp := range responses {
			fmt.Println(resp.Resp)
		}
	}
}

func makeRequest(out chan<- Response, url string) {
	res := Response{}
	client := http.Client{
		Timeout: time.Duration(time.Duration(timeout) * time.Millisecond),
	}

	resp, err := client.Get(url)
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	res.Resp = string(byt)
	out <- res
}

func main() {
	barrier(
		"http://httpbin.org/headers",
		"http://httpbin.org/user-agent",
	)
}
```
