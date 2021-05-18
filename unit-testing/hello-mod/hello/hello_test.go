package hello

import "testing"

func TestHello(t *testing.T) {

	expected := "Hello, world."

	if actual := Hello(); actual != expected {
		t.Errorf("Hello() = %q, want %q", actual, expected)
	}

}
