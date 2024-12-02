### Day 2 - One Off

Line-by-line parser. Split by \s. Atoi loop.

#### Pt 1

Loop thru and check delta's range.

#### Pt 2

Instead of returning when unsafe is found, test out the array with suspects removed. The suspects being either the left or the right element of the unsafe delta. If either passes, return 1.

1st mistake: Shallow copy. It was only until I print out the original array did I realize them being disfigured after the appends. I quickly wrote my own `func cp()` but there is slices.Clone() available.

2nd mistake: Off by one. The first possible moment of detecting an unsafe is between position 1 and 2, while it could be pos 0 causing the bad level. I hardcoded the monotonicity by pos 1 - pos 0, fully trusting the validness of pos 0. As an afterthought, bruteforcing (try removing every single level) is more apparent logically. Lesson learned is to think about edge cases when using hardcoded values, or resort to bruteforce if typing/thinking speed is the shortest plate of the barrel.
