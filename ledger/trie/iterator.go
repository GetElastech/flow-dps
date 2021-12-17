// Copyright 2021 Optakt Labs OÜ
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy of
// the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.

package trie

import (
	"github.com/onflow/flow-go/ledger"
)

// NodeIterator is an iterator over the nodes in a trie.
// It guarantees a DESCENDANTS-FIRST-RELATIONSHIP in the sequence of nodes it generates:
//   * Consider the sequence of nodes, in the order they are generated by NodeIterator.
//     Let `node[k]` denote the node with index `k` in this sequence.
//   * Descendents-First-Relationship means that for any `node[k]`, all its descendents
//     have indices strictly smaller than `k` in the iterator's sequence.
// The Descendents-First-Relationship has the following important property:
// When re-building the Trie from the sequence of nodes, one can build the trie on the fly,
// as for each node, the children have been previously encountered.
type NodeIterator struct {
	// The node iterator is initialized with an empty stack and the trie's root node assigned
	// to unprocessedRoot. On the FIRST call of Next(), the NodeIterator traverses the trie
	// starting from the root in a depth-first search (DFS) order (prioritizing the left child
	// over the right, when descending). It pushes the nodes it encounters on the stack,
	// until it hits a leaf node (which then forms the head of the stack).
	// On each subsequent call of Next(), the NodeIterator pops the head of the stack.
	// Let `n` be the node which was popped from the stack.
	// If `n` has a parent, denominated as `p`, the parent is now the head of the stack.
	// Parent `p` can either have one or two children.
	//   * If `p` has only one child, there is no other child of `p` to enumerate.
	//   * If `p` has two children:
	//       - if `n` is the left child, we haven't searched through `p.RightChild()`
	//         (as priority is given to the left child)
	//         => we search p.RightChild() and push nodes in DFS manner on the stack
	//            until we hit the first leaf node again
	// By induction, it follows that the head of the stack always contains a node,
	// whose descendents have already been recalled:
	//   * after the initial call of Next(), the head of the stack is a leaf node, which has
	//	   no children, it can be recalled without restriction.
	//   * When popping node `n` from the stack, its parent `p` (if it exists) is now the
	//     head of the stack.
	//     - If `p` has only one child, this child is must be `n`.
	//       Therefore, by recalling `n`, we have recalled all ancestors of `p`.
	//     - If `n` is the right child, we have already searched through all of `p`
	//       descendents (as the `p.LeftChild` must have been searched before).
	//       Therefore, by recalling `n`, we have recalled all ancestors of `p`
	// Hence, it follows that the head of the stack always satisfies the
	// Descendents-First-Relationship. As we search the trie in DFS manner, each
	// node of the trie is recalled (once). Therefore, the algorithm iterates all
	// nodes of the MTrie while guaranteeing Descendents-First-Relationship.

	// unprocessedRoot contains the trie's root before the first call of Next().
	// Thereafter, it is set to nil (which prevents repeated iteration through the trie).
	// This has the advantage, that we gracefully handle tries whose root node is nil.
	unprocessedRoot Node
	stack           []Node
}

// NewNodeIterator returns a new NodeIterator, which iterates through all nodes
// comprising the MTrie. The Iterator guarantees a DESCENDANTS-FIRST-RELATIONSHIP in
// the sequence of nodes it generates.
func NewNodeIterator(tr *Trie) *NodeIterator {
	// For a Trie with height H (measured by number of edges), the longest possible path contains H+1 vertices.
	i := &NodeIterator{
		stack:           make([]Node, 0, ledger.NodeMaxHeight+1),
		unprocessedRoot: tr.RootNode(),
	}

	return i
}

// Next moves the iterator to the next node in the DFS search.
func (i *NodeIterator) Next() bool {
	if i.unprocessedRoot != nil {
		// Initial call to Next() for a non-empty trie.
		i.dig(i.unprocessedRoot)
		i.unprocessedRoot = nil
		return true
	}

	// The current head of the stack, `n`, has been recalled.
	// Now, inspect `n`'s parent and dig into its right child, if it exists.
	n := i.pop()
	if len(i.stack) > 0 {
		// If there are more elements on the stack, the next element on the stack is `n`'s parent `p`.
		// Before `p` is recalled, we dig into the parent's right child, if we have not done so already.
		// As we descend into the left child with priority, the only case where we still dig into the
		// right child is if `n` is `p`'s left child.
		parent := i.peek()
		if parent.LeftChild() == n {
			i.dig(parent.RightChild())
		}
		return true
	}

	// Since `len(i.stack) == 0`, there are no more elements to recall.
	return false
}

// Value returns the value of the current node at the top of the iterator's stack.
func (i *NodeIterator) Value() Node {
	if len(i.stack) == 0 {
		return nil
	}
	return i.peek()
}

// pop pops the current node at the top of the iterator's stack.
func (i *NodeIterator) pop() Node {
	if len(i.stack) == 0 {
		return nil
	}
	headIdx := len(i.stack) - 1
	head := i.stack[headIdx]
	i.stack = i.stack[:headIdx]
	return head
}

// peek returns the node at the top of the stack.
func (i *NodeIterator) peek() Node {
	return i.stack[len(i.stack)-1]
}

// dig adds the children on the given node to the stack.
func (i *NodeIterator) dig(n Node) {
	if n == nil {
		return
	}

	// Go through each of the node's children, from left to right.
	for {
		i.stack = append(i.stack, n)
		if lChild := n.LeftChild(); lChild != nil {
			n = lChild
			continue
		}
		if rChild := n.RightChild(); rChild != nil {
			n = rChild
			continue
		}
		return
	}
}
