## Cryptographic Algorithms Demonstration

This project demonstrates the implementation of two fundamental cryptographic concepts: the S-Box (Substitution-Box) and the P-Box (Permutation-Box). These algorithms are pivotal in many encryption and decryption processes, providing the foundation for data security and integrity.
Description

The S-Box algorithm applies a non-linear transformation (substitution) to the input data, enhancing the data's resistance to cryptanalysis by making the relationship between the ciphertext and the plaintext complex.

The P-Box algorithm, on the other hand, permutes (shuffles) the bits of the input data. This step disperses the input bits across the output, contributing to diffusion in the encrypted data.

This implementation showcases the effect of these algorithms on an 8-bit binary input and demonstrates their reversible nature, allowing the original data to be retrieved through inverse operations.

No external dependencies are required to run this project. It is developed in Go and utilizes the standard library packages:

How to Run

    Ensure Go is installed on your machine.
    Clone this repository or download the source code.
    Navigate to the project directory via the command line.
    Execute the program by running:

go run .

#### Program Usage

Upon running the program, you will be prompted to choose between the S-Box Substitution 

(1) and the P-Box Permutation 

(2). Input the number corresponding to your choice and press Enter.

Next, you'll be asked to enter an 8-bit binary number (e.g., 10111011). 

Ensure your input is exactly 8 bits long and consists of 0s and 1s only.

Example Interaction

Choose algorithm: 
1. S-Box Substitution
2. P-Box Permutation 
 
Enter the number of the algorithm: 1

Enter an 8-bit input (binary): 11010011

S-Box substitution: 11010011 -> 01101001

Reverse S-Box substitution: 01101001 -> 11010011

Expected Results

The program will display the result of the chosen cryptographic operation on your input, followed by the inverse operation's result, verifying the reversibility of the process.

For example, for S-Box Substitution:

    The output will show the substituted binary value.
    Following, the program will show the result of reversing the S-Box substitution, which should match the original input.

For P-Box Permutation:

    The permuted binary value based on the defined permutation pattern will be displayed.
    The inverse permutation result, which should revert to the original input, will be shown afterward.

