package translatedassert

import (
	"testing"
)

type Example struct {
	a        interface{}
	b        interface{}
	expected interface{}
}

func TestOpADD(t *testing.T) {
	tests := []Example{
		// int
		{2, 3, 2 + 3},
		{int32(1), 1, int32(1) + 1},
		{1, int32(1), 1 + int32(1)},
		{int64(2147483647), 1, int64(2147483647) + 1},
		{1, int64(2147483647), 1 + int64(2147483647)},
		// uint
		{uint(2), uint(4), uint(2) + uint(4)},
		{uint8(2), uint8(4), uint8(2) + uint8(4)},
		{uint16(2), uint16(4), uint16(2) + uint16(4)},
		{uint32(2), uint32(4), uint32(2) + uint32(4)},
		{uint64(2), uint64(4), uint64(2) + uint64(4)},
		// float
		{2.3, 2.5, 2.3 + 2.5},
		{2.3, 1, 2.3 + 1},
		{1, 2.3, 1 + 2.3},
		{float32(2.3), float32(1), float32(2.3) + float32(1)},
		{float32(2.3), 1, float32(2.3) + 1},
		{1, float32(2.3), 1 + float32(2.3)},
		{float64(2.3), float64(1), float64(2.3) + float64(1)},
		{float64(2.3), 1, float64(2.3) + 1},
		{1, float64(2.3), 1 + float64(2.3)},
		// complex
		{(1 + 1i), (2 + 1i), (1 + 1i) + (2 + 1i)},
		{1, (2 + 1i), 1 + (2 + 1i)},
		{(2 + 1i), 1, (2 + 1i) + 1},
		{1.1, (2 + 1i), 1.1 + (2 + 1i)},
		{(2 + 1i), 1.1, 1.1 + (2 + 1i)},
		{complex(1, 2), complex(1, 1), complex(1, 2) + complex(1, 1)},
		{complex(1, 2), 1, complex(1, 2) + 1},
		{1, complex(1, 2), 1 + complex(1, 2)},
		{complex64(complex(1, 2)), complex64(complex(1, 1)), complex64(complex(1, 2)) + complex64(complex(1, 1))},
		{complex64(complex(1, 2)), 1, complex64(complex(1, 2)) + 1},
		{1, complex64(complex(1, 2)), 1 + complex64(complex(1, 2))},
		{complex128(complex(1, 2)), complex128(complex(1, 1)), complex128(complex(1, 2)) + complex128(complex(1, 1))},
		{complex128(complex(1, 2)), 1, complex128(complex(1, 2)) + 1},
		{1, complex128(complex(1, 2)), 1 + complex128(complex(1, 2))},
		// string
		{"a", "b", "ab"},
	}

	for _, test := range tests {
		got := Op("ADD", test.a, test.b)
		if !eq(test.expected, got) {
			t.Errorf(`Op("ADD", %v, %v) = %v, but got %v`, test.a, test.b, test.expected, got)
		}
	}
}

func TestOpWithUserDefinedType(t *testing.T) {
	type myint int
	var a myint = 1
	type myfloat float64
	var b myfloat = 2.5
	tests := []Example{
		{myint(1), 3, myint(1) + 3},
		{3, myint(1), 3 + myint(1)},
		{a, 3, a + 3},
		{3, a, 3 + a},
		{a, a, a + a},
		{b, 1, b + 1},
		{1, b, 1 + b},
		{b, b, b + b},
		{b, b, b + b},
	}

	for _, test := range tests {
		got := Op("ADD", test.a, test.b)
		if !eq(test.expected, got) {
			t.Errorf(`Op("ADD", %v, %v) = %v, but got %v`, test.a, test.b, test.expected, got)
		}
	}
}

func TestOpShiftWithUserDefinedType(t *testing.T) {
	type myint int
	type myuint uint
	tests := []Example{
		{myint(1), 1, myint(1) << 1},
		{1, myuint(1), 1 << myuint(1)},
		{myint(1), myuint(1), myint(1) << myuint(1)},
	}

	for _, test := range tests {
		got := OpShift("SHL", test.a, test.b)
		if !eq(test.expected, got) {
			t.Errorf(`OpShift("SHL", %v, %v) = %v, but got %v`, test.a, test.b, test.expected, got)
		}
	}
}

func TestOpSUB(t *testing.T) {
	tests := []Example{
		{2, 3, -1},
		{uint(5), uint(4), uint(1)},
		{2.3, 2.5, -0.2},
		{1 + 1i, 2 + 1i, -1 + 0i},
	}

	for _, test := range LeftRightTests(tests) {
		got := Op("SUB", test.a, test.b)
		if !eq(test.expected, got) {
			t.Errorf(`Op("SUB", %v, %v) = %v, but got %v`, test.a, test.b, test.expected, got)
		}
	}
}

func TestOpMUL(t *testing.T) {
	tests := []Example{
		{2, 3, 6},
		{uint(5), uint(4), uint(20)},
		{2.3, 2.5, 5.75},
		{1 + 1i, 2 + 1i, 1 + 3i},
	}

	for _, test := range LeftRightTests(tests) {
		got := Op("MUL", test.a, test.b)
		if !eq(test.expected, got) {
			t.Errorf(`Op("MUL", %v, %v) = %v, but got %v`, test.a, test.b, test.expected, got)
		}
	}
}

func TestOpQUO(t *testing.T) {
	tests := []Example{
		{15, 3, 5},
		{14, 3, 4},
		{uint(5), uint(2), uint(2)},
		{5.0, 2.5, 2.0},
		{1.0 + 1.0i, 2.0 + 1.0i, 0.6 + 0.2i},
	}

	for _, test := range LeftRightTests(tests) {
		got := Op("QUO", test.a, test.b)
		if !eq(test.expected, got) {
			t.Errorf(`Op("QUO", %v, %v) = %v, but got %v`, test.a, test.b, test.expected, got)
		}
	}
}

func TestOpREM(t *testing.T) {
	tests := []Example{
		{15, 3, 0},
		{14, 3, 2},
		{uint(5), uint(2), uint(1)},
	}

	for _, test := range LeftRightTests(tests) {
		got := Op("REM", test.a, test.b)
		if !eq(test.expected, got) {
			t.Errorf(`Op("REM", %v, %v) = %v, but got %v`, test.a, test.b, test.expected, got)
		}
	}
}

func TestOpAND(t *testing.T) {
	tests := []Example{
		// 010, 110, 010
		{2, 6, 2},
		// 101, 001, 001
		{uint(5), uint(1), uint(1)},
	}

	for _, test := range LeftRightTests(tests) {
		got := Op("AND", test.a, test.b)
		if !eq(test.expected, got) {
			t.Errorf(`Op("AND", %v, %v) = %v, but got %v`, test.a, test.b, test.expected, got)
		}
	}
}

func TestOpOR(t *testing.T) {
	tests := []Example{
		// 010, 110, 110
		{2, 6, 6},
		// 101, 010, 111
		{uint(5), uint(2), uint(7)},
	}

	for _, test := range LeftRightTests(tests) {
		got := Op("OR", test.a, test.b)
		if !eq(test.expected, got) {
			t.Errorf(`Op("OR", %v, %v) = %v, but got %v`, test.a, test.b, test.expected, got)
		}
	}
}

func TestOpXOR(t *testing.T) {
	tests := []Example{
		// 010, 110, 100
		{2, 6, 4},
		// 101, 010, 111
		{uint(5), uint(2), uint(7)},
	}

	for _, test := range LeftRightTests(tests) {
		got := Op("XOR", test.a, test.b)
		if !eq(test.expected, got) {
			t.Errorf(`Op("XOR", %v, %v) = %v, but got %v`, test.a, test.b, test.expected, got)
		}
	}
}

func TestOpANDNOT(t *testing.T) {
	tests := []Example{
		// 010, 110, 000
		{2, 6, 0},
		// 010, 000, 010
		{2, 0, 2},
		// 101, 010, 101
		{uint(5), uint(2), uint(5)},
	}

	for _, test := range LeftRightTests(tests) {
		got := Op("ANDNOT", test.a, test.b)
		if !eq(test.expected, got) {
			t.Errorf(`Op("ANDNOT", %v, %v) = %v, but got %v`, test.a, test.b, test.expected, got)
		}
	}
}

func TestOpSHL(t *testing.T) {
	tests := []Example{
		{1, uint(2), 4},
		{uint(2), uint(2), uint(8)},
	}

	for _, test := range leftIntRightUintTests(tests) {
		got := OpShift("SHL", test.a, test.b)
		if !eq(test.expected, got) {
			t.Errorf("OpSHL(%v, %v) = %v, but got %v", test.a, test.b, test.expected, got)
		}
	}
}

func TestOpSHR(t *testing.T) {
	tests := []Example{
		{4, uint(2), 1},
		{uint(2), uint(2), uint(0)},
		{uint(2), uint8(2), uint(0)},
	}

	for _, test := range leftIntRightUintTests(tests) {
		got := OpShift("SHR", test.a, test.b)
		if !eq(test.expected, got) {
			t.Errorf(`Op("SHR", %v, %v) = %v, but got %v`, test.a, test.b, test.expected, got)
		}
	}
}

func TestOpLAND(t *testing.T) {
	tests := []Example{
		{true, true, true},
		{true, false, false},
		{false, true, false},
		{false, false, false},
	}

	for _, test := range tests {
		got := Op("LAND", test.a, test.b)
		if !eq(test.expected, got) {
			t.Errorf(`Op("LAND", %v, %v) = %v, but got %v`, test.a, test.b, test.expected, got)
		}
	}
}

func TestOpLOR(t *testing.T) {
	tests := []Example{
		{true, true, true},
		{true, false, true},
		{false, true, true},
		{false, false, false},
	}

	for _, test := range tests {
		got := Op("LOR", test.a, test.b)
		if !eq(test.expected, got) {
			t.Errorf(`Op("LOR", %v, %v) = %v, but got %v`, test.a, test.b, test.expected, got)
		}
	}
}

// t `op` t = t
func LeftRightTests(tests []Example) []Example {
	for _, test := range tests {
		if _, ok := test.a.(int); ok {
			a := test.a.(int)
			b := test.b.(int)
			expected := test.expected.(int)

			tests = append(
				tests,
				Example{int8(a), int8(b), int8(expected)},
				Example{int16(a), int16(b), int16(expected)},
				Example{int32(a), int32(b), int32(expected)},
				Example{int64(a), int64(b), int64(expected)},
			)
		}

		if _, ok := test.expected.(uint); ok {
			a := test.a.(uint)
			b := test.b.(uint)
			expected := test.expected.(uint)

			tests = append(
				tests,
				Example{uint8(a), uint8(b), uint8(expected)},
				Example{uint16(a), uint16(b), uint16(expected)},
				Example{uint32(a), uint32(b), uint32(expected)},
				Example{uint64(a), uint64(b), uint64(expected)},
			)
		}

		if _, ok := test.expected.(float64); ok {
			a := test.a.(float64)
			b := test.b.(float64)
			expected := test.expected.(float64)

			tests = append(
				tests,
				Example{float32(a), float32(b), float32(expected)},
			)
		}

		if _, ok := test.expected.(complex128); ok {
			a := test.a.(complex128)
			b := test.b.(complex128)
			expected := test.expected.(complex128)

			tests = append(
				tests,
				Example{complex64(a), complex64(b), complex64(expected)},
			)
		}
	}

	return tests
}

// uint* `op` int*|uint* = int*|uint*
func leftIntRightUintTests(tests []Example) []Example {
	for _, test := range tests {
		if a, ok := test.a.(int); ok {
			if b, ok := test.b.(uint); ok {
				expected := test.expected.(int)

				tests = append(
					tests,
					// int, uint*
					Example{int(a), uint8(b), int(expected)},
					Example{int(a), uint16(b), int(expected)},
					Example{int(a), uint32(b), int(expected)},
					Example{int(a), uint64(b), int(expected)},
					// int8, uint*
					Example{int8(a), uint8(b), int8(expected)},
					Example{int8(a), uint16(b), int8(expected)},
					Example{int8(a), uint32(b), int8(expected)},
					Example{int8(a), uint64(b), int8(expected)},
					// int16, uint*
					Example{int16(a), uint8(b), int16(expected)},
					Example{int16(a), uint16(b), int16(expected)},
					Example{int16(a), uint32(b), int16(expected)},
					Example{int16(a), uint64(b), int16(expected)},
					// int32, uint*
					Example{int32(a), uint8(b), int32(expected)},
					Example{int32(a), uint16(b), int32(expected)},
					Example{int32(a), uint32(b), int32(expected)},
					Example{int32(a), uint64(b), int32(expected)},
					// int64, uint*
					Example{int64(a), uint8(b), int64(expected)},
					Example{int64(a), uint16(b), int64(expected)},
					Example{int64(a), uint32(b), int64(expected)},
					Example{int64(a), uint64(b), int64(expected)},
				)
			}
		}
		if a, ok := test.a.(uint); ok {
			if b, ok := test.b.(uint); ok {
				expected := test.expected.(uint)

				tests = append(
					tests,
					// uint, uint*
					Example{uint(a), uint8(b), uint(expected)},
					Example{uint(a), uint16(b), uint(expected)},
					Example{uint(a), uint32(b), uint(expected)},
					Example{uint(a), uint64(b), uint(expected)},
					// uint8, uint*
					Example{uint8(a), uint8(b), uint8(expected)},
					Example{uint8(a), uint16(b), uint8(expected)},
					Example{uint8(a), uint32(b), uint8(expected)},
					Example{uint8(a), uint64(b), uint8(expected)},
					// uint16, uint*
					Example{uint16(a), uint8(b), uint16(expected)},
					Example{uint16(a), uint16(b), uint16(expected)},
					Example{uint16(a), uint32(b), uint16(expected)},
					Example{uint16(a), uint64(b), uint16(expected)},
					// uint32, uint*
					Example{uint32(a), uint8(b), uint32(expected)},
					Example{uint32(a), uint16(b), uint32(expected)},
					Example{uint32(a), uint32(b), uint32(expected)},
					Example{uint32(a), uint64(b), uint32(expected)},
					// uint64, uint*
					Example{uint64(a), uint8(b), uint64(expected)},
					Example{uint64(a), uint16(b), uint64(expected)},
					Example{uint64(a), uint32(b), uint64(expected)},
					Example{uint64(a), uint64(b), uint64(expected)},
				)
			}
		}
	}

	return tests
}

func eq(a, b interface{}) bool {
	if a, af := a.(float64); af {
		if b, bf := b.(float64); bf {
			return (a-b) < 0.00000001 && (b-a) < 0.00000001
		}
	}
	if a, af := a.(float32); af {
		if b, bf := b.(float32); bf {
			return (a-b) < 0.0000001 && (b-a) < 0.0000001
		}
	}

	return a == b
}
