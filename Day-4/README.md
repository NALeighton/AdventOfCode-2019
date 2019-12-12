# Day 4

The premise of day is to determine how many possible passwords exist within the given rule set. All passwords must be with the range of `145852-616942`

## Part 1

For the first part, the rule set was as follows:
* Password must be a 6 digit number
* Digits must either remain the same or increase when read from left to right
    * 111111 - Valid
    * 101111 - Invalid
    * 123456 - Valid
    * 123444 - Valid
* Password must contain a double
    * 112345 - Valid
    * 123456 - Invalid
    * 123344 - Valid
    * 112334 - Valid

### Execution Time

| Average Time (ms) | Standard Deviation (ms) |
| :---: | :---: |
| 0.969711 | 0.0177385986 |

## Part 2

For the second part, the rule set is everything from part 1 plus:
* Repeated digits can only be repeated twice
    * 112233 - Valid
    * 111234 - Invalid
    * 111233 - Valid
    * 111133 - Invalid

### Execution Time

| Average Time (ms) | Standard Deviation (ms) |
| :---: | :---: |
| 31.2656715 | 0.979624262 |