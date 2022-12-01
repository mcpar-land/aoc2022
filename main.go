package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type DayFunc func(input string) error

var days = map[int]DayFunc{
	1: Day01,
}

func main() {
	dayFlag := flag.Int("d", 0, "Day to run")
	flag.Parse()

	day := *dayFlag

	if day == 0 {
		fmt.Println("Please supply a day.")
		os.Exit(1)
	}
	df, ok := days[day]
	if !ok {
		fmt.Printf("Day %d not found\n", day)
		os.Exit(1)
	}

	inputFilePath := fmt.Sprintf("./inputs/%02d.txt", day)
	fmt.Println("Reading", inputFilePath)
	buf, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	input := string(buf)

	err = df(input)
	if err != nil {
		fmt.Printf("Error running day %d\n", day)
		fmt.Println(err)
		os.Exit(1)
	}
}
