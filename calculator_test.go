package calculator

import "testing"

func TestAddNumbers(t *testing.T) {
	tests := []struct {
		a int32
		b int32
		result int32
	}{
		{1,2,3},
		{0,0,0},
		{-1, -2,-3},
	}

	for _, entry := range tests {
		result := Add(entry.a, entry.b)
		if result != entry.result {
			t.Error("Add(%d,%d) = %d, want %d", entry.a, entry.b, result, entry.result)
		}
	}
}

func  TestDifference(t *testing.T)  {
	tests := []struct{
		a int32
		b int32
		result int32
	}{
		{1,2,-1},
		{0,0,0},
		{-1, -2,1},
	}
	for _, entry := range tests {
		result := Difference(entry.a, entry.b)
		if result != entry.result {
			t.Error("Difference(%d,%d) = %d, want %d", entry.a, entry.b, result, entry.result)
		}
	}
}
