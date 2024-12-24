### Day 24 - Good Night

#### Part 1

Store the variables in a map, process the logic gates if both inputs are present, otherwise throw it into a queue (q). Loop with the unprocessed gates, until there is no unprocessed ones left.

For every entry in the map starting with 'z', add `pow(2, z_index)*val` to the result.

#### Part 2 - analysis

Bruteforcing does not work. Some manual work, on the other hand...

The spirit of solving a system of equations is to reduce the number of unknowns. I can reduce the entire system into only x, y and z's. Even though the output is long and stinky.

Reading the first few lines of the output should give you an idea about how the system works. Write things down on paper, I find the binary program consisting of several fixed structures, and I started to name them:

- x01 XOR y01 => z01-, ignoring previous bits, the result of bitwise add

- x01 AND y01 => z01+, ignoring previous bits, the overflow flag of the bitwise add

And advanced ones:

- z01+ OR (z00+ AND z01-) => z01\*, overflow flag of the bitwise add, taking previous overflow flag into consideration. 

which later becomes z02+ OR (z01\* AND z02-) => z02\*, to use the real overflow flag.

To make things easier, there is also:

- z02- AND z01\* => z02!, the overflow flag caused by previous bits and the resulting bit of bitwise add, where z02+ OR z02! => z02\*. This also explains the usage of OR since only one of the resulting bit or the overflow flag of bitwise add can be 1.

And finally, z02- XOR z01\* => z02, the true result, a.k.a. the output bit.

#### Part 2 - execution

I figured it would be easier to take things by hand instead of writing code that I cannot fathom, and in the end it would require manual inspections anyway. Vim's regex search and replace functionality rocks for its support of extended syntax (especially backreferences), so I can save some work from writing regex for every new index. It's also really comforting to watch the verbose lines getting shortened into a single XOR.

There is one thing that could have been done differently: by printing out the binary representation of the output and the expected output, you see exactly 4 turning points (where the bits start to align or misalign with the expected binary). Those are the problematic ones.

Anyhow, look out for anomalies and you'd find the swapped ones being exposed easily. Find which variable is placed in whose place, trace to the underlying wires. A simple `python> print(",".join(sorted(wires)))` to cross the finish line.

I'm glad that I pulled out all those by myself without any external help, while it's not the most efficient solution for certain. I've seen my improvements since last year. 2024 has been a great year. Tonight will be a good night.
