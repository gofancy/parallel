/*
Package parallel implements an abstraction of the Go blog,
Go Concurrency Patterns: Pipelines and cancellation,
which is a great article with great examples. After writing several scripts based on
bounded.go (https://blog.golang.org/pipelines/bounded.go), I realized that it
would be a good idea to build a tiny library from it, so that I do not need to
copy & paste the code every time. Therefore, this library was built.

What does it do?

* It creates channels to set up fan-in, fan-out mechanisms.

* It starts a number of goroutines, which call user-defined functions to do some
actual work.

* It creates a "done" channel for cancellation. All goroutines will exit when the
"done" channel is closed.

Usage

The file, examples/md5all.go, realized same functionality as the bounded.go.
You can see how the library is used from it.

1. create a runner:

	r := parallel.New()

2. set a ProduceFunc and a ConsumeFunc

	r.Produce = func(products chan<- interface{}, done <-chan struct{}) {
        ...
	}

	r.Consume = func(product interface{}, done <-chan struct{}) {
        ...
	}

3. optionally set the number of producers and the number of consumers

	r.NumProducers = 2
	r.NumConsumers = 4

4. optionally enable cancellation by CTRL-C

	parallel.TrapInterrupt(r)

5. optionally you can also configure it if it should process all data in channels
on cancellation

	r.ConsumeAll = false // the goroutines will exit immidiately on cancellation

6. start to run

	r.Run()

## Get it

    go get github.com/gofancy/parallel
*/
package parallel
