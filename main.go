package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := strings.Split(os.Args[1], "\n")

	fonts := []string{"standard.txt", "shadow.txt", "thinkertoy.txt"}
	fmt.Println("\nChoose font, type:\n1 for standard\n2 for shadow\n3 for thinkertoy\n-------------------")
	var userinput string
	fmt.Scanln(&userinput)
	number, _ := strconv.Atoi(userinput)
	font := fonts[number-1]

	for _, word := range input {
		Lines(word, font)
	}
}

func Lines(word string, filename string) {
	for linecount := 0; linecount < 9; linecount++ { // prints all characters line by line
		for _, char := range word {
			line := ReadLine(filename, int(char-32)*9+1+linecount)
			fmt.Print(line)
		}
		fmt.Println()
	}
}

func ReadLine(filename string, lineNum int) (line string) { // input line nr and output that given line text as a string
	f, _ := os.Open(filename) // opens file
	defer f.Close()           // close file when the rest of the func is done
	scanner := bufio.NewScanner(f)
	lastline := 0

	for scanner.Scan() { //scan file until given line
		lastline++
		if lastline == lineNum {
			return scanner.Text() //return given line
		}
	}
	return line
}
