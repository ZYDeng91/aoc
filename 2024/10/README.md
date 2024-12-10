### Day 10 - Zero

Matrix parser. 42x42.

#### Part 1 & 2

I did part 2 unknowingly having misread the part 1 requirements, I'll put them together here.

It's way past the time when I got smashed by dynamic programming. For part 1 a global variable to track visited "9"s would be helpful. For every "0" in the matrix, reset the visited map and start the DP function with its position and a state (byte) '0'.

Terminate condition is when the state is '9', set the position to true in the visited map, and return 1 for one route finished in part 2.

For every direction of the current location, excluding out-of-bound ones, check if the corresponding value in the matrix is exactly `state + 1`, i.e. height exactly one higher than current. 

If yes, recursively call the DP function with the new location and the new state, adding the return value to part 2 result.

Count the visited map, sum the returned values.
