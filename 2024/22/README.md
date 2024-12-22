### Day 22 - Monkey Business

Line parser.

#### Part 1

An insignificant day towards the end - literally do what the instruction says, time is not a problem.

#### Part 2

The example changed in part 2, take note of that.

Generate a sequence of the prices (secret%10) and iterate through it. Keep a `[4]int` tracking the 4 consecutive price changes, slide FIFO. For every never seen 4 price changes, add the price at the cursor position to a `map[[4]int]int`. Repeat on every other buyer, use the same map.

Find the highest value in the map. That's it.
