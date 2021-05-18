package calc

import "testing"

func TestAdd(t *testing.T) {
	expected := 5

	if actual := Add(3, 2); actual != expected {
		t.Errorf("Add() = %v, want %v", actual, expected)
	}
}
