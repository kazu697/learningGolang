package app

import "testing"

func TestAverage(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}
	a := Average(s)
	expect := 3
	if a != expect {
		t.Errorf("expected %v but %v", expect, a)
	}
}
