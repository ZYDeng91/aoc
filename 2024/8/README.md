### Day 8 - Sunday

Matrix parser. 50x50.

#### Part 1

Copying a lot from day 6 since many util functions can be reused. Use a map to categorize signals of each frequency(symbol), and for every pair of same frequency find the two antinodes. The antinode positon follows `anti.x, anti.y = 2*ante1.x - ante2.x, 2*ante2.y - ante2.y`, exchange ante1 and ante2 for the other direction.

Exclude the out-of-bound ones. Sum up.

#### Part 2

The main loops are identical to part 1, just replace the find two antinodes part with loops to find all antinodes on two directions. A diff variable would be handy.

Wasted some time debugging only to find out I forgot to change the `if` to `for` when copypasting.

A simpler puzzle on sunday to save us some precious weekend time? Thanks, Eric.
