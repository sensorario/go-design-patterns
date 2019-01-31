package main

import "testing"

func TestProcessZeroItemsWheneverQueueIsEmpty(t *testing.T) {
	c := CommandInvoker{}
	c.processQueue()
	if c.ProcessedItems() != 0 {
		t.Error("no items should be processed with empty queue")
	}
}

func TestCountProcessedItems(t *testing.T) {
	c := CommandInvoker{}
	c.addToQueue(&SomeCommand{"foo"})
	c.processQueue()
	if c.ProcessedItems() != 1 {
		t.Error("no items should be processed with empty queue")
	}
}

func TestQueueIsProcessedWheneverContainsThreeItems(t *testing.T) {
	c := CommandInvoker{}

	c.addToQueue(&SomeCommand{"foo"})
	c.addToQueue(&SomeCommand{"foo"})
	if c.ProcessedItems() != 0 {
		t.Error("no items should be processed")
	}

	c.addToQueue(&SomeCommand{"foo"})
	if c.ProcessedItems() != 3 {
		t.Error("three items should be processed")
	}
}
