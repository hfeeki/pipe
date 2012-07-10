// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	channel := Repeat(5)

	for i := 0; i < 5; i++ {
		in := <-channel
		if in != 5 {
			t.Fatal("Got unexpected result from Repeat function, got:", in, ", expected:", 5)
		}
	}
}

func TestRepeatBounded(t *testing.T) {
	channel := Repeat(5, 3)

	for i := 0; i < 3; i++ {
		in := <-channel
		if in != 5 {
			t.Fatal("Got unexpected result from Repeatedly function, got:", in, ", expected:", 5)
		}
	}

	_, ok := <-channel
	if ok {
		t.Fatal("Expected the bounded repeating channel to close after 3, but it stayed open")
	}
}
