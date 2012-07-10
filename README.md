# pipe

Concurrent, sequential, transformations along Golang channels.

## Usage

```
import "github.com/paulbellamy/pipe"
```

## Interface

Pipes are created with the ```NewPipe(input, output chan interface{}) *Pipe``` method.

After that there are several chaining methods to build up the processing. Once the pipe is prepared, simply pipe items into the input channel and retrieve the results from the output channel.

Be careful, because some of the transformations (e.g. Reduce, Skip) result in channels which are 'leaky'. Meaning that one item in may not equal one item out.

For example, to count the number of items passing through a channel:

```Go
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
```

You can, of course, modify the items flowing through the pipe:

```Go
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
```

## Available Transformations

* Filter(func(item interface{}) bool)
* ForEach(func(item interface{}))
* Map(func(item interface{}) interface{})
* Reduce(initial interface{}, func(accumulator interface{}, item interface{}) interface{})
* Skip(n int64)
* SkipWhile(func(item interface{}) bool)
* Take(n int64)
* TakeWhile(func(item interface{}) bool)
* Zip(other chan interface{})

## Godoc

```
type FilterFunc func(item interface{}) bool
    A function which filters

type ForEachFunc func(item interface{})
    A function which foreachs

type MapFunc func(item interface{}) interface{}
    A function which mappers

type Pipe struct {
    // contains filtered or unexported fields
}
    A Pipe is a set of transforms being applied along the channel

func NewPipe(in, out chan interface{}) *Pipe
    Return a new Pipe object which echoes input to output

func (p *Pipe) Filter(fn FilterFunc) *Pipe
    Only pass through items when the filter returns true

func (p *Pipe) ForEach(fn ForEachFunc) *Pipe
    Execute a function for each item (without modifying the item)

func (p *Pipe) Map(fn MapFunc) *Pipe
    Pass through the result of the map function for each item

func (p *Pipe) Reduce(initial interface{}, fn ReduceFunc) *Pipe
    Accumulate the result of the reduce function being called on each item,
    then when the input channel is closed, pass the result to the output
    channel

func (p *Pipe) Skip(num int64) *Pipe
    Skip a given number of items from the input pipe. After that number has
    been dropped, the rest are passed straight through.

func (p *Pipe) SkipWhile(fn SkipWhileFunc) *Pipe
    Skip the items from the input pipe until the given function returns
    true. After that , the rest are passed straight through.

func (p *Pipe) Take(num int64) *Pipe
    Accept only the given number of items from the input pipe. After that
    number has been received, all input messages will be ignored and the
    output channel will be closed.

func (p *Pipe) TakeWhile(fn TakeWhile) *Pipe
    Accept items from the input pipe until the given function returns false.
    After that, all input messages will be ignored and the output channel
    will be closed.

func (p *Pipe) Zip(other chan interface{}) *Pipe
    Group each message from the input channel with it's corresponding
    message from the other channel. This will block on the first channel
    until it receives a message, then block on the second until it gets one
    from there. At that point an array containing both will be sent to the
    output channel.

    For example, if channel a is being zipped with channel b, and output on
    channel c:

  a <- 1
  b <- 2
  result := <-c

    Here, result will equal []interface{}{1, 2}

type ReduceFunc func(result, item interface{}) interface{}
    A function which reduces

type SkipWhileFunc func(item interface{}) bool
    A function which skipwhiles

type TakeWhileFunc func(item interface{}) bool
    A function which takewhiles
```
