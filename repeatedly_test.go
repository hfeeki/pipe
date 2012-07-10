// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestRepeatedly(t *testing.T) {
	i := 0
	channel := Repeatedly(func() interface{} {
		i++
		return i
	})

	for i := 0; i < 5; i++ {
		in := <-channel
		if in != i+1 {
			t.Fatal("Got unexpected result from Repeatedly function, got:", in, ", expected:", i+1)
		}
	}
}

func TestRepeatedlyBounded(t *testing.T) {
	i := 0
	channel := Repeatedly(func() interface{} {
		i++
		return i
	}, 3)

	for i := 0; i < 3; i++ {
		in := <-channel
		if in != i+1 {
			t.Fatal("Got unexpected result from Repeatedly function, got:", in, ", expected:", i+1)
		}
	}

	_, ok := <-channel
	if ok {
		t.Fatal("Expected the bounded repeating channel to close after 3, but it stayed open")
	}
}
