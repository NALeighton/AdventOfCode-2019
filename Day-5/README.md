# Day 5

The premise of day is to upgrade the opcode computer that was built on [day 2](https://github.com/NLeighton/AdventOfCode-2019/tree/master/Day-2) to include more functionality.

## Part 1

For the first part, two more opcodes were added to the valid instruction set.
* `3` - takes in input and stores it in a given location
* `4` - outputs data in a given location

### Execution Time

| Average Time (ms) | Standard Deviation (ms) |
| :---: | :---: |
| 0.6335971 | 0.0172876784 |

## Part 2

For the second part, another four additional opcodes were added.
* `5` - jump if true
* `6` - jump if false
* `7` - less than
* `8` - equals

### Execution Time

| Average Time (ms) | Standard Deviation (ms) |
| :---: | :---: |
| 0.6102898 | 0.0123086502 |