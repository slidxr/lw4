package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Choose algorithm:")
	fmt.Println("1. S-Box Substitution")
	fmt.Println("2. P-Box Permutation")
	fmt.Print("Enter the number of the algorithm: ")

	algorithmInput, _ := reader.ReadString('\n')
	algorithmInput = strings.TrimSpace(algorithmInput)
	algorithmChoice, err := strconv.Atoi(algorithmInput)

	if err != nil || (algorithmChoice != 1 && algorithmChoice != 2) {
		fmt.Println("Invalid input. Please choose either 1 or 2.")
		return
	}

	fmt.Print("Enter an 8-bit input (binary): ")
	bitInput, _ := reader.ReadString('\n')
	bitInput = strings.TrimSpace(bitInput)

	if len(bitInput) != 8 {
		fmt.Println("Invalid input. Please enter exactly 8 bits.")
		return
	}

	input, err := strconv.ParseUint(bitInput, 2, 8)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid 8-bit binary number.")
		return
	}

	switch algorithmChoice {
	case 1:
		sboxOutput := SBoxSubstitution(uint8(input))
		fmt.Printf("S-Box substitution: %08b -> %08b\n", input, sboxOutput)
		sboxReversed := ReverseSBoxSubstitution(sboxOutput)
		fmt.Printf("Reverse S-Box substitution: %08b -> %08b\n", sboxOutput, sboxReversed)
	case 2:
		pboxOutput := PBoxPermutation(uint8(input))
		fmt.Printf("P-Box Permutation: %08b -> %08b\n", input, pboxOutput)
		pboxReversed := ReversePBoxPermutation(pboxOutput)
		fmt.Printf("P-Box Permutation Output: %08b -> %08b\n", pboxOutput, pboxReversed)
	}
}
