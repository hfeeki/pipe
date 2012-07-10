// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// A function which takewhiles
type TakeWhileFunc func(item interface{}) bool

// Accept items from the input pipe until the given function returns false.
// After that, all input messages will be ignored and the output channel will
// be closed.
func (p *Pipe) TakeWhile(fn TakeWhile) *Pipe {
	p.addStage()
	go p.takewhileHandler(fn, p.length-1)()

	return p
}

func (p *Pipe) takewhileHandler(fn TakeWhileFunc, pos int) func() {
	return func() {
		for {
			item, ok := <-p.prevChan(pos)
			if !ok {
				break
			}

			// check if we should continue
			if !fn(item) {
				break
			}

			p.nextChan(pos) <- item
		}

		// hit the toggle, close the channel
		close(p.nextChan(pos))

		// drop any extra messages
		for {
			_, ok := <-p.prevChan(pos)
			if !ok {
				break
			}
		}
	}
}
