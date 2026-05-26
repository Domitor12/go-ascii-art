ASCII Art Generator

A powerful command-line tool written in Go that converts text into beautiful ASCII art banners. This tool supports multiple font styles (standard, shadow, and thinkertoy) and can handle multi-line text input with \n support.

Features

    Multiple Font Styles: Choose from three distinct ASCII art styles

        standard - Classic clean banner style using spaces, underscores, and pipes

        shadow - Creates a 3D shadow effect with filled characters

        thinkertoy - Artistic style with unique symbols and patterns

    Multi-line Support: Use \n in your text to create multiple lines of ASCII art

    Full Character Support: Handles all printable ASCII characters (codes 32-126)

        Letters: A-Z, a-z

        Numbers: 0-9

        Punctuation: .,!?;: etc.

        Special characters: @#$%^&*() etc.

    Blank Line Handling: Automatically adds appropriate spacing between text lines

Installation
Prerequisites

    Go 1.16 or higher installed on your system

Clone and Build
bash

git clone https://github.com/Domitor12/ascii-art-generator-go.git
cd ascii-art-generator-go
go build -o asciiart

Or Run Directly
bash

go build -o asciiart && ./asciiart "Hello World" standard

Usage
Basic Syntax
bash

go run . "your text here" [banner_style]

Or using the compiled binary:
bash

./asciiart "your text here" [banner_style]

Arguments
Argument	Required	Description
"text"	Yes	The text to convert to ASCII art (use quotes for multi-word text)
[banner_style]	No	Banner style: standard, shadow, or thinkertoy (defaults to standard)
Examples
1. Basic Usage with Standard Banner
bash

go run . "Hello"

Output:
text

 _    _          _   _
| |  | |        | | | |
| |__| |   ___  | | | |   ___
|  __  |  / _ \ | | | |  / _ \
| |  | | |  __/ | | | | | (_) |
|_|  |_|  \___| |_| |_|  \___/

2. Shadow Banner Style
bash

go run . "GO" shadow

Output (shadow style):
text

▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒
▒▒░  ░▒▒▒▒▒▒▒▒▒▒▒▒░  ░▒▒
▒▒░  ░▒▒▒▒▒▒▒▒▒▒▒▒░  ░▒▒
▒▒▒▒▒░░░▒▒▒▒▒▒▒░░░▒▒▒▒▒▒
      ░░▒▒▒▒▒▒▒░░░
      ░░▒▒▒▒▒▒▒░░░
▒▒▒▒▒░░░▒▒▒▒▒▒▒░░░▒▒▒▒▒▒
▒▒░  ░▒▒▒▒▒▒▒▒▒▒▒▒░  ░▒▒

3. Thinkertoy Banner Style
bash

go run . "Art" thinkertoy

4. Multi-line Text
bash

go run . "Hello\nWorld" standard

This creates two separate ASCII art blocks with a blank line between them:
text

 _    _          _   _
| |  | |        | | | |
| |__| |   ___  | | | |   ___
|  __  |  / _ \ | | | |  / _ \
| |  | | |  __/ | | | | | (_) |
|_|  |_|  \___| |_| |_|  \___/

__        ___   _       _
\ \      / / \ | |     | |
 \ \ /\ / / _ \| |     | |
  \ V  V / ___ \ |___  | |
   \_/\_/_/   \_\____| |_|

5. Complex Examples
bash

# Multiple words with punctuation
go run . "Hello, World!" shadow

# Numbers and symbols
go run . "123!@#" thinkertoy

# Multi-paragraph text
go run . "Line1\n\nLine2" standard

Banner Style Comparison
Style	Description	Best For
standard	Clean, classic look with spaces and underscores	General purpose, code comments, headers
shadow	3D filled character effect	Eye-catching titles, posters
thinkertoy	Artistic symbolic style	Creative projects, artistic displays
Technical Details
How It Works

    Banner File Loading: The program reads banner files (standard.txt, shadow.txt, thinkertoy.txt) which contain ASCII art representations of characters

    Character Mapping: Each printable ASCII character (32-126) is stored as an 8-line block in the banner file

    Text Processing: Input text is parsed, supporting \n for multi-line output

    ASCII Art Generation: Each character is looked up in the banner and printed horizontally across 8 rows

Banner File Format

    Each character occupies exactly 9 lines in the banner file

    8 lines contain the actual ASCII art

    1 empty line serves as a separator between characters

    Characters are stored in ASCII code order (space first, then !, ", #, etc.)

Supported Characters

The program supports all ASCII characters from code 32 (space) to 126 (tilde):
text

 !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstuvwxyz{|}~

Characters outside this range (emojis, accented letters, etc.) are silently ignored.
Error Handling

The program handles common errors gracefully:

    Missing arguments: Shows usage instructions

    Invalid banner style: Prints error message and exits

    Missing banner file: Displays file not found error

    Empty text input: Exits without output

Complete Usage Examples
bash

# Different banner styles
go run . "ASCII" standard
go run . "ASCII" shadow
go run . "ASCII" thinkertoy

# Multi-line with shadow effect
go run . "Top\nBottom" shadow

# Empty line handling
go run . "First\n\nSecond\n\nThird" standard

# Special characters
go run . "Hello@123#456" thinkertoy

# Long text
go run . "The quick brown fox jumps over the lazy dog" standard

Project Structure
text

ascii-art-generator-go/
├── asciiart.go          # Main program source code
├── standard.txt         # Standard banner font file
├── shadow.txt          # Shadow banner font file
├── thinkertoy.txt      # Thinkertoy banner font file
├── README.md           # This documentation
└── LICENSE             # License file (recommended)

Key Functions
Function	Purpose
loadBanner()	Reads and parses banner file into memory
getCharArt()	Extracts ASCII art for a specific character
printWord()	Prints a word as ASCII art row by row
main()	Entry point, handles arguments and flow control
Limitations

    Only supports printable ASCII characters (32-126)

    Requires banner files to be in the same directory as the executable

    Text input must be wrapped in quotes when containing spaces

    Banner files must follow the exact format (8 lines + 1 separator per character)

Contributing

Contributions are welcome! Here are some areas for improvement:

    Add support for custom banner files

    Implement color output support

    Add more font styles

    Support for Unicode characters

    Create a web API version

    Add alignment options (left, center, right)

    Implement file input/output redirection

To contribute:

    Fork the repository

    Create a feature branch

    Commit your changes

    Push to the branch

    Open a Pull Request

License

This project is open source and available under the MIT License.
Acknowledgments

    Built with Go's standard library packages: fmt, os, strings

    Inspired by classic ASCII art generators and banner tools

    Banner file formats based on common ASCII art standards

Author Contact

VINCENT ONYECHEREM IKENNA

For questions, suggestions, or contributions, please open an issue on GitHub or contact the author directly.
Quick Reference Card
bash

# Standard banner (default)
go run . "TEXT"

# Shadow banner
go run . "TEXT" shadow

# Thinkertoy banner  
go run . "TEXT" thinkertoy

# Multi-line with shadow
go run . "LINE1\nLINE2" shadow

# Show help
go run .

Transform your text into stunning ASCII art! 🎨

Happy coding and enjoy creating beautiful terminal art!
