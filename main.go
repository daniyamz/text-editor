package main

import (
	"fmt"
	"os"
	"strings"
	process "text-editor/modifiers"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Try this: go run . sample.txt result.txt")
	}
	inputfile := os.Args[1]
	outputfile := os.Args[2]
	cont, err := os.ReadFile(inputfile)
	if err != nil {
		fmt.Println("error occured", err)
	}
	data := string(cont)
	tokens := process.Split(data)
	tokens = process.BaseConv(tokens)
	tokens = process.AlphaConv(tokens)
	tk := strings.Join(tokens, " ")
	err = os.WriteFile(outputfile, []byte(tk), 0644)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("Process.....")
}
