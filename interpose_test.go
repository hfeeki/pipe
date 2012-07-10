// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestInterposePipe(t *testing.T) {
	in := make(chan interface{}, 5)
	out := make(chan interface{}, 5)
	other := make(chan interface{}, 5)
	NewPipe(in, out).Interpose(false)

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
