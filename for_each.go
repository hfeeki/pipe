// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// A function which foreachs
type ForEachFunc func(item interface{})

// Execute a function for each item (without modifying the item)
func (p *Pipe) ForEach(fn ForEachFunc) *Pipe {
	p.addStage()
	go p.foreachHandler(fn, p.length-1)()

	return p
}

func (p *Pipe) foreachHandler(fn ForEachFunc, pos int) func() {
	return func() {
		for {
			item, ok := <-p.prevChan(pos)
			if !ok {
				break
			}

			fn(item)
			p.nextChan(pos) <- item
		}
		close(p.nextChan(pos))
	}
}
