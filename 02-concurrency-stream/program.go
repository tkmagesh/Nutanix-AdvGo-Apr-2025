package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
Example to demonstrate reading & write file and convert string to int
*/
func main() {
	file, err := os.Open("data1.dat")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		if val, err := strconv.Atoi(line); err == nil {
			total += val
		}
	}
	resultFile, err := os.Create("total.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer resultFile.Close()
	fmt.Fprintf(resultFile, "Total : %d\n", total)
	fmt.Println("Done")
}
