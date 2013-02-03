// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// A Pipe is a set of transforms being applied along the channel
type Pipe struct {
	Output chan interface{}
}

// Return a new Pipe object which echoes input to output
func NewPipe(input chan interface{}) *Pipe {
	// echo input to output
	return &Pipe{Output: input}
}
