// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// A function which can be used to iteratively generate data
type IterateFunc func(previous interface{}) interface{}

// Generate an infinite sequence by repeatedly calling the given function with
// the previous value. The output will be x, f(x), f(f(x)), etc...
func Iterate(fn IterateFunc, x interface{}) chan interface{} {
	out := make(chan interface{})

	go func() {
		this := x
		for {
			out <- this
			this = fn(this)
		}
	}()

	return out
}
