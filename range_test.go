// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestRange(t *testing.T) {
	channel := Range(1, 10, 2)

	for i := 1; i < 10; i += 2 {
		in := <-channel
		if in != i {
			t.Fatal("Got unexpected result from Range function, got:", in, ", expected:", i)
		}
	}

	_, ok := <-channel
	if ok {
		t.Fatal("expected range channel to close after reaching limit, but it didn't")
	}
}

func TestRangeNoStep(t *testing.T) {
	channel := Range(1, 10)

	for i := 1; i < 10; i++ {
		in := <-channel
		if in != i {
			t.Fatal("Got unexpected result from Range function, got:", in, ", expected:", i)
		}
	}

	_, ok := <-channel
	if ok {
		t.Fatal("expected range channel to close after reaching limit, but it didn't")
	}
}
