// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// Alternate messages from each channel. Sending a message from the first
// channel, then one from the second, etc... If any input channel closes,
// the output will be closed.
func Interleave(chans ...chan interface{}) chan interface{} {
	output := make(chan interface{})
	go func() {
		for {
			for _, c := range chans {
				value, ok := <-c
				if !ok {
					close(output)
					return
				}

				output <- value
			}
		}
	}()
	return output
}

// Helper for the chained constructor
func (p *Pipe) Interleave(other ...chan interface{}) *Pipe {
	chans := []chan interface{}{p.Output}
	chans = append(chans, other...)

	p.Output = Interleave(chans...)
	return p
}
