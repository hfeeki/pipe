// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestInterposePipe(t *testing.T) {
	in := make(chan interface{}, 5)
	out := Interpose(in, false)

	in <- true
	in <- true
	in <- true

	expected := false
	for i := 0; i < 5; i++ {
		result := <-out
		expected = !expected
		if result != expected {
			t.Fatal("expected channel output to match", expected, "but got", result.(int))
		}
	}
}

func TestInterposeChainedConstructor(t *testing.T) {
	in := make(chan interface{}, 5)
	out := NewPipe(in).Interpose(false).Output

	in <- true
	in <- true
	in <- true

	expected := false
	for i := 0; i < 5; i++ {
		result := <-out
		expected = !expected
		if result != expected {
			t.Fatal("expected channel output to match", expected, "but got", result.(int))
		}
	}
}
