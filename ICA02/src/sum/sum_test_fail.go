package main

import 	"testing"


// Test int8
var test_int8 = []struct {
	n1       int8
	n2       int8
	expected int8
}{
	// Gir feil da resultatet blir feil
	{1, 2, 3},
	{4, 5, 3},
	{123, 2, 125},
}


func TestSumInt8(t *testing.T) {
	for _, v := range test_int8 {
		if val := SumInt8(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}



// Test int 32
var test_int32 = []struct {
	// Gir feil da det er int8
	n1       int32
	n2       int8
	expected int32
}{
	{1, 2, 3},
	{4, 5, 9},
	{123, 2, 125},
}

func TestSumInt32(t *testing.T) {
	for _, v := range test_int32 {
		if val := SumInt32(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

