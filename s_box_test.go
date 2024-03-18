package main

import (
	"testing"
)

func TestSboxSubstitution(t *testing.T) {
	tests := []struct {
		input    byte
		expected byte
	}{
		{0x00, 0xEE},
		{0b10100110, 0b01101011},
	}

	for _, tt := range tests {
		substituted := SBoxSubstitution(tt.input)
		if substituted != tt.expected {
			t.Errorf("sboxSubstitution(%02X) = %02X, want %02X", tt.input, substituted, tt.expected)
		}
	}
}

func TestReverseSboxSubstitution(t *testing.T) {
	tests := []struct {
		substituted byte
		original    byte
	}{
		{0xEE, 0x00},
		{0b01101011, 0b10100110},
	}

	for _, tt := range tests {
		reversed := ReverseSBoxSubstitution(tt.substituted)
		if reversed != tt.original {
			t.Errorf("reverseSboxSubstitution(%02X) = %02X, want %02X", tt.substituted, reversed, tt.original)
		}
	}
}

func TestSBoxAndReverseSBoxConsistency(t *testing.T) {
	for i := 0; i < 256; i++ { // Test all 8-bit values
		original := byte(i)
		substituted := SBoxSubstitution(original)
		reversed := ReverseSBoxSubstitution(substituted)
		if original != reversed {
			t.Errorf("Mismatch after full S-Box and Reverse S-Box cycle for %02X: got %02X", original, reversed)
		}
	}
}
