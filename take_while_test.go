// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestTakeWhilePipe(t *testing.T) {
	in := make(chan interface{}, 5)
	out := make(chan interface{}, 5)
	take := true
	NewPipe(in, out).TakeWhile(func(item interface{}) bool {
		return take
	})

	in <- 7
	in <- 4
	take = false
	in <- 5

	<-out
	<-out
	if _, ok := <-out; ok {
		t.Fatal("takewhile pipe should have closed the channel after turning it off")
	}

	close(in)
}
