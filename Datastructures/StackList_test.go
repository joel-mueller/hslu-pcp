package Datastructures

import (
	"testing"
)

func TestPushAndPopList(t *testing.T) {
	s := StackList[int]{}

	s.Push(10)
	s.Push(20)

	if s.Size() != 2 {
		t.Errorf("Expected length 2, got %d", s.Size())
	}

	val := s.Pop()
	if val != 20 {
		t.Errorf("Expected 20, got %d", val)
	}

	val = s.Pop()
	if val != 10 {
		t.Errorf("Expected 10, got %d", val)
	}
}

func TestPeekList(t *testing.T) {
	s := StackList[string]{}
	s.Push("hello")
	s.Push("world")

	if s.Peek() != "world" {
		t.Errorf("Expected 'world', got %s", s.Peek())
	}
	if s.Size() != 2 {
		t.Errorf("Peek should not remove element, but length is %d", s.Size())
	}
}

func TestEmptyList(t *testing.T) {
	s := StackList[float64]{}

	if !s.Empty() {
		t.Error("Stack should be empty")
	}
	s.Push(3.14)
	if s.Empty() {
		t.Error("Stack should not be empty after push")
	}
	s.Pop()
	if !s.Empty() {
		t.Error("Stack should be empty after pop")
	}
}

func TestPopOnEmptyPanicsList(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic on Pop from empty stack")
		}
	}()
	s := StackList[int]{}
	s.Pop()
}

func TestPeekOnEmptyPanicsList(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic on Peek from empty stack")
		}
	}()
	s := StackList[int]{}
	s.Peek()
}
