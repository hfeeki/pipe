// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// Generate an sequence of numbers from start (inclusive) to end (exclusive)
// incrementing by step (default 1)
func Range(start, end int, step ...int) chan interface{} {
	out := make(chan interface{})

	step_size := 1
	if len(step) > 0 {
		step_size = step[0]
	}

	go func() {
		for i := start; i < end; i += step_size {
			out <- i
		}
		close(out)
	}()

	return out
}
