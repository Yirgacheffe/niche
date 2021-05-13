package calc

import "testing"

// TestMathAdd do unit test for add integer number
func TestMathAdd(t *testing.T) {

	result, _ := Add(1, 2, 3)

	if result == 6 {
		t.Logf("Add(1,2,3) PASSED, expected: %v, got: %v", 6, result)
	} else {
		t.Errorf("Add(1,2,3) FAILED, expected: %v, got: %v", 6, result)
	}

}
