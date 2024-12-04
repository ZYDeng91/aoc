### Day 4 - Enum

Matrix parser. 140x140.

#### Pt 1

`wc 4.in` tells the matrix is 140x140, which is rather small and bruteforcing works. Loop through x, y of the matrix and find the first letter 'X' of the word "XMAS". Starting from there, find all 4-letter words for all possible directions (try-except would look much nicer here), compare with the string literal "XMAS" and res+=1 for each match.

I went for a more generalized solution (works on other words) but hardcoding could save some time here.

#### Pt 2

Similarly, this time find all occurences of 'A' that allows a X-shape around it (i.e. not on the border of the matrix). Since the task asks for only a tiny cross, a more specialized/hardcoded solution looks more elegant to me. Check if the surrounding chars are 2 'M's and 2 'S's, and make sure there is no 'MAM' or 'SAS' by checking two diagonals. Orders do not matter here and the alphabet is limited, comparing the product of the rune values will suffice.
