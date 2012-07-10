// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// Alternate messages from each channel. Sending a message from the first
// channel, then one from the second, etc... If either input channel closes,
// the output will be closed.
func (p *Pipe) Interleave(other chan interface{}) *Pipe {
	p.addStage()
	go p.interleaveHandler(other, p.length-1)()

	return p
}

func (p *Pipe) interleaveHandler(other chan interface{}, pos int) func() {
	return func() {
		// only send num items
		for {
			a, ok := <-p.prevChan(pos)
			if !ok {
				break
			}

			p.nextChan(pos) <- a

			b, ok := <-other
			if !ok {
				break
			}

			p.nextChan(pos) <- b
		}

		close(p.nextChan(pos))
	}
}
