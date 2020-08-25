package main

import (
	"fmt"

	"github.com/lambher/blop"
)

func main() {
	err := blop.LoadSound("blop", "sounds/blop.wav")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = blop.LoadSound("rock", "sounds/rock.mp3")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		fmt.Println("press [ENTER] to play blop")
		fmt.Scanln()
		blop.Play("blop")
		fmt.Println("press [ENTER] to play rock")
		fmt.Scanln()
		blop.Play("rock")
	}
}
