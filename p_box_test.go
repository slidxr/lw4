package main

import (
	"testing"
)

func TestPboxPermutation(t *testing.T) {
	// Test cases
	tests := []struct {
		input    byte
		expected byte
	}{
		{0b10110011, 0b11001101},
		{0b11110000, 0b00001111},
		{0b00000000, 0b00000000},
		{0b11111111, 0b11111111},
		{0x0F, 0xF0},
		{0xFF, 0xFF},
		{0x00, 0x00},
	}

	for _, tt := range tests {
		// Test the P-Box permutation
		permuted := PBoxPermutation(tt.input)
		if permuted != tt.expected {
			t.Errorf("pBoxPermutation(%08b) = %08b, want %08b", tt.input, permuted, tt.expected)
		}

		// Test the reverse P-Box permutation
		reversed := ReversePBoxPermutation(permuted)
		if reversed != tt.input {
			t.Errorf("reversePBoxPermutation(%08b) = %08b, want %08b", permuted, reversed, tt.input)
		}
	}
}

func TestPBoxAndReversePBoxConsistency(t *testing.T) {
	for i := 0; i < 256; i++ { // Test all 8-bit values
		original := byte(i)
		substituted := PBoxPermutation(original)
		reversed := ReversePBoxPermutation(substituted)
		if original != reversed {
			t.Errorf("Mismatch after full S-Box and Reverse S-Box cycle for %02X: got %02X", original, reversed)
		}
	}
}
