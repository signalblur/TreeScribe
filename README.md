# TreeScribe

TreeScribe is a simple, lightweight, and minimal tool that visually maps out your directory structures, providing a clear and concise view of your project hierarchy. Perfect for developers who appreciate straightforward file organization.

## Features

- Outputs a visual diagram of folder and filenames.
- Easy-to-use command-line interface.
- Uses only the Go standard library.
- Statically linked and compiled for simplicity and security.

## Installation

First, ensure you have [Go](https://golang.org/doc/install) installed on your system. Then, compile TreeScribe from the source code:

```
go build -o TreeScribe main.go
```

Move the binary to a directory in your PATH. On Linux and macOS, a common directory for user-installed binaries is /usr/local/bin.

```
sudo mv TreeScribe /usr/local/bin/
```

## Usage

TreeScribe provides a clear view of your directory structure through the command line. Here are the available options:

```
./TreeScribe [options]

Options:
  -h, --help             Show help
  -p, --path             Path to the directory
  -i, --include-hidden   Include hidden folders
```

## Examples

Display the help message:

```
./TreeScribe -h
```

or

```
./TreeScribe --help
```

Print the directory structure of a specified path:

```
./TreeScribe -p /path/to/directory
```

or

```
./TreeScribe --path /path/to/directory
```

Include hidden folders (e.g., .git):

```
./TreeScribe -p /path/to/directory -i
```

or

```
./TreeScribe --path /path/to/directory --include-hidden
```

## Output

TreeScribe will output:

```
myApp/
    ├── config/
    │   └── config.yml
    ├── static/
    │   ├── templates/
    │   │   └── derp.html
    │   └── css/
    │       └── styles.css
```

## Adding an Alias

To make it easier to run TreeScribe, you can create an alias ts in your shell configuration file.

### On macOS and Linux

Open your zsh configuration file (.zshrc) or bash configuration file (.bashrc or .bash_profile):

```
nano ~/.zshrc
```

```
nano ~/.bashrc
```

Add the alias at the end of the file:

```
alias ts='TreeScribe'
```

Save the file and exit the editor.

Reload your shell configuration to apply the changes:

```
source ~/.zshrc
```

or

```
source ~/.bashrc
```

Verify the alias by running:

```
ts -p /path/to/directory
```