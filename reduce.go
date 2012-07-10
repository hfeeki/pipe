// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// A function which reduces
type ReduceFunc func(result, item interface{}) interface{}

// Accumulate the result of the reduce function being called on each item, then
// when the input channel is closed, pass the result to the output channel
func (p *Pipe) Reduce(initial interface{}, fn ReduceFunc) *Pipe {
	p.addStage()
	go p.reducerHandler(initial, fn, p.length-1)()

	return p
}

func (p *Pipe) reducerHandler(initial interface{}, fn ReduceFunc, pos int) func() {
	var result interface{} = initial
	return func() {
		for {
			item, ok := <-p.prevChan(pos)
			if !ok {
				break
			}

			result = fn(result, item)
		}
		// Input was closed, send the result
		p.nextChan(pos) <- result
		close(p.nextChan(pos))
	}
}
