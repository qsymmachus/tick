# tick ⏰

Tick is a simple timer utility for the command line. Think of it as an improved `sleep`.

## Installation

The only dependency is a [Go installation](https://golang.org/doc/install). To install tick, run this command:

```
go install github.com/qsymmachus/tick
```

## Usage

Running `tick -help` will print usage guidelines:

```
Usage of tick:
  -m int
    	Number of minutes the timer should run.
  -q	Quiet mode, suppresses countdown.
  -qq
    	Extra quiet mode, suppresses countdown and bell.
  -s int
    	Number of seconds the timer should run.
```

You can use tick as general purpose timer. By default, tick displays a countdown and rings your terminal bell when it's complete.

Run timer for 25 minutes and 10 seconds:

```
tick -m 25 -s 10
```

You can suppress the countdown output with the `-q` flag. It will still ring the terminal bell when it's done.

Let's send this tick to the background with `&` so we can run other commands while it runs:

```
tick -m 5 -q &
```

If you want to suppress the countdown _and_ the bell, use the `-qq` flag. In this mode, tick works the same way as `sleep`:

```
tick -s 3 -qq
```

An implementation of [sleep sort](https://kevlinhenney.medium.com/need-something-sorted-sleep-on-it-11fdf8453914) using tick:

```bash
#!/bin/bash

function sleeper() {
  tick -qq -s $1
  echo $1
}

while [ -n "$1" ]
do
  sleeper $1 &
  shift
done
wait
```
