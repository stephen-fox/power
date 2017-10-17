# power

## What is it?
A Go library for managing power states.

## How do I use it?
If you would like the library:
```
go get github.com/stephen-fox/power
```

Or if you would like the command line application:
```
go get github.com/stephen-fox/power/command
```

### Library API

* `Restart` - Restarts the machine
* `Shutdown` - Shuts the machine down
* `Sleep` - Puts the machine to sleep

### Command line application usage
Example: `power -h`

```
Specify action with second argument: <application> [-s, -r, -o]
     -s     Sleep the machine
     -r     Restart the machine
     -o     Power the machine off
```