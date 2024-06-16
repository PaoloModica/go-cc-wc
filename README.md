# Go `wc` tool

## Overview 

This project aims at building Unix `wc` tool using `Go`.

## Usage

### Build

```bash
go build -o ccwc cmd/cli/main.go
```

### Run

```bash
./ccwc [option] [file]
```

Example

```bash
$ ./ccwc -l test.txt
[7146] test.txt
```

## Acknowledgement

Coding Challenge #1: Build Your Own `wc` tool.
Go check out John Crickett's [Coding Challenges newsletter](https://codingchallenges.fyi) for more inspiring challenges.
