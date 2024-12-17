### Day 17 - Reverse

The great filter is here. This is it.

#### Part 1

I had to read the problem statement over and over to understand what's going on. To simplify it, the registers define 3 variables, and the program is an array that should be processed in pairs. Each pair contains an opcode and an operand, and depending on the opcode the operand can mean different things.

Use switches to handle the opcodes and combo operands, and a jump is merely a change to the looping counter/instruction pointer.

#### Part 2

What looks like a reverse engineering problem (crackme) is a binary manipulation problem in disguise, after drawing the pseudocode what the program does becomes clear.

Flowing through each every opcode-operand pair (based on my input):

- b = a % 8 => b equals to the 3 LSBs of a

- b = 2 ^ b

- c = a / pow(2, b)

- a = a / 8 => right shift a by 3 bits

- b = b ^ c

- b = b ^ 7

- outputs 3 LSB of b

- jump to beginning if a is not 0

Since the output's length is 16, the loop is restarted 15 times. Therefore pow(8,15) < a < pow(8,16).

Chain the operations, the program can be reduced as:

- b = popLSB(a, 3)

- output = b ^ 5 ^ a[-b^2-3:-b^2]

Simple as it looks, I spent a lot time reversing the operation from the known output. Without seeing the output I assumed there will be multiple answers, some of which contradicts itself at some point, and to get the lowest one I should compare them after finding them all. This is wrong.

The last output correspond to the 3 MSBs of A. A 0-7 loop breaking on the first catch can find the smallest B, which easily converts to 3 bits of A. Everything in front of the MSB of A is 0. Going back to the front of the output array, every variable is known from there.
