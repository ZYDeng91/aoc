### Day 20 - Speedrun

Matrix parser. 141x141.

#### Part 1

Dijkstra as always. A bruteforce algorithm (set every '#' to '.' once) works.

#### Part 2

Didn't read that same start/end counts as same cheats, wasted some time modifying the algorithm to skip blocks with a cheat timer.

Read the instructions and the input again, and you just realize there is no dead end, and a cheat must start/end on a position that is seen on the no-cheat loop.

Extract the gScore, the problem becomes simple: select every pair of 2 points from gScore, if one point is reachable from another within the allowed cheat time, and the shortcut will cut the time/score by at least 100, this is a valid route, res++. Remember to compensate for the time used during the cheat.

#### Extra

Curiously, I got the answer right but the example wrong with one off (want 285, got 284). After some debugging I found the starting point has a score of 2. It turns out the 2nd position treated its neighbor-parent as uninitialized because it has a score of 0. This never mattered in shortest path problems, or when with directions I disallowed turning back. If it wasn't for this off-by-one mistake, it would have slipped through many of my copy-pastes unnoticed.

The quick fix is to use the second output of `val, ok := mymap[key]` instead of val to check for uninitialized entries. Guess I was busy implementing the pseudocode and not remembering the cool feature Go has one year ago.

While we are here: 

> `time go run 20.go`

> `go run 20.go  1.35s user 0.09s system 89% cpu 1.595 total`
