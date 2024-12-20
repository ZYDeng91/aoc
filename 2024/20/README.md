### Day 20 - Speedrun

Matrix parser. 141x141.

#### Part 1

Dijkstra as always. A bruteforce algorithm (set every '#' to '.' once) works.

#### Part 2

Didn't read that same start/end counts as same cheats, wasted some time modifying the algorithm to skip blocks with a cheat timer.

Read the instructions and the input again, and you just realize there is no dead end, and a cheat must start/end on a position that is seen on the no-cheat loop.

Extract the gScore, the problem becomes simple: select every pair of 2 points from gScore, if one point is reachable from another within the allowed cheat time, and the shortcut will cut the time/score by at least 100, this is a valid route, res++. Remember to compensate for the time used during the cheat.
