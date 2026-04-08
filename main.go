package main

import (
	"fmt"
	process "text-editor/modifiers"
)

func main() {
	tokens := process.Split("hello, world am (up) here too! 1F (hex), I will be 10 (bin) years (cap) OLD (low)")
	tokens = process.BaseConv(tokens)
	tokens = process.AlphaConv(tokens)
	fmt.Println(tokens)
}
