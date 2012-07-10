// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestReducePipe(t *testing.T) {
	in := make(chan interface{}, 5)
	out := make(chan interface{}, 5)
	NewPipe(in, out).Reduce(0, func(sum, item interface{}) interface{} {
		return sum.(int) + item.(int)
	})

	in <- 5
	in <- 10
	in <- 20
	close(in)

	result, ok := <-out
	if !ok {
		t.Fatal("output channel was closed before we retrieved the result")
	}

	if result.(int) != 35 {
		t.Fatal("reducing (sum) pipe received 5, 10, and 20 items but output ", result.(int))
	}
}
