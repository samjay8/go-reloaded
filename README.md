# go-reloaded


## Overview

go-reloaded is a command-line text transformation tool written in Go. It receives two arguments; an input file containing text that needs modification, and an output file where the corrected text will be written. The tool handles text completion, editing, and auto-correction based on a defined set of rules.

---
## Requirements

+ Go 1.22 or higher

---
## Installation & Usage
Clone the repository and navigate into the project directory:

```bash
git clone https://acad.learn2earn.ng/git/sojetund/go-reloaded.git
cd go-reloaded
```

Run the program with an input file and an output file as arguments:
```bash
go run . input.txt output.txt
```

## Project Structure
```
go-reloaded/
├── main.go
├── go.mod
├── README.md
└── transformation/
    ├── hex_bin.go
    ├── hex_bin_test.go
    ├── cases.go
    ├── cases_test.go
    ├── punctuation.go
    ├── punctuation_test.go
    ├── quotes.go
    ├── quotes_test.go
    ├── article.go
    └── article_test.go
```
---
## Examples:

```bash
$ cat sample.txt
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.

$ go run . sample.txt result.txt

$ cat result.txt
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.
```

```bash
$ cat sample.txt
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.

$ go run . sample.txt result.txt

$ cat result.txt
Simply add 66 and 2 and you will see the result is 68.
```

```bash
$ cat sample.txt
There is no greater agony than bearing a untold story inside you.

$ go run . sample.txt result.txt

$ cat result.txt
There is no greater agony than bearing an untold story inside you.
```

```bash
$ cat sample.txt
Punctuation tests are ... kinda boring ,what do you think ?

$ go run . sample.txt result.txt

$ cat result.txt
Punctuation tests are... kinda boring, what do you think?
```

## Features:

### 1. Hexadecimal to Decimal Conversion — `(hex)`
Every instance of the `(hex)` marker causes the word immediately before it to be converted from hexadecimal (base 16) to its decimal (base 10) equivalent. Both the original word and the marker are replaced by the converted result.

### 2. Binary to Decimal Conversion — `(bin)`
Every instance of the `(bin)` marker causes the word immediately before it to be converted from binary (base 2) to its decimal (base 10) equivalent. Both the original word and the marker are replaced by the converted result.

### 3. Uppercase Conversion — `(up)`
Every instance of the `(up)` marker converts the word immediately before it to uppercase. The marker is then removed from the text.

### 4. Lowercase Conversion — `(low)`
Every instance of the `(low)` marker converts the word immediately before it to lowercase. The marker is then removed from the text.

### 5. Capitalization — `(cap)`
Every instance of the `(cap)` marker converts the word immediately before it so that its first letter is uppercase and all remaining letters are lowercase. The marker is then removed from the text.

### 6. Multiple Word Conversion — `(up, N)` / `(low, N)` / `(cap, N)`
The `(up)`, `(low)`, and `(cap)` markers each accept an optional number. When a number is provided, the transformation is applied to that many words preceding the marker, counting backwards from the marker's position.

### 7. Punctuation Formatting
The punctuation marks `.` `,` `!` `?` `:` and `;` must always be placed directly after the preceding word with no space before them, and exactly one space after them.

When punctuation marks appear consecutively, they are treated as a single group and kept together with no spaces between them. The group as a whole follows the same rule — no space before it, one space after it.

### 8. Single Quote Formatting
Single quote marks ' always appear in pairs. The content between them must have no space immediately after the opening quote and no space immediately before the closing quote. Spaces between words inside the quotes are preserved.

### 9. Article Correction — `a `to `an`
Every instance of the word `a` is replaced with an when the immediately following word begins with a vowel (`a`, `e`, `i`, `o`, `u`) or the letter `h`. The original casing of the article is preserved — `A` becomes `An` and `a` becomes `an`.

## Running Tests

```bash
# Run all tests
cd transformation
go test -v

# Run a specific test

go test -v run TestHexConversion
```

## Author
**Samuel Ojetunde**

