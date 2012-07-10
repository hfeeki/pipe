// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestInterleavePipe(t *testing.T) {
	in := make(chan interface{}, 5)
	out := make(chan interface{}, 5)
	other := make(chan interface{}, 5)
	NewPipe(in, out).Interleave(other)

	in <- 5
	in <- 7
	in <- 9
	other <- 6
	other <- 8

	for i := 5; i <= 9; i++ {
		result := <-out
		expected := i
		if result != expected {
			t.Fatal("expected channel output to match", expected, "but got", result.(int))
		}
	}
}
