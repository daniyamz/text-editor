## Go-Reloaded ##

The objective of this project is making a simple text completion/editing/auto-correction tool. It's read the content of a file and make the necessary changes base on the project requirements.

## How to run the program ##
1. clone the repo by entering this `git clone https://github.com/daniyamz/text-editor` after you have go installed on your machine.
2. what a text in the sample.txt file example "hello (up) world!".
3. Open your terminal and run the program using this `go run .`
4. Open the result.txt to see the output.

## Files ##
1. **Main.go:** The file that run the whole program, it's where all the other functions are called.
2. **Sample.txt:** That's where the program reads from, that's where the user enter text.
3. **Result.txt:** This is the file the program write after processing the text.
4. **Modifiers:** This folder contains all the files of the program.

## Functions ##
1. **alpha.go:** This function check if the first letter of user input is a vowel and string before the string is letter "a", the letter "a" is convereted to a "an".
2. **punt.go:** This function handles punctuations in user input.  