# pipe

Concurrent, sequential, transformations along Golang channels.

## Usage

```
import "github.com/paulbellamy/pipe"
```

## Interface

Pipes are created with the ```NewPipe(input, output chan interface{}) *Pipe``` method.

After that there are several chaining methods to build up the processing. Once the pipe is prepared, simply pipe items into the input channel and retrieve the results from the output channel. Pipes can be modified at any time, even after messages have been sent.

Be careful, because some of the transformations (e.g. Reduce, Skip) result in channels which are 'leaky'. Meaning that one item in may not equal one item out.

For example, to count the number of items passing through a channel:

```Go
// Define our counter
count := 0
counter_func := func(item interface{}) {
  count++
}

// Set up our pipe
input := make(chan interface{}, 5)

// Add our counter into the pipe
output := ForEach(input, counter_func)

// Now we send some items
input <- true
input <- true
input <- true

// Check how many have gone through
fmt.Println(count) // prints "3"
```

You can, of course, modify the items flowing through the pipe:

```Go
// Set up our pipe
input := make(chan interface{}, 5)
output := NewPipe(input, output).
  Filter(func(item interface{}) bool {
    // Only allow ints divisible by 5
    return (item.(int) % 5) == 0
  }).
  Map(func(item interface{}) interface{} {
    // Add 2 to each
    return item.(int) + 2
  }).
  Output

// Now we send some items
input <- 1 // will be dropped
input <- 5 // will come through as 7
```

## Creating Pipes

### NewPipe(out chan interface{})

Return a new Pipe object which echoes input to output. Additional
transformations can then be 'chained' onto the pipe, to modify the output.

## Generators

There are several generator functions included which can be used as a pipe's
input channel.

### Iterate(func(item interface{}) interface{}, x interface{}) chan interface{}

Generate an infinite sequence by repeatedly calling the given function
with the previous value. The output will be x, f(x), f(f(x)), etc...

### Range(start, end int, step ...int) chan interface{}

Generate an sequence of numbers from start (inclusive) to end
(exclusive) incrementing by step (default 1)

### Repeat(item interface{}, x ...int) chan interface{}

Generate an infinite sequence by repeating the value. Can be bounded by
passing x.

### Repeatedly(func() interface{}, x ...int) chan interface{}

Generate an infinite sequence by repeatedly calling the given function.
The function should take no arguments, and ideally be side-effect free.
The output will be x, f(x), f(f(x)), etc...

## Available Transformations

### Filter(func(item interface{}) bool)

Only pass through items when the filter returns true

### ForEach(func(item interface{}))

Execute a function for each item (without modifying the item)

### Interleave(other ...chan interface{})

Alternate messages from each channel. Sending a message from the first
channel, then one from the second, etc... If any input channel
closes, the output will be closed.

### Interpose(item interface{})

Alternate messages from the channel with repeating the item. Sending a
message from the channel, then the item, etc... If the input channel
closes, the output will be closed. The final thing through the input
channel will be the final item from the output channel. e.g.

```Go
out := NewPipe(Range(0,3)).Interpose('a').Output
<-out // 0
<-out // 'a'
<-out // 1
<-out // 'a'
<-out // 2
// output is now closed
```

### Map(func(item interface{}) interface{})

Pass through the result of the map function for each item

### Reduce(initial interface{}, func(accumulator interface{}, item interface{}) interface{})

Accumulate the result of the reduce function being called on each item,
then when the input channel is closed, pass the result to the output
channel

### Skip(n int)

Skip a given number of items from the input pipe. After that number has
been dropped, the rest are passed straight through.

### SkipWhile(func(item interface{}) bool)

Skip the items from the input pipe until the given function returns
true. After that , the rest are passed straight through.

### Take(n int)

Accept only the given number of items from the input pipe. After that
number has been received, all input messages will be ignored and the
output channel will be closed.

### TakeWhile(func(item interface{}) bool)

Accept items from the input pipe until the given function returns false.
After that, all input messages will be ignored and the output channel
will be closed.

### Zip(other chan interface{})

Group each message from the input channel with it's corresponding
message from the other channel. This will block on the first channel
until it receives a message, then block on the second until it gets one
from there. At that point an array containing both will be sent to the
output channel.

For example, if channel a is being zipped with channel b, and output on
channel c:

```Go
  a <- 1
  b <- 2
  result := <-c
```

Here, result will equal []interface{}{1, 2}

## More Info

See the ```godoc``` command for more information.
