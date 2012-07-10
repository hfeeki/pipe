// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// A function which mappers
type MapFunc func(item interface{}) interface{}

// Pass through the result of the map function for each item
func (p *Pipe) Map(fn MapFunc) *Pipe {
	p.addStage()
	go p.mapperHandler(fn, p.length-1)()

	return p
}
