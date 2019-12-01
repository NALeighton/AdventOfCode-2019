# Advent of Code 2019

Advent of Code is an annual language agnostic coding challenge consisting of 25 2-Part challenges. 1 challenge is released each day starting on December 1st and ending on December 25th (Hence the name, Advent of Code). The 2019 challenges can be found [here](https://adventofcode.com) as they are released.

## Language of choice

This year I am using [Go](https://golang.org) as my language of choice. This choice was made because I have been looking for an excuse to write in Go (As far as excuses go, that is pretty lame).

## Execution Times

A part of the fun of Advent of Code is creating not only a working solution, but one that also executes in a decent amount of time. My method for timing is as follows: `The timer starts on execution and stops once the result has been returned. The times from 10 runs are then averaged to smooth out the variability that comes from processor load. All times are measured in milliseconds`. All times are taken using the compiled program.

| Challenge | Average Time (ms) | Standard Deviation (ms) |
| :---: | :---: | :---: |
| Day 1 - Part 1 | 0.1203382 | 0.0152516995 |
| Day 1 - Part 2 | 0.1337176 | 0.0080913641 |

## Repository Structure

All solutions are self-contained within directories marked with the day of the challenge. Any special dependencies are declared in that day's README.

## License

All solutions are licensed under [GNU AGPLv3](https://choosealicense.com/licenses/agpl-3.0/).