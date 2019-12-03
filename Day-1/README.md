# Day 1

The premise of day is to calculate the amount of fuel required to launch a rocket given the mass of all of the modules. The equation for the calculation being: `⌊ mass / 3 ⌋ - 2`.

## Part 1

For the first part, only the total fuel was required.

### Execution Time

| Average Time (ms) | Standard Deviation (ms) |
| :---: | :---: |
| 0.1203382 | 0.0152516995 |

## Part 2

For the second part, it was discovered that fuel requires fuel which requires fuel or in otherwords fuel fueling fuel fueling fueling fuel (I'm not sorry for doing that, it had to be done).

In order to arrive at the correct answer, the fuel calcuation had to be ran until the total mass including fuel was accounted for by the appropriate amount of fuel.
> At first, a module of mass 1969 requires 654 fuel. Then, this fuel requires 216 more fuel (654 / 3 - 2). 216 then requires 70 more fuel, which requires 21 fuel, which requires 5 fuel, which requires no further fuel. So, the total fuel required for a module of mass 1969 is 654 + 216 + 70 + 21 + 5 = 966.

### Execution Time

| Average Time (ms) | Standard Deviation (ms) |
| :---: | :---: |
| 0.1337176 | 0.0080913641 |