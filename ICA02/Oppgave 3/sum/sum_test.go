package main

import 	"testing"


// Test int8
var test_int8 = []struct {
	n1       int8
	n2       int8
	expected int8
}{
	// Number + Number = Number
	{1, 2, 3},
	{4, 5, 9},
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
	n1       int32
	n2       int32
	expected int32
}{
	// Number + Number = Number
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

// Test int64
var test_int64 = []struct {
	n1       int64
	n2       int64
	expected int64
}{
	// Number + Number = Number
	{1, 2, 3},
	{4, 5, 9},
	{123, 2, 125},
}

func TestSumInt64(t *testing.T) {
	for _, v := range test_int64 {
		if val := SumInt64(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

// Test uint32
var test_uint32 = []struct {
	n1       uint32
	n2       uint32
	expected uint32
}{
	// Number + Number = Number
	{1, 2, 3},
	{4, 5, 9},
	{123, 2, 125},
}

func TestSumUint32(t *testing.T) {
	for _, v := range test_uint32 {
		if val := SumUint32(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

// Test float32
var test_SumFloat32 = []struct {
	n1       float32
	n2       float32
	expected float32
}{
	// Number + Number = Number
	{1, 2, 3},
	{4, 5, 9},
	{123, 2, 125},
}

func TestSumFloat32(t *testing.T) {
	for _, v := range test_SumFloat32 {
		if val := SumFloat32(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

// Test float64
var test_float64 = []struct {
	n1       float64
	n2       float64
	expected float64
}{
	// Number + Number = Number
	{1, 2, 3},
	{4, 5, 9},
	{123, 2, 125},
}

func TestSumFloat64(t *testing.T) {
	for _, v := range test_float64 {
		if val := SumFloat64(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}