package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFlag := flag.String("csv", "problems.csv", "IMP: csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFlag)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(lines)
}
