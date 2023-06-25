package main

import (
	"fmt"

	"github.com/konduruganesh/puppy"
)

func main() {
	fmt.Println("main thread is getting executed")
	str := puppy.Bark()
	fmt.Println("learning1+++++++++++++", str)

	puppy.BigBark()
}
