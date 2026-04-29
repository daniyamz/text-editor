package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	process "text-editor/modifiers"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Try this: go run . sample.txt result.txt")
	}
	inputfile := os.Args[1]
	outputfile := os.Args[2]
	// cont, err := os.ReadFile(inputfile)
	// if err != nil {
	// 	fmt.Println("error occured", err)
	// }
	cont, err := os.Open(inputfile)
	if err != nil {
		log.Fatal()
	}
	result := ""
	scanner := bufio.NewScanner(cont)
	for scanner.Scan() {
		data := scanner.Text()
		tokens := process.Split(data)
		tokens = process.BaseConv(tokens)
		tokens = process.AlphaConv(tokens)
		tokens = process.Alpha(tokens)
		tokens1 := process.PunctControl(tokens)
		tk := process.QuotControl(tokens1)
		result += tk + "\n"

	}
	cont.Close()

	//tk := strings.Join(tokens, " ")
	err = os.WriteFile(outputfile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("Process.....")
}
