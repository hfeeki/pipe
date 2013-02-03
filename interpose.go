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
func Interpose(input chan interface{}, item interface{}) chan interface{} {
	output := make(chan interface{})
	first := true
	go func() {
		// only send num items
		for {
			a, ok := <-input
			if !ok {
				break
			}

			if first {
				first = false
			} else {
				output <- item
			}

			output <- a
		}

		close(output)
	}()
	return output
}

// Helper for the chained constructor
func (p *Pipe) Interpose(item interface{}) *Pipe {
	p.Output = Interpose(p.Output, item)
	return p
}
