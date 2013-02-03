// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// Generates a pipe of the elements.
func Sequence(items ...interface{}) chan interface{} {
	out := make(chan interface{})

	go func() {
		for _, item := range items {
			out <- item
		}
		close(out)
	}()

	return out
}
