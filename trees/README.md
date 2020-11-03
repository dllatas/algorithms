# Trees

### Kind of trees
- Binary tree (BT): each node has up to 2 children.
- Binary search tree (BST): binary tree which has this ordering: all left descendents are lower than all right descendents

#### Problems

1 BST with minimal height

A BST's minimal height is achieved when the tree is balanced.

Given that the input is an array of ints, and we want a balanced BST, the perfect root candidate is the median element of the sorted array.

How to keep the BST balanced while performing the insert? Always insert the median of the remaining sorted input array.
