package main

import (
	"errors"
	"testing"
)

func TestSuccessCallbackIsCalledWhenExecuteFuncSucceeds(t *testing.T) {
	called := false

	s := Subject{}
	s.Success(func(m string) {
		called = true
	}).Failure(func(e error) {
		t.Error("expected success, got failure")
	}).Execute(func(_ int) (string, error) {
		return "ok", nil
	})

	if !called {
		t.Error("expected success callback to be called")
	}
}

func TestFailureCallbackIsCalledWhenExecuteFuncFails(t *testing.T) {
	called := false

	s := Subject{}
	s.Success(func(m string) {
		t.Error("expected failure, got success")
	}).Failure(func(e error) {
		called = true
	}).Execute(func(_ int) (string, error) {
		return "", errors.New("forced error")
	})

	if !called {
		t.Error("expected failure callback to be called")
	}
}

func TestSuccessCallbackReceivesReturnedString(t *testing.T) {
	var received string

	s := Subject{}
	s.Success(func(m string) {
		received = m
	}).Failure(func(e error) {
		t.Error("unexpected failure")
	}).Execute(func(_ int) (string, error) {
		return "hello future", nil
	})

	if received != "hello future" {
		t.Errorf("expected 'hello future', got '%s'", received)
	}
}

func TestFailureCallbackReceivesReturnedError(t *testing.T) {
	var received error
	expected := errors.New("something went wrong")

	s := Subject{}
	s.Success(func(m string) {
		t.Error("unexpected success")
	}).Failure(func(e error) {
		received = e
	}).Execute(func(_ int) (string, error) {
		return "", expected
	})

	if received == nil || received.Error() != expected.Error() {
		t.Errorf("expected error '%v', got '%v'", expected, received)
	}
}

func TestSuccessMethodReturnsSubjectForChaining(t *testing.T) {
	s := &Subject{}
	result := s.Success(func(string) {})
	if result != s {
		t.Error("expected Success to return the same Subject pointer")
	}
}

func TestFailureMethodReturnsSubjectForChaining(t *testing.T) {
	s := &Subject{}
	result := s.Failure(func(error) {})
	if result != s {
		t.Error("expected Failure to return the same Subject pointer")
	}
}