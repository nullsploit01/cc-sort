# CC-Sort - Another Command-Line Sorting Utility

CC-Sort is a custom command-line sorting utility developed as part of a coding challenge found here: https://codingchallenges.fyi/challenges/challenge-sort. It supports multiple sorting algorithms including Radix, Merge, Quick, Heap, and Random. This utility allows users to sort the contents of a file with options to treat lines as unique, thereby eliminating duplicates, or to include all lines.

## Features

- Support for Multiple Sorting Algorithms: Choose from Radix, Merge, Quick, Heap, and Random sorting methods to tailor the sorting process to your needs.
- Unique Sorting Capability: Optionally eliminate duplicate lines from the output, allowing for unique-only content sorting.
- Command-line Interface (CLI): Easy-to-use CLI powered by Cobra, providing clear options and commands.
- Efficient Performance: Optimized to handle large files and complex sorting operations efficiently.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- You need to have Go installed on your machine (Go 1.15 or later is recommended).
- You can download and install Go from https://golang.org/dl/.

### Installing

Clone the repository to your local machine:

```bash
git clone https://github.com/nullsploit01/cc-sort.git
cd cc-sort
```

### Building

Compile the project using:

```bash
go build -o ccsort
```

### Testing

Run tests to ensure the application is working correctly:

```bash
go test ./...
```

### Usage

To use the utility, execute the compiled binary with the desired options:

#### Sort a file with a specified algorithm:

```bash
./ccsort --algorithm radix --unique path/to/file.txt
```

### Use the unique flag to remove duplicates:

```bash
./ccsort -u --algorithm quick path/to/file.txt
```

### Examples

Sort a file using the Quick sort algorithm without removing duplicates:

```bash
./ccsort --algorithm quick path/to/file.txt
```

Sort a file using the Random sort algorithm and eliminate duplicate lines:

```bash
./ccsort --unique --algorithm random path/to/file.txt
```

Sort a file using the Random sort algorithm and eliminate duplicate lines:

```bash
./ccsort -u --algorithm random path/to/file.txt
```
