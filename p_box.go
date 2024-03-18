package main

// permutation is a slice of integers representing the order of bit permutation
var permutation = []int{7, 6, 5, 4, 3, 2, 1, 0}

// PBoxPermutation takes a byte as input and returns a permuted byte
// based on the predefined permutation order
func PBoxPermutation(input byte) byte {
	var output byte = 0

	for i, bitPos := range permutation {
		// Extract the bit at the current position from the input byte
		bit := (input >> bitPos) & 1

		// Assign the extracted bit to the new position in the output byte
		output |= (bit << i)
	}

	return output
}

// ReversePBoxPermutation takes a permuted byte as input and returns
// the original byte by reversing the permutation
func ReversePBoxPermutation(input byte) byte {
	var output byte = 0

	for i, bitPos := range permutation {
		// Extract the bit at the current position from the input byte
		bit := (input >> i) & 1

		// Assign the extracted bit to the original position in the output byte
		output |= (bit << bitPos)
	}

	return output
}
