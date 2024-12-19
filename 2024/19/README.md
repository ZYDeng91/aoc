### Day 19 - Staircase

Line parser. Split line 1 by ", ". Skip line 2.

#### Part 1

This is quite reminiscent of the staircase problem: you can take from 1 to k stairs at a time, return the number of ways to climb a staircase with certain height.

Take the length of the longest pattern as k. Get the first 1 to k chars of the design, if the string slice is a valid pattern, then proceed to recursively call the function itself with the remaining chars of the string. Terminates when the string is exactly one pattern (res++) or unable to continue (impossible).

#### Part 2

Nothing much different except the big numbers. Memoization does it.
