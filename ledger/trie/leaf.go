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

// Leaf nodes are found at the end of each path of the trie. They do not contain
// a part of the path, so they could, in theory, be shuffled around easily. This
// is made more difficult by the Flow implementation of the trie, which hashes
// the height of a leaf as part of the node hash, and uses the a height based on
// the sparse trie, which changes as the trie fills up, instead of the height in
// terms of path traversed, which would always be the same.
type Leaf struct {

	// The hash of the leaf node is computed whenever it changes as part of
	// insertions, so that it is never dirty.
	hash [32]byte

	// The path is kept as a byte slice, which allows us to share the path
	// between all of the nodes on that path when inserting, reducing memory use
	// significantly.
	path [32]byte

	// We insert the payload into the KV store and keep its hash here. Using the
	// payload hash rather than the node hash improves insertion performance by
	// avoiding storing a payload again when the leaf hash changes.
	key [32]byte
}

// Hash returns the leaf hash.
func (l *Leaf) Hash(height int) [32]byte {
	return l.hash
}
