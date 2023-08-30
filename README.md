# Password List Generator

This is a simple Go program that generates random strings based on various options.

## Compilation

To compile the program, make sure you have Go installed on your system. Then, run the following command in your terminal:

```bash
go build pwlist.go
```

## Usage

The program generates random strings with customizable options. Here are the available command line arguments:

- `-n` (default: 10): Length of the random strings to generate.
- `-q` (default: 10): Number of random strings to generate.
- `-a`: Include lowercase letters in the generated strings.
- `-A`: Include uppercase letters in the generated strings.
- `-1`: Include numbers in the generated strings.
- `-spe`: Include special characters in the generated strings.

Example usage:

```bash
./pwlist -n 12 -a -1 -q 5
```

This command generates 5 random strings, each with a length of 12 characters, including lowercase letters and numbers.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
