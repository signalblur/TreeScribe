# TreeScribe

TreeScribe is a simple, lightweight, and minimal tool that visually maps out your directory structures, providing a clear and concise view of your project hierarchy. Perfect for developers who appreciate straightforward file organization.

## Features

- Outputs a visual diagram of folder and filenames.
- Easy-to-use command-line interface.
- Uses only the Go standard library.
- Statically linked and compiled for simplicity and security.

## Installation

First, ensure you have [Go](https://golang.org/doc/install) installed on your system. Then, compile TreeScribe from the source code:

```sh
go build -o TreeScribe main.go
```

## Usage

TreeScribe provides a clear view of your directory structure through the command line. Here are the available options:

```
./TreeScribe [options]

Options:
  -h, --help       Show help
  -p, --path       Path to the directory
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

## Output

For a given directory structure TreeScribe will output:


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