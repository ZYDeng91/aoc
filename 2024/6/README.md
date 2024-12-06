### Day 6 - Stall

Matrix parser. 130x130.

#### Part 1

A classic idea that is seen in last year's puzzles. Loop over steps until the next step is out of the map: `new_x, new_y = x + dir.x, y + dir.y`, turn when '#' is reached. Sum up the visited matrix.

#### Part 2

The input size is still small and bruteforcing is arrangeable. There are places to optimize the bruteforcing process, and I took the follow:

- Obstacle placement: Only consider visited spots to add obstacles

- Breakout of infinite loops: If the same state (position + dir) is reached again, the guard must be looping

I did simplify the breakout condition to be if the same position is visited >4 times. Like the Birthday Problem, with >4 times visited, there must be at least 2 times where the directions are the same.

That being said, it really does not take long even if no optimization is applied. I guess this is one of the perks of using a compiled language unlike python. It only took 1 second if the infinite loops are broken at the 100000th step.

`go run 6.go  1.03s user 0.08s system 114% cpu 0.962 total`
