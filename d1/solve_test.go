package main

import "testing"

func TestSum(t *testing.T) {
	s1, s2 := sum("2x3one")

	if s1 != 23 {
		t.Errorf("sum1 = %d, want %d", s1, 23)
	}
	if s2 != 21 {
		t.Errorf("sum2 = %d, want %d", s2, 21)
	}
}

func TestSum2(t *testing.T) {
	// vvbfnnine3ngv 33 93

	s1, s2 := sum("vvbfnnine3ngv")

	if s1 != 33 {
		t.Errorf("sum1 = %d, want %d", s1, 33)
	}
	if s2 != 93 {
		t.Errorf("sum2 = %d, want %d", s2, 93)
	}
}

func TestSum3(t *testing.T) {
	// 3lzfjpcthreeonenine 33 39

	s1, s2 := sum("3lzfjpcthreeonenine")

	if s1 != 33 {
		t.Errorf("sum1 = %d, want %d", s1, 33)
	}
	if s2 != 39 {
		t.Errorf("sum2 = %d, want %d", s2, 39)
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum("vvbfnnine3ngv")
	}
}
