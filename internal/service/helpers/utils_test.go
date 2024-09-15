package helpers

import "testing"

func TestConvertHexToInt_Success(t *testing.T) {
	testCases := []struct {
		hexStr   string
		expected int64
	}{
		{"0x0", 0},
		{"0x1", 1},
		{"0xA", 10},
		{"0x10", 16},
		{"0xFF", 255},
		{"0x1000", 4096},
		{"0x7FFFFFFFFFFFFFFF", 9223372036854775807}, // max int64 value
	}

	for _, tc := range testCases {
		result, err := ConvertHexToInt(tc.hexStr)
		if err != nil {
			t.Errorf("ConvertHexToInt(%s) error = %v", tc.hexStr, err)
			continue
		}
		if result != tc.expected {
			t.Errorf("ConvertHexToInt(%s) = %d; want %d", tc.hexStr, result, tc.expected)
		}
	}
}

func TestConvertHexToInt_InvalidHex(t *testing.T) {
	testCases := []struct {
		hexStr string
	}{
		{"0xG"},                   // Invalid character
		{"0x"},                    // Empty value
		{"123"},                   // Missing '0x' prefix
		{"0x123ABCDEF1234567890"}, // Too large value
	}

	for _, tc := range testCases {
		_, err := ConvertHexToInt(tc.hexStr)
		if err == nil {
			t.Errorf("ConvertHexToInt(%s) = nil; expected error", tc.hexStr)
		}
	}
}
