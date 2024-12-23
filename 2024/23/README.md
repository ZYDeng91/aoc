### Day 23 - Connected

#### Part 1

Use a map to track all nodes connected to each node. Then for every node, if two of its neighbors are connected, it's a valid trio. Check if any of them starts with 't'. Divide by 3.

#### Part 2

A oneliner with python's networkx. Remember the all-connected set of nodes is a clique.

Because initially I wanted to visualize the network with GraphViz, I wrote some Go to output the .dot file. Afterwards the same function is used to output a list of edges (python code) for networkx to add edges from.
