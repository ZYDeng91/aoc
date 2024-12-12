### Day 12 - Wall

Matrix parser. 140x140.

#### Part 1

Pretty simple, for every unvisited cell, get neighbors and label the ones with the same type of plant as visited and add 1 to area, or add 1 to perimeter if the neighbor is of a different type. Sum up the products.

#### Part 2

I spent one hour thinking there is some formula I can derive knowing area, perimeter, free spaces and whatnot. I did find the number of sides is equal to 4 + number of 90 degree angles inside + 2 * number of 90 degree angles outside. But I don't quite feel like doing a flood fill right now, I went on with another idea to analyze the placement of the fences.

Every fence can be represented by its position and facing direction. When adding fences, keep track of their placements. Check the sides of the fence to be potentially connected, if there are no nearby fence to be connected, this will be a new side.

Running against the example shows there is an extra 13 to the result, and it's got to be plant E counted 1 extra side. Output the fence placement when a new side is added, I realized there are conditions when a side has fence(s) placed incompletely, and the sequence of iterations made it to the other end of the side. My solution is to reduce the sides count by one if the newly placed fence is sandwiched by two fences around, which can only happen when the side is already counted twice.

A challenging one today for sure, at least I am not walled, just yet.
