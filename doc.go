// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/*
Package pipe provides concurrent and (relatively) transparent transformations
along Golang channels.

For example, to count the number of items passing through a channel:

  // Define our counter
  count := 0

  // Set up our pipe
  input := make(chan interface{}, 5)
  output := make(chan interface{}, 5)
  pipe := NewPipe(input, output)

  // Add our counter into the pipe
  pipe.ForEach(func(item interface{}) {
    count++
  })

  // Now we send some items
  input <- true
  input <- true
  input <- true

  // Check how many have gone through
  fmt.Println(count) // prints "3"

You can, of course, modify the items flowing through the pipe:

  // Set up our pipe
  input := make(chan interface{}, 5)
  output := make(chan interface{}, 5)

  NewPipe(input, output).Filter(func(item interface{}) bool {
    // Only allow ints divisible by 5
    return (item.(int) % 5) == 0
  }).Map(func(item interface{}) interface{} {
    // Add 2 to each
    return item.(int) + 2
  })

  // Now we send some items
  input <- 1 // will be dropped
  input <- 5 // will come through as 7
*/
package pipe
