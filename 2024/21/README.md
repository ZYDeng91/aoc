### Day 21 - Overdue

#### Part 1

The fluctuating difficulty has made me underestimate things. Now that we've hit Blackjack days, I can only expect more advanced tomfoolery to come.

I naively thought the sequence length only depends on the x/y difference of the two-button pairs, and the order does not matter as long as they are grouped together. Wrong! Even the example won't let me pass.

I tried to hardcode some ordering decisions from the given example, thinking the empty space could have mattered somehow. I was half right, it still doesn't make it. There is no one-size-fits-all solution, and I have to consider every arrangement that makes sense (i.e. either move x or y first, no inbetweens).

I have to rewrite my solution completely. And merely chaining the strings together takes too much time complexity. The solution is to atomize the strings, into moves from one button to another. The function should return all efficient paths (where up/down or left/right are grouped together) from the start button to the end button. The situation now becomes similar to the Day 11 puzzle.

Recursively call the function for each layer/depth of robots it has, and for each pair of buttons the path has. Terminate at depth 1, where it should return the length of the shortest sequence of paths between input buttons. (alternatively, terminate at depth 0 where it should return 1)

#### Part 2

Memoization again. The little `var cache map[State]int` does it all.

The "big number" (featuring 10 '9's) that starts as an uninitialized minimum when counting minimum happens to be not enough for the bigger numbers of sequence lengths, which I found out because of the number of 9s in the outputed answer. I had to resort to the max int64, even though one more 9 would have made it.
