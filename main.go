package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	process "text-editor/modifiers"
)

func main() {
	// Validating the number of arguments the user should pass
	if len(os.Args) < 3 {
		fmt.Println("Try this: go run . sample.txt result.txt")
		return
	}

	inputfile := os.Args[1]
	outputfile := os.Args[2]

	if inputfile != "sample.txt" {
		fmt.Println("inputfile should be sample.txt")
		return
	}
	//openig the file to read from
	cont, err := os.Open(inputfile)
	if err != nil {
		log.Fatal()
	}
	var result strings.Builder
	// reading the file content line by line.
	scanner := bufio.NewScanner(cont)
	err = scanner.Err()
	if err != nil {
		fmt.Println("Error occurd", err)
		return
	}

	for scanner.Scan() {
		data := scanner.Text()
		tokens, err := process.Split(data)
		if err != nil {
			fmt.Println("Error Occured: ", err)
			return
		}
		tokens = process.BaseConv(tokens)
		tokens, err = process.AlphaConv(tokens)
		tokens, err = process.Alpha(tokens)
		if err != nil {
			fmt.Println("occured", err)
			return
		}
		tokens1 := process.PunctControl(tokens)
		tk, err := process.QuotControl(tokens1)
		if err != nil {
			fmt.Println("Error occured: ", err)
			return
		}
		result.WriteString(tk)
		result.WriteString("\n")

	}
	cont.Close()

	//tk := strings.Join(tokens, " ")
	err = os.WriteFile(outputfile, []byte(result.String()), 0644)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("Process.....")
}
