### Day 16 - Pathfinder

Matrix parser. 140x140.

#### Part 1

A classic pathfinding problem, solvable by Dijkstra. I learned from my mistakes I made last year.

#### Part 2

With the gScore from part1, backtrack from the end tile by reversing the path.

As always, I can't finish one day without mistakes - I set the starting direction to (0,0) when it's stated the Reindeer start facing *East*. It slipped through part 1 somehow, and in part 2 I lost track of one branch from the starting tile because of it.

The good thing is, inspired from day 14, I now understand the importance of visualization in debugging. I printed out the best paths calculated on the map, and by manual inspection I found the missing part easily.
