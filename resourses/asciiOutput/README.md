# ASCII Art Generator

This project is an ASCII Art Generator that converts input text into ASCII art using different fonts and optional color coding and either outputs  on the terminal or on a specific file of the user's choosing. 

## Prerequisites

This program was created entirely with Golang programming language so ensure your workstation supports golang version 1.18 or later before cloning the repository.

You can install golang from there website depending on the type of operating system your using at https://go.dev/doc/install and follow instructions on how to set-up

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. Clone the repository:
    ```sh
    git clone  https://github.com/jesee-kuya/ascii-art-output.git
    ```

2. Navigate to the project directory:
    ```sh
    cd ascii-art-output
    ```

3. Build the project:
    ```sh
    go build -o ascii-art-output
    ```

## Usage

Run the binary with the required arguments:

```sh
./ascii-art-output -output=<text.txt>  "Your text here" <banner_file>
```
or this other optionl commands which are relevant to this program

this for base functionality
```sh
./ascii-art-output "your text here"
```

or this that specifies bannerfiles
```sh
./ascii-art-output "your text here" <banner_file>
```

or this that colors 
```sh
./ascii-art-output -color=<color|color-code> "your text here" 
```

or letters to be colored 
```sh
./ascii-art-output -color=<color|color-code> <letters to color> "your text here" 
```

All this are possible and can be combined with the output flag to output whatever the user's output he|she desires to the chosen text file
i.e
```sh
./ascii-art-ouput -output=<text.txt> --color=<color> "Your text file" <banner_file>

```
 
### Parameters

- **-banner_file**: The font file to use (default: `standard`).
- **-color**: The color to use for the text.
- **-letters**: The specific letters to color.
- **"Your text here"**: The text you want to convert to ASCII art.
- **-output**: The flag for the file the user chooses


# **Examples**
## **Example 1: Basic Usage**
```go
./ascii-art-output -output=test.txt "Hello, World!" thinkertoy

cat -e test.txt
                                                    $
o  o     o o             o       o         o    o o $
|  |     | |             |       |         |    | | $
O--O o-o | | o-o         o   o   o o-o o-o |  o-O o $
|  | |-' | | | |          \ / \ /  | | |   | |  |   $
o  o o-o o o o-o o         o   o   o-o o   o  o-o O $
                 |                                  $
                                                    $

```


# **Contributing**
We welcome contributions to improve the ASCII Art Generator. If you find a bug or have a feature request, please open an issue. If you would like to contribute code, please fork the repository and submit a pull request.

### **Steps to Contribute**

1. **Fork the repository.**
```sh
   https://github.com/jesee-kuya/ascii-art-output.git
```
2. **Create a new branch**:
```sh
   git checkout -b feature-branch
```
3. **Make your changes.**
4. **Commit your changes**
```sh
git commit -am 'Add new feature
```
5. **Push to the branch**
```sh
git push origin feature-branch
```
5. **Open a pull request**

## License

This project is licensed under the MIT License. See the [LICENSE]( https://github.com/jesee-kuya/ascii-art-color.git/main/LICENSE)  file for details.


Brought to you by [Jesee Kuya](https://github.com/jesee-kuya) , [James Muchiri](https://github.com/j1mmy7z7) and [Fena Onditi][def]


[def]: https://github.com/konditi1
