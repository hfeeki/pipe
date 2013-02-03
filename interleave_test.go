// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestInterleavePipe(t *testing.T) {
	other := make(chan interface{}, 5)
	in := make(chan interface{}, 5)
	out := Interleave(in, other)

	in <- 5
	in <- 7
	in <- 9
	other <- 6
	other <- 8

	for i := 5; i <= 9; i++ {
		result := <-out
		expected := i
		if result != expected {
			t.Fatal("expected channel output to match", expected, "but got", result)
		}
	}
}

func TestMultiInterleavePipe(t *testing.T) {
	other := make(chan interface{}, 5)
	other2 := make(chan interface{}, 5)
	in := make(chan interface{}, 5)
	out := Interleave(in, other, other2)

	in <- 1
	in <- 4
	in <- 7
	other <- 2
	other <- 5
	other2 <- 3
	other2 <- 6

	for i := 1; i <= 6; i++ {
		result := <-out
		expected := i
		if result != expected {
			t.Fatal("expected channel output to match", expected, "but got", result)
		}
	}
}

func TestInterleaveChainedConstructor(t *testing.T) {
	other := make(chan interface{}, 5)
	other2 := make(chan interface{}, 5)
	in := make(chan interface{}, 5)
	out := NewPipe(in).Interleave(other, other2).Output

	in <- 1
	in <- 4
	in <- 7
	other <- 2
	other <- 5
	other2 <- 3
	other2 <- 6

	for i := 1; i <= 6; i++ {
		result := <-out
		expected := i
		if result != expected {
			t.Fatal("expected channel output to match", expected, "but got", result)
		}
	}
}
