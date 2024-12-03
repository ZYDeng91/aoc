### Day 3 - Parsing

Line-by-line parser. Split by `,`. Atoi.

#### Pt 1 & 2

The programming side is brutally straightforward: slice the `mul(a,b)` string with [4:-1] and get the numbers by splitting then atoi. Even with pt2 there's only 2 more cases of exact string matches. All it's left is a times b, no extra function required.

The catch of today's puzzle is to filter out the garbage, and I find Regex extremely efficient with tasks like this. Last year I experimented on solving puzzles with vim only, even though I have to stop (due to the lack of brain power) after a few ones, I realized the power of data preprocessing with vim commands.

Here are the commands to remove the useless characters:

Pt 1: `:%s/\vmul\(\d+,\d+\)/\0\r/g` `<enter>` `:%s/\v.*(mul\(\d+,\d+\))/\1/g`

Explanation: Find every Regex match of `mul(\d+,\d+)`, add a newline after each one, then remove the stuff in front of it. Could always append an anchor `$` to locate line ends.

I wasted plenty of time for Pt 2 thinking how you can inverse the Regex matches, but there is always the '\|' aka OR to do the same as Pt 1. Something like `do()\|don't()\|\vmul\(\d+,\d+\)` will do. I kept getting confused at the magic and backslashes.

There is also a grep/sed/awk line I won't go into.

#### Good Stuff

A little extra exercise on how I'd imagine a Vim-only solution will be (continueing from preprocessed input):

`:%s/\vmul\((\d+),(\d+)\)/+\1*\2/g` is enough for part 1, and

`:%s/do()/)+1*(0/g` and `:%s/don't()/)+0*(0/g` to convert dos and don'ts into multipliers

`:%s/\v\n//g` and `:%s/^/(0/` and `:%s/$/)/` to clean up

Arithmetics time. I'll show keystrokes to use the arithmetic/expression register with the default/unamed register.

`^<Shift-c><Ctrl-r>=<Ctrl-r>"<Enter>`
