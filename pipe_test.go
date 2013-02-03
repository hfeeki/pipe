// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestNullPipe(t *testing.T) {
	in := make(chan interface{})
	out := make(chan interface{})
	NewPipe(in, out)

	in <- 5
	if result := <-out; result != 5 {
		t.Fatal("Null pipe received: 5 but output ", result)
	}

	close(in)
}

func TestMultiPipe(t *testing.T) {
	in := make(chan interface{})
	out := make(chan interface{})
	NewPipe(in, out).Filter(func(item interface{}) bool {
		return (item.(int) % 5) == 0
	}).Filter(func(item interface{}) bool {
		return (item.(int) % 2) == 0
	})

	in <- 2
	in <- 5
	in <- 10
	if result := <-out; result != 10 {
		t.Fatal("mod 2 and mod 5 pipe received 2, 5 and 10 but output ", result)
	}

	close(in)
}

func TestClosingPipe(t *testing.T) {
	in := make(chan interface{})
	out := make(chan interface{})
	NewPipe(in, out)

	close(in)
	if _, ok := <-out; ok {
		t.Fatal("closing the input pipe did not cascade to output")
	}
}

func TestModifyingRunningPipe(t *testing.T) {
	in := make(chan interface{})
	out := make(chan interface{})
	pipe := NewPipe(in, out)

	in <- 5
	if result := <-out; result != 5 {
		t.Fatal("Unmodified pipe received: 5 but output ", result)
	}

	pipe.Filter(func(item interface{}) bool {
		return item.(int) < 5
	})

	in <- 6
	in <- 3
	if result := <-out; result != 3 {
		t.Fatal("Modified pipe received: 6 and 3 but output ", result)
	}

	close(in)
}
