// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestForEachPipe(t *testing.T) {
	in := make(chan interface{}, 5)
	out := make(chan interface{}, 5)
	count := 0
	NewPipe(in, out).ForEach(func(item interface{}) {
		count++
	})

	in <- 5
	in <- 6
	in <- 7

	// drain the pipe
	for i := 5; i <= 7; i++ {
		result := <-out
		if result.(int) != i {
			t.Fatal("counting ForEach pipe modified ", i, " into ", result.(int))
		}
	}

	if count != 3 {
		t.Fatal("counting ForEach pipe received 3 items but counted ", count)
	}

	close(in)
}
