package main

import (
	"fmt"
	//"bufio"
	//"io"
	"io/ioutil"
	"os"
)

// Check if an error occured
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// This is the interpreter loop which reads and parses each instruction
func parse(instructions []byte) {
	var memory [30000]byte // Brainfuck has 30000 bytes of memory
	var pc, p int

	for ; pc < len(instructions); pc++ {
		switch instructions[pc] {
		case '>':
			p++
		case '<':
			p--
		case '+':
			memory[p]++
		case '-':
			memory[p]--
		case '.':
			fmt.Printf("%c", memory[p])
		case ',':
			// fmt.Scanln(memory[p])
		case '[':
			if memory[p] == 0 {
				loop := 1
				for loop > 0 {
					pc++
					if (instructions[pc] == '[') {
						loop++;
					} else if (instructions[pc] == ']') {
						loop--;
					}
				}
			}
		case ']':
			if memory[p] != 0 {
				loop := 1
					for loop > 0 {
						pc--
						if (instructions[pc] == '[') {
							loop--;
						} else if (instructions[pc] == ']') {
							loop++;
						}
					}
				}
		}
	}
}

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Printf("Usage: %s filename", args[0])
		return
	}

	filename := args[1];
	file, err := ioutil.ReadFile(filename)
	check(err)

	//source := string(file)
	//fmt.Println(source)

	code := []byte(file)
	parse(code)
}
