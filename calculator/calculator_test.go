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
			t.Errorf("Add(%d,%d) = %d, want %d", entry.a, entry.b, result, entry.result)
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
			t.Errorf("Difference(%d,%d) = %d, want %d", entry.a, entry.b, result, entry.result)
		}
	}
}

func TestInt32ToBytes(t *testing.T) {
	tests := []struct {
		input int32
		result []byte
	}{
		{4, []byte{4,0,0,0}},
	}

	for _, test := range tests {
		result := Int32ToBytes(test.input)
		for i, _ := range result {
			if result[i] != test.result[i] {
				t.Errorf("Int32ToBytes(%d)[%d] = %d, want %d", test.input, i, result[i], test.result[i])
			}
		}
	}

}

func TestGetBitOnPosition(t *testing.T) {
	tests := []struct {
		bytes    []byte
		position uint
		expected bool
	}{
		{[]byte{0,0,0,0}, 0, false},
		{[]byte{1,0,0,0}, 0, true},
		{[]byte{4,0,0,0}, 0, false},
		{[]byte{4,0,0,0}, 2, true},
		{[]byte{0,2,0,0}, 8, false},
		{[]byte{0,2,0,0}, 9, true},
	}

	for _, test := range tests {
		result := GetBitFromPosition(test.bytes, test.position)
		if result != test.expected {
			t.Errorf("GetBitFromPosition([% x], %d) = %b, want %b", test.bytes, test.position, result, test.expected)
		}
	}
}