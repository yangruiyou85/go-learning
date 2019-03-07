package loop_test

import "testing"

func TestWhileLop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}

}
