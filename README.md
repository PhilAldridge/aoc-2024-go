# 🎄 Advent of Code 2024 🎄

Working through problems for [Advent of Code 2024](https://adventofcode.com/2024).

Fork of template repo: github.com/kylehoehns/aoc-2023-go/ 

Using [Go](https://go.dev/)

## Setup

1. Install Go from [here](https://golang.org/doc/install).
2. Run the following script to install dependencies.
    ```shell
    go mod tidy
    ```

## Running

Run the following command to test a given day's problem.

```shell
go test ./puzzles/day00
```


Run the following to run the given puzzle input for the day.

```shell
go run ./puzzles/day00/main.go
```
or
```shell
./run-day.ps1 00
```

## Template

To create a template of files needed for a new day's puzzle, run the following command.

```shell
./scripts/create-day 01
```

This will create a new folder named `day01` pre-created with files for the main code, test code, and input files.

## Solution Times
### Day 1
Part 1: 12.6985ms,
Part 2: 522.1µs
### Day 2
Part 1: 15.7083ms,
Part 2: 1.1514ms
### Day 3
Part 1: 14.5766ms,
Part 2: 608.4µs
### Day 4
Part 1: 23.1866ms,
Part 2: 538.6µs
### Day 5
Part 1: 12.8562ms,
Part 2: 2.9952ms
### Day 6
Part 1: 20.5151ms,
Part 2: 800.537ms
### Day 7
Part 1: 17.2347ms,
Part 2: 244.6467ms
### Day 8
Part 1: 22.2637ms
Part 2: 627.3µs