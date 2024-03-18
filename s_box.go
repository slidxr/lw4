package main

var sBox = [16]byte{
	0xE, 0x4, 0xD, 0x1,
	0x2, 0xF, 0xB, 0x8,
	0x3, 0xA, 0x6, 0xC,
	0x5, 0x9, 0x0, 0x7,
}

var reverseSBox = [16]byte{
	0xE, 0x3, 0x4, 0x8,
	0x1, 0xC, 0xA, 0xF,
	0x7, 0xD, 0x9, 0x6,
	0xB, 0x2, 0x0, 0x5,
}

// SBoxSubstitution performs substitution using the sBox lookup table
func SBoxSubstitution(input byte) byte {
	// Split input into two 4-bit tetrads
	leftNibble := input >> 4
	rightNibble := input & 0x0F

	// Look up substitution values from sBox
	newLeftNibble := sBox[leftNibble]
	newRightNibble := sBox[rightNibble]

	// Combine substituted nibbles into a single byte
	return (newLeftNibble << 4) | newRightNibble
}

// ReverseSBoxSubstitution performs inverse substitution using the reverseSBox lookup table
func ReverseSBoxSubstitution(input byte) byte {
	leftNibble := input >> 4
	rightNibble := input & 0x0F

	// Look up substitution values from reverseSBox
	newLeftNibble := reverseSBox[leftNibble]
	newRightNibble := reverseSBox[rightNibble]

	// Combine substituted nibbles into a single byte
	return (newLeftNibble << 4) | newRightNibble
}
