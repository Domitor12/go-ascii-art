# Go Text Formatter

A powerful command-line text processing tool written in Go that applies various transformations to text files, including hexadecimal/binary conversion, case modifications, punctuation fixing, and grammar corrections.

## Features

- **Number Conversion**: Convert hexadecimal `(hex)` and binary `(bin)` markers to decimal numbers
- **Case Transformations**: Apply capitalization `(cap)`, uppercase `(up)`, and lowercase `(low)` commands to single or multiple words
- **Grammar Correction**: Automatically change "a" to "an" before vowel sounds or silent 'h'
- **Punctuation Fixing**: Properly space punctuation marks (commas, periods, exclamation points, etc.)
- **Quote Spacing**: Remove extra spaces around single and double quotation marks
- **Batch Processing**: Handle multiple transformations in a single pass

## Installation

### Prerequisites
- Go 1.16 or higher installed on your system

### Clone and Build

```bash
git clone https://github.com/Domitor12/text-formatter-go.git
cd text-formatter-go
go build -o textfmt

Or run directly without building:
bash

go build -o textfmt && ./textfmt input.txt output.txt

Usage
Basic Command
bash

go run . input.txt output.txt

Or using the compiled binary:
bash

./textfmt input.txt output.txt

Arguments

    input.txt - Path to the source text file

    output.txt - Path where the formatted text will be saved

Transformation Rules
1. Number Conversions
Input	Output
1E (hex)	30
1010 (bin)	10
2. Case Transformations
Command	Effect	Example
(cap)	Capitalizes the previous word	hello (cap) → Hello
(cap, 2)	Capitalizes 2 previous words	hello world (cap, 2) → Hello World
(up)	Uppercases the previous word	hello (up) → HELLO
(up, 3)	Uppercases 3 previous words	hello world go (up, 3) → HELLO WORLD GO
(low)	Lowercases the previous word	HELLO (low) → hello
(low, 2)	Lowercases 2 previous words	HELLO WORLD (low, 2) → hello world
3. Grammar Rules
Input	Output
a apple	an apple
a hour	an hour
a honest	an honest
4. Punctuation Spacing
Input	Output
hello , world	hello, world
BAMM !!	BAMM!!
why ?	why?
5. Quote Handling
Input	Output
' hello '	'hello'
He said ' hi '	He said 'hi'
Example
Input File (input.txt)
text

This is 1E (hex) in decimal. hello (cap) world! it (up) is a amazing day. 
I have 1010 (bin) apples . He said ' awesome ' and left . 
(up,2) this is important (cap, 3) example text 

Output File (output.txt)
text

This is 30 in decimal. Hello world! IT is an amazing day. 
I have 10 apples. He said 'awesome' and left. 
THIS IS Important Example Text

How It Works

The program processes text in the following order:

    Read the input file

    Split text into words (handling all whitespace automatically)

    Convert hex and binary numbers to decimal

    Apply case transformations (cap, up, low)

    Fix grammar (a → an)

    Clean punctuation spacing

    Fix quote spacing

    Write formatted text to output file

Error Handling

The program checks for common errors:

    Missing input/output file arguments

    Input file not found

    Permission issues when reading/writing files

    Invalid number formats in conversion commands

Development
Project Structure
text

text-formatter-go/
├── readwords.go      # Main program with all functions
├── input.txt         # Example input file
├── output.txt        # Example output file
└── README.md         # This file

Key Functions
Function	Purpose
processHexBin()	Converts hexadecimal and binary numbers
processCapUpLow()	Handles case transformation commands
applyTransformation()	Applies individual case changes
capitalize()	Properly capitalizes a single word
fixAtoAn()	Changes "a" to "an" before vowels/h
fixPunctuation()	Fixes spacing around punctuation
fixQuoteSpacing()	Removes extra spaces around quotes
isPunctuation()	Checks if a character is punctuation
Limitations

    Commands must be placed AFTER the words they affect

    Nested or overlapping commands may produce unexpected results

    Only supports single-byte characters (ASCII/UTF-8 single-byte)

    The (cap, n) format requires proper spacing or no space between comma and number

Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
Areas for Improvement

    Add support for Unicode characters

    Implement recursive command processing

    Add more punctuation handling options

    Create a web interface

    Add unit tests

License

This project is open source and available under the MIT License.
Author

VINCENT ONYECHEREM IKENNA
Acknowledgments

    Built as a text formatting tool following standard text processing requirements

    Uses Go's standard library packages: fmt, os, strings

Happy Text Formatting! 📝
