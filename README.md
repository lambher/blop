# Blop 

A little library that simplifies the way to load sounds and play them in any go applications.

```
go get -u github.com/lambher/blop
```

## Features

Blop is based on the [Beep package](https://github.com/faiface/beep) 
- **Load WAV, MP3 and FLAC.**
- **Play sound in easily way**
- **Very small codebase.** The core is just ~1K LOC.

## Usage

```go
package main

import "github.com/lambher/blop"

func main() {
    blop.LoadSound("blop", "sounds/blop.wav")

    // somewhere in your application
    blop.Play("blop")
}
```

## Examples

| [Example](https://github.com/lambher/blop/blob/master/example/main.go) 


## Dependencies

Blop uses [Beep package](https://github.com/faiface/beep) under the hood. Check its requirements to see what you need to install for building your application.

Running an already built application should work with no extra dependencies.

## Licence

[MIT](https://github.com/lambher/blop/blob/master/LICENSE)
