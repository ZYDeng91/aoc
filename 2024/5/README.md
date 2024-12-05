### Day 5 - Rulebook

2 sections. Line by line parser. Split by '|'/','. Atoi.

Data types: `map[[2]int]bool`, `[]int`

#### Part 1

The rules are long and boring. Surely someone can figure out the ranks among the pages, but what's more apparent to me is the rules also tell what orders are invalid.

Make a map as a *rulebook* to support tuple-indexed (or [2]int) lookups, populate it with the rules. For every line of update, Compare every 2-element combination while keeping the order, and check if they violate the rulebook. If so, the line should be deemed incorrect. Sum up the correct lines.

#### Part 2

It might seem unintuitive to sort an array without directly knowing what's bigger or smaller, but we go with what we have here. I wanted to go for bubble sort but there could be hazards to mess things up (e.g. not explicitly ordered pages). I went for a solution which swaps the positions of any invalid pairs until the whole thing is correct (cannot make any more swaps). The relatively short lengths of updates really helped.

Nested loops are confusing. Thankfully there is `goto`.
