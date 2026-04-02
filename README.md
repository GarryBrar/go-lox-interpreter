# lox-go

A tree-walking interpreter for the Lox programming language, written in Go. 

Implemented following Robert Nystrom's [Crafting Interpreters](https://craftinginterpreters.com).

## Requirements

- Go 1.22 or later

## Build

```sh
git clone https://github.com/yourname/lox-go
cd lox-go
go build ./...
```

## Usage

Run a source file:

```sh
./lox script.lox
```

Start the REPL:

```sh
./lox
```

## Example

```lox
fun greet(name) {
  print "Hello, " + name + "!";
}

greet("world");
```
