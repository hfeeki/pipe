// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestIterate(t *testing.T) {
	channel := Iterate(func(i interface{}) interface{} {
		return i.(int) + 1
	}, 1)

	for i := 0; i < 5; i++ {
		in := <-channel
		if in != i+1 {
			t.Fatal("Got unexpected result from Iterate function, got:", in, ", expected:", i+1)
		}
	}
}
