// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// Generate an infinite sequence by repeating the value. Can be bounded by
// passing x.
func Repeat(item interface{}, x ...int) chan interface{} {
	out := make(chan interface{})

	if len(x) > 0 {
		go boundedRepeatHandler(item, x[0], out)
	} else {
		go unboundedRepeatHandler(item, out)
	}

	return out
}

func boundedRepeatHandler(item interface{}, bound int, out chan interface{}) {
	for i := int(0); i < bound; i++ {
		out <- item
	}
	close(out)
}

func unboundedRepeatHandler(item interface{}, out chan interface{}) {
	for {
		out <- item
	}
	close(out)
}
