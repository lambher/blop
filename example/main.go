package main

import (
	"fmt"

	"github.com/lambher/blop"
)

func main() {
	blop.LoadSound("blop", "sounds/blop.wav")
	blop.LoadSound("rock", "sounds/rock.mp3")

	for {
		fmt.Println("press [ENTER] to play blop")
		fmt.Scanln()
		blop.Play("blop")
		fmt.Println("press [ENTER] to play rock")
		fmt.Scanln()
		blop.Play("rock")
	}
}
