# Markdown Table Of Contents

<p>
    <a href="https://github.com/anotherhadi/markdown-table-of-contents/releases"><img src="https://img.shields.io/github/release/anotherhadi/markdown-table-of-contents.svg" alt="Latest Release"></a>
    <a href="https://pkg.go.dev/github.com/anotherhadi/markdown-table-of-contents?tab=doc"><img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="GoDoc"></a>
    <a href="https://goreportcard.com/report/github.com/anotherhadi/markdown-table-of-contents"><img src="https://goreportcard.com/badge/github.com/anotherhadi/markdown-table-of-contents" alt="GoReportCard"></a>
</p>

Markdown Table Of Contents is a straightforward command-line tool designed to generate a table of contents for your Markdown files based on titles.
This project use the [markdown reader/writer](https://github.com/anotherhadi/markdown) package.

## Features

- **Automatic Generation**: Quickly generate a table of contents for your Markdown files.
- **Title Detection**: Detects headings in your Markdown file and generates links to them in the table of contents.

## How to Use

1. **Installation**:

  ```bash
  go install github.com/anotherhadi/markdown-table-of-contents@latest
  ```

2. **Usage**: Run the executable with your Markdown file as an argument.

  ```bash
  markdown-table-of-contents <markdownfile.md>
  ```

3. **Output**: The tool will generate a table of contents based on the headings in your Markdown file.

**Arguments:**

```txt
  -depth int
        Depth of the table of contents (1,2,...,6). 2 will print only H1 and H2 (default 3)
  -indent string
        Indentation string ('  ','tab') (default "  ")
  -start-by int
        Start by specific header (1,2,...,6). 2 will start by H2 (default 1)
```

## Example

Suppose you have a Markdown file named `example.md` with the following headings:

```md
# Introduction
## Getting Started
### Installation
### Usage
## Features
# Conclusion
```

Running the command:

```bash
markdown-table-of-contents example.md
```

Would generate the following table of contents:

- [Introduction](#introduction)
  - [Getting Started](#getting-started)
    - [Installation](#installation)
    - [Usage](#usage)
  - [Features](#features)
- [Conclusion](#conclusion)

## License

This project is licensed under the [MIT License](LICENSE).

## Todolist

- Replace from file with {table_of_contents} tag or something
- Ignore H1, H2 or anything from the left
