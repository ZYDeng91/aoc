### Day 9 - Low

No parser. 

#### Part 1

The puzzle input is a single line with 19999 characters, where 10000 of them indicate files and the other 9999 indicate spaces. A naive implementation of the original instructions could work but there is always space to optimize things.

Iterate through the scary noodle of a string, we can stop the count when at some point: the spaces to the left is more than the files to the right. To save time from constructing the disk block by block, we traverse the filesystem by files/spaces. The middle point happen to include a file that is not quite moved by the fragmenter yet part of it is. There will be trouble.

Construct arrays to record the positions and IDs of the files that are to the left (not moved), and the empty spaces to be filled together with the files to be moved in reverse order. Split the middle file between moved and not moved parts. Calculate dot products of the position array over the id array, sum up and we have the result.

My solution is totally not mentally efficient and especially difficult to debug. I was stuck on the result being too low for quite a while. Bruteforce wins another round.

#### Part 2

To me part 2 is easier than part 1, despite I had to rewrite the solution function. This time to better model the files and spaces, I made a new type `File` which contains 3 integer fields: start, size and id. Put them into arrays. Then starting from the right for each file, starting from the left for each space until it goes beyond the file, fit the file into the space if possible. Since each file is processed once there is no need to cover up the moved files' remnants. The spaces though will receive a reduction in size and a right shift of its "start" property. 

Remember to write to the array with a new object.
