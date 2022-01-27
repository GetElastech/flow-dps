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
	"github.com/onflow/flow-go/ledger/common/hash"
)

// Branch is a node is an intermediary node which has children.
// It does not need to contain a path, because its children are ordered
// based on their own path differences.
type Branch struct {
	hash  hash.Hash
	dirty bool
	left  Node
	right Node
}

// Hash returns the branch hash. If it is currently dirty, it is recomputed first.
func (b *Branch) Hash(height uint8) [32]byte {
	if b.dirty {
		b.computeHash(height)
	}
	return b.hash
}

// computeHash computes the branch hash by hashing its children.
func (b *Branch) computeHash(height uint8) {
	if b.left == nil && b.right == nil {
		panic("branch node should never have empty children")
	}

	var lHash, rHash hash.Hash
	if b.left != nil {
		lHash = b.left.Hash(height)
	} else {
		lHash = ledger.GetDefaultHashForHeight(int(height))
	}

	if b.right != nil {
		rHash = b.right.Hash(height)
	} else {
		rHash = ledger.GetDefaultHashForHeight(int(height))
	}

	b.hash = hash.HashInterNode(lHash, rHash)
	b.dirty = false
}