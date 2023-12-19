# Advent Of Code 2023

Archiving my ideas and prompts about the yearly event [Advent Of Code](https://adventofcode.com/).

2023 is my 1st year into AoC, and this is my 1st hands-on adventure with Golang (I don't have what it takes to play competitive). Not gonna clean up the code since they are reflective of my thought process, instead I will try to provide reviews here. Enjoy!

[13](#day-13) - [14](#day-14) - [15](#day-15)

## Personal Times

      --------Part 1--------   --------Part 2--------
Day       Time   Rank  Score       Time   Rank  Score
 19   01:19:39   4844      0   02:49:26   3363      0
 18   01:17:09   4063      0   02:26:53   3054      0
 17   05:22:43   5525      0   05:33:30   4858      0
 16   00:43:46   2669      0   00:58:13   2767      0
 15   00:14:29   4770      0   00:58:07   4901      0
 14   00:31:26   4782      0   02:00:03   4833      0
 13   01:10:48   5218      0   01:54:39   4973      0
 12   04:37:53  11749      0   06:07:31   5475      0
 11   01:03:00   7219      0   01:05:38   5622      0
 10   00:41:27   2425      0   03:28:37   4726      0
  9   00:37:24   6840      0   00:39:56   5951      0
  8   00:27:39   7573      0   00:53:07   4099      0
  7   02:08:48  12286      0   02:50:50  11084      0
  6   00:38:44  10406      0   00:41:25   9254      0
  5   01:14:31   9490      0   01:38:58   3384      0
  4   01:19:25  16225      0   01:30:36  12227      0
  3   03:27:57  17211      0   03:47:21  14338      0
  2   00:52:11  10583      0   01:51:29  14549      0
  1   00:31:30  10214      0   01:03:41   7145      0

### Day 13

#### Part 1

The problem statement can be a bit confusing at some point, my interpretation is: 

> For each pattern if we trim off several rows/cols from a SINGLE side, the shape left is symmetrical, find the line of symmetry, be it horizontal or vertical. The rest is simple arithmetics.

I don't have in mind any refined way to find symmetry. Relying on the assumption that all symmetric shapes are of even length/width, the symmetry center must have 2 identical lines beside it. 

We start by iterating to find such candidates (which fulfills mat[i-1]==mat[i]), then spread out to further lines until one edge is reached (mat[i-1-j]==mat[i+j]). Problem solved.

I am being lazy here, some shortcuts I took include: 1. direct line/string comparison, instead of element by element (this resulted in some computational inefficiency in part2) 2. transpose the matrix for vertical cases, just to reuse the horizontal code

It also turns out that one shape can only be either horizontal or vertically symmetrical.

#### Part 2

I think here it's better to modify part 1 such that it gives a leeway of 1 character when checking symmetry, and deduce the new result. My part 1 implementation happens to be not ready for this though, so I went with bruteforcing i.e. flipping each char. Luckily the numbers are not big enough to stop me.

There are 2 potential smudges, symmetrical over the new symmetry line. Either one of them works, yet I computed both then took a /2.

### Day 14

#### Part 1

'O's falling towards north. My immediate reaction is we can take '#'s as receivers, and reconstruct 'O' values using arithmetic series (start+end)\*n/2. Tranverse the matrix and count how many 'O's falls to each '#' with a map/dict. The border also is also considered as a line of '#'s. Sum up all the arithmetic series and we have the answer.

#### Part 2

Well I definitely didn't see this coming. My implementation does not model the final positition of each rock and reconstructing is inefficient. But then, ain't no way we are supposed to rotate *1000000000 cycles* right? I can assume the rotations reach an equilibrium, only repeat itself after a certain time is reached.

I went on to complete the code for other directions and the reconstruction, it is painful I need to rewrite the logics for 3 more times. At least I was quick to notice the question asks for cycles, meaning 1 of each direction.

Printing out all matrices after each cycle will be too verbose, instead we inspect the part 1 load value. Unsuprisingly it starts repeating after 180 cycles, with a period equal to today's date - 14. Now we can simply get the result after (dumb number-180)%14+180 cycles.

### Day 15

#### Part 1

Parse and HASH, nothing much. Rune using ASCII is convinient.

#### Part 2

Interesting to see this part 2 deviates so far from part 1, leaving the latter only as a helper function. The HASHMAP hint is strong. Yet I still got lost on the box indexing, only to find out the box index is equal to the part 1 hash of the label. Out of 256 they really could use better examples.

There are now '-' and '=', two types of operations. My approach is to use 2 maps, one for focal length indexed by label (map\[string\]int), and the other for in-box slots indexed by box number (map\[int\]\[\]string). For '-' we can simply take out the label from the slot and reset its focal length, and for '=' we can modify the focal length, and append to the box if not present (indicated by a 0 focal length).

Loop through the boxes' content, multiply and sum up.
