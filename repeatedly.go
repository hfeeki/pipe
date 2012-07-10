// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

// A function which can be used to iteratively generate data
type RepeatedlyFunc func() interface{}

// Generate an infinite sequence by repeatedly calling the given function. The output will be x, f(x), f(f(x)), etc...
func Repeatedly(fn RepeatedlyFunc, x ...int64) chan interface{} {
	out := make(chan interface{})

  if len(x) > 0 {
    go boundedRepeatedlyHandler(fn, x[0], out)
  } else {
    go unboundedRepeatedlyHandler(fn, out)
  }

	return out
}

func boundedRepeatedlyHandler(fn RepeatedlyFunc, bound int64, out chan interface{}) {
  for i := int64(0); i < bound; i++ {
    out <- fn()
  }
  close(out)
}

func unboundedRepeatedlyHandler(fn RepeatedlyFunc, out chan interface{}) {
  for {
    out <- fn()
  }
  close(out)
}
