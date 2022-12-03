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
	2: Day02,
	3: Day03,
}

func main() {
	dayFlag := flag.Int("d", 0, "Day to run")
	flag.Parse()

	day := *dayFlag

	err := executeDay(day)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func executeDay(day int) error {
	if day == 0 {
		return fmt.Errorf("Please supply a day")
	}
	df, ok := days[day]
	if !ok {
		return fmt.Errorf("Day %d not found", day)
	}
	inputFilePath := fmt.Sprintf("./inputs/%02d.txt", day)
	buf, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		return err
	}
	err = df(string(buf))
	if err != nil {
		return fmt.Errorf("Error running day %d: %v", day, err)
	}
	return nil
}
