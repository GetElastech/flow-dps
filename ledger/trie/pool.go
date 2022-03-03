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
	"sync"
)

type Pool struct {
	extensions *sync.Pool
	branches   *sync.Pool
	leaves     *sync.Pool
}

func NewPool(number int) *Pool {
	ePool := &sync.Pool{}
	ePool.New = func() interface{} {
		return &Extension{}
	}
	bPool := &sync.Pool{}
	bPool.New = func() interface{} {
		return &Branch{}
	}
	lPool := &sync.Pool{}
	lPool.New = func() interface{} {
		return &Leaf{}
	}

	// Pre allocate each node type.
	for i := 0; i < number; i++ {
		ePool.Put(ePool.New())
		bPool.Put(bPool.New())
		lPool.Put(lPool.New())
	}

	p := Pool{
		extensions: ePool,
		branches:   bPool,
		leaves:     lPool,
	}

	return &p
}
