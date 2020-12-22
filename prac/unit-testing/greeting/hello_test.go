package main

import "testing"

const (
	expEmpty  = "Hello dude!"
	expResult = "Hello Mike!"
)

func TestHello(t *testing.T) {

	actEmpty := hello("")
	if expEmpty != actEmpty {
		t.Errorf("hello(\"\") failed, excepted %#v, got %#v", expEmpty, actEmpty)
	} else {
		t.Logf("hello(\"\") success, excepted %#v, got %#v", expEmpty, actEmpty)
	}

	actResult := hello("Mike")
	if actResult != expResult {
		t.Errorf("hello(\"Mike\") failed, excepted %#v, got %#v", expResult, actResult)
	} else {
		t.Logf("hello(\"Mike\") success, excepted %#v, got %#v", expResult, actResult)
	}

}
