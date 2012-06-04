// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// Implement this interface in your object to pass it to Pipe.Filter
type Filter interface {
	Filter(item interface{}) bool
}

// A function which filters
type FilterFunc func(item interface{}) bool

// Only pass through items when the filter returns true
func (p *Pipe) FilterFunc(fn FilterFunc) *Pipe {
	p.addStage()
	go p.filterHandler(fn, p.length-1)()

	return p
}

// Only pass through items when the filter returns true
func (p *Pipe) Filter(t Filter) *Pipe {
	p.FilterFunc(func(item interface{}) bool {
		return t.Filter(item)
	})

	return p
}

func (p *Pipe) filterHandler(fn FilterFunc, pos int) func() {
	return func() {
		for {
			item, ok := <-p.prevChan(pos)
			if !ok {
				break
			}

			if fn(item) {
				p.nextChan(pos) <- item
			}
		}
		close(p.nextChan(pos))
	}
}
