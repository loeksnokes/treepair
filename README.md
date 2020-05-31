# treepair
Implements a Go package treepair using our Go package prefcode.  treepair supports standard R. Thompson's group calculations for the groups F, T, and V.

From treepair.go
======================================================================
TreePair: An interface built specifically for the type treePair.

treePair::

Has:
Two prefix codes with the same base alphabet and same cardinality.  One
called the range and one called the domain.  Either side can carry a permutation,
as the PrefixCode interface inherently assocites a bijection from the code to a set
{0, 1, ..., k-1} while the words in the prefix code have a natural order.

Can:
1) Expand both trees at corresponding location in domain or range.
2) Permute domain or range permutation
3) ResetLabels sets the
   domain tree labels as 0 1 ... k-1 corresponding to its dictionary order and
   modifies range tree so the initial prefix map is unchanged.
4) Minimise and Minimize (same function but for British English or American English.)
5) Multiply tree pairs based off of same alphabet.
6) Invert an element.
6) Detect if the element is in F, T, or V.
7) Initialise (trivial permutation elt) from a list of expansions in D/R
8) Initialise from DFS notation representation string: e.g. "{11000,10100,1 2 0}"" (an elt of T)
9) Initialise from Full representation string: e.g. "D: [00 0], [01 1], [1 2], R: [0 1], [10 2], [11 0]"
10) Return domain/range permutations (natural permutation from prefix code in
	dictionary order to the numeric labels of leaves)
