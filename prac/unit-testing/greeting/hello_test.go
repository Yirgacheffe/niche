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

func TestHello_With_Empty_Arg(t *testing.T) {

	emptyExp := "Hello dude!"
	emptyAct := hello("")

	if emptyAct == emptyAct {
		t.Logf("Succeed!  Exp: %#v, Got %#v", emptyExp, emptyAct)
	} else {
		t.Errorf("Failed! Exp: %#v, Got %#v", emptyExp, emptyAct)
	}

}

func TestHello_With_Valid_Arg(t *testing.T) {

	resultExp := "Hello Mike!"
	resultAct := hello("Mike")

	if resultAct == resultExp {
		t.Logf("Succeed!  Exp: %#v, Got %#v", resultExp, resultAct)
	} else {
		t.Errorf("Failed! Exp: %#v, Got %#v", resultExp, resultAct)
	}

}
