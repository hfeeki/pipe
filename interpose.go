// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// Alternate messages from the channel with repeating the item. Sending a
// message from the channel, then the item, etc... If the input channel closes,
// the output will be closed. The final thing through the input channel will be the final item from the output channel. e.g.
//
//   out := make(chan interface{})
//   NewPipe(Range(0,3), out).Interpose('a')
//
//   <-out // 0
//   <-out // 'a'
//   <-out // 1
//   <-out // 'a'
//   <-out // 2
//   // output is now closed
//
func (p *Pipe) Interpose(item interface{}) *Pipe {
	p.addStage()
	go p.interposeHandler(other, p.length-1)()

	return p
}

func (p *Pipe) interposeHandler(item interface{}, pos int) func() {
	first := true
	return func() {
		// only send num items
		for {
			a, ok := <-p.prevChan(pos)
			if !ok {
				break
			}

			if first {
				first = false
			} else {
				p.nextChan(pos) <- item
			}

			p.nextChan(pos) <- a
		}

		close(p.nextChan(pos))
	}
}
