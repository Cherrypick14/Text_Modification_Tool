
# Text-Modification-Tool

This project is a command-line utility designed to manipulate text files according to specified functions. Upon execution, the program reads an input text file (specified as the first argument) and applies a series of text transformation functions to the content. The transformed text is then written to an output file (specified as the second argument).

## Dependencies

This project relies on an external package, reloaded/functions, which provides the implementation for the text manipulation functions.


## Features

- Text Transformation: The program offers a range of text transformation functions, including capitalization, changing case (upper and lower), removing vowels, apostrophes, and punctuation marks, and converting hexadecimal and binary numbers to decimal.

- Command-Line Interface (CLI): It operates as a command-line utility, allowing users to specify input and output files as command-line arguments, enhancing usability and integration into scripts or automation workflows.

- Error Handling: The program includes basic error handling for file input/output operations and reports errors to the console, ensuring robustness and user awareness in case of issues like missing files or permission errors.

- Modular Design: The text transformation functions are implemented in a separate package (reloaded/functions), promoting modularity, code organization, and potential reuse in other projects or contexts.

- Scalability: The codebase can easily accommodate additional text manipulation functions or enhancements, fostering scalability and extensibility to support evolving requirements or user needs in the future.


## Usage

Clone the project

```bash
 
  git clone https://github.com/Cherrypick14/Text_Modification_Tool
```

Go to the project directory

```bash
  cd Text_Modification_Tool
```

Run the program

```bash
  go run main.go <sample.txt> <result.txt>
```

## Running Tests

To run tests, run the following command

```bash
  go test -v
```


## Authors

- [@Cherrypick14](https://github.com/Cherrypick14)



