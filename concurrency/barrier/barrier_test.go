package main

import "testing"

func TestDidAllUrlsDidLoaded(t *testing.T) {
	barrier(
		"http://httpbin.org/headers",
		"http://httpbin.org/user-agent",
	)
}
