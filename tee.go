// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// Additionally send each incoming value to each other channel specified.
func Tee(input chan interface{}, others ...chan interface{}) chan interface{} {
	output := make(chan interface{})
	go func() {
		// only send num items
		for item := range input {
			// Send it to the original channel
			output <- item

			// Send the item to all the other chanels
			for _, c := range others {
				c <- item
			}
		}

		close(output)

		// Close all the other channels as well
		for _, c := range others {
			close(c)
		}
	}()
	return output
}

// Helper for the chained constructor
func (p *Pipe) Tee(others ...chan interface{}) *Pipe {
	p.Output = Tee(p.Output, others...)
	return p
}
