// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestMapPipe(t *testing.T) {
	in := make(chan interface{}, 5)
	out := make(chan interface{}, 5)
	count := 0
	NewPipe(in, out).Map(func(item interface{}) interface{} {
		count++
		return count
	})

	in <- 7
	in <- 4
	in <- 5
	for i := 1; i <= 3; i++ {
		if result := <-out; result.(int) != i {
			t.Fatal("mapping pipe received ", i, " items but output ", result.(int))
		}
	}

	close(in)
}
