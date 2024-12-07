### Day 7 - Numbers

Line-by-line parser. Sum.

#### Part 1

I probably shot myself in the foot for not going for a total bruteforce strategy since the numbers on the right are too few to have an impact on performance.

I imagined the problem as a 2-choice decision making at every intersection, starting from the right of the equation. If the left number is divisible by the last element of the right list, branch into 2 possible operators + and x, else + is the only option. Recursively call the function itself until the right array has only one element left, and compare to the left.

#### Part 2

Similarly implement the reverse operation of the new concat operator, check if concat is possible and go backwards.

A significant mistake I made is ignoring the sub-conditions on 3-way branches. When filtering out options the control flow should be hierarchical not parallel, neglecting the check for multiplication when the check for concatenation has passed is unwise.

I personally blame the compiler for converting division results to int automatically, and the examples for not covering cases where I would fail. Ultimately bruteforcing left-to-right would solve much logical confusions, mea culpa.
