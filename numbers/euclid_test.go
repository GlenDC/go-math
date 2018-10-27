package numbers

import "testing"

func TestGCD(t *testing.T) {
	testCases := []struct {
		a, b, gcd int
	}{
		{0, 0, 0},
		{2322, 0, 0},
		{0, 100, 100},
		{2322, 654, 6},
		{654, 2322, 6},
		{10, -100, 10},
		{-100, 3, 1},
		{-100, -3, 1},
	}
	for idx, testCase := range testCases {
		gcd := GCD(testCase.a, testCase.b)
		if gcd != testCase.gcd {
			t.Errorf("%d. GCD(%d,%d)=%d != %d", idx, testCase.a, testCase.b, gcd, testCase.gcd)
		}
	}
}
