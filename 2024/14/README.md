### Day 14 - Automaton

Line parser. Split by " v=" then ",".

#### Part 1

Pretty simple use of modulo.

#### Part 2

An epitome of human creation in terms of computing puzzles. You aren't told what to expect but are supposed to figure out by yourself.

I was so overwhelmed by such little information, shutting down my monkey brain for quite a while until I recouped with the idea that a picture of a Christmas tree must have some straight line feature. It turned out to be true - what a nice tree with a frame around it.

My guess is the answer wouldn't be too big, and we aren't supposed to watch the simulated field with a 1-second interval either. Going back to the basics, for every second print out the simulated state. There are two ways, to create a new empty state and fill it every tick or to operate on the same variable without refreshing it. I chose the latter to save some space complexity, and for that I have to model the robots as a new data type (or it could be a new class). For every tick, for every bot, remove 1 from the old position, move it by adding velocity and modulo board size, add 1 to the new position. Each tick prints out the time and the state, with 'X' representing a cell with robots in it, and ' ' or blankspace for an empty cell.

I even pulled out some ANSI escape code magic wanting to refresh the display, only to find I don't really need it. I wrote the result to a file (`go run 14.go > 14.out`), and searched a sequence of several 'X's throughout the output, until it hit the one and only.

With the AoC submit feedback, one can roughly gauge the magnitude of the answer - since it tells too high or too low from a faraway wrong answer.
