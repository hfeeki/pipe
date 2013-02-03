// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestSequence(t *testing.T) {
	channel := Sequence(0, 1, 2, 3, 4)

	for i := 0; i < 5; i++ {
		in := <-channel
		if in != i {
			t.Fatal("Got unexpected result from Sequence function, got:", in, ", expected:", i)
		}
	}

	if _, ok := <-channel; ok {
		t.Fatal("Sequence should close the channel.")
	}
}
