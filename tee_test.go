// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestTeePipe(t *testing.T) {
	in := make(chan interface{}, 10)
	out2 := make(chan interface{}, 10)
	out3 := make(chan interface{}, 10)
	out1 := Tee(in, out2, out3)

	for i := 0; i < 5; i++ {
		in <- i
	}

	for _, c := range []chan interface{}{out1, out2, out3} {
		results := []int{}
		for i := 0; i < 5; i++ {
			value, ok := <-c
			if !ok {
				break
			}

			if value != i {
				t.Fatal("Tee pipe output unexpected value. Expected: ", i, " Got: ", value)
			}

			results = append(results, value.(int))
		}

		if len(results) != 5 {
			t.Fatal("Tee pipe received 5 items and should have output 5, but output ", len(results))
		}
	}
}

func TestTeeChainedConstructor(t *testing.T) {
	in := make(chan interface{}, 10)
	out2 := make(chan interface{}, 10)
	out3 := make(chan interface{}, 10)
	out1 := NewPipe(in).Tee(out2, out3).Output

	for i := 0; i < 5; i++ {
		in <- i
	}

	for _, c := range []chan interface{}{out1, out2, out3} {
		results := []int{}
		for i := 0; i < 5; i++ {
			value, ok := <-c
			if !ok {
				break
			}

			if value != i {
				t.Fatal("Tee pipe output unexpected value. Expected: ", i, " Got: ", value)
			}

			results = append(results, value.(int))
		}

		if len(results) != 5 {
			t.Fatal("Tee pipe received 5 items and should have output 5, but output ", len(results))
		}
	}
}
