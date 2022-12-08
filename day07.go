package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type FileSysItem struct {
	Name     string
	IsDir    bool
	Size     int
	Children []FileSysItem
}

func (f *FileSysItem) CalcSize() {
	if f.IsDir {
		size := 0
		for i := range f.Children {
			f.Children[i].CalcSize()
			size += f.Children[i].Size
		}
		f.Size = size
	}
}

func ParseCmds(
	stack []*FileSysItem,
	lines []string,
) []string {
	if len(lines) == 0 {
		return []string{}
	}
	current := stack[len(stack)-1]
	fmt.Println("Parsing line:", lines[0])
	line := strings.Split(lines[0], " ")

	if line[0] == "$" { // if it's a command
		if line[1] == "cd" {
			if line[2] == ".." {
				return ParseCmds(stack[:len(stack)-1], lines[1:])
			}
			cd := line[2]
			for i := range current.Children {
				if current.Children[i].Name == cd {
					return ParseCmds(append(stack, &current.Children[i]), lines[1:])
				}
			}
			panic(fmt.Sprintf("Child dir %s not found", cd))
		} else if line[1] == "ls" {
			lsLines := []FileSysItem{}
			_i := 0
			for i := 1; i < len(lines) && lines[i][0] != '$'; i++ {
				fl := strings.Split(lines[i], " ")
				size := -1
				if fl[0] != "dir" {
					size, _ = strconv.Atoi(fl[0])
				}
				lsLines = append(lsLines, FileSysItem{
					Name:     fl[1],
					Size:     size,
					IsDir:    fl[0] == "dir",
					Children: []FileSysItem{},
				})
				_i = i
			}
			current.Children = lsLines
			fmt.Println("Skipping", _i, "lines")
			return ParseCmds(stack, lines[_i+1:])
		}
		panic("Unrecognized command: " + lines[0])
	} else { // if it's not a command
		panic("Got a non command: " + lines[0])
	}
}

func SumOfDirsOverThreshold(threshold int, dir FileSysItem, sum int) int {
	if dir.Name != "/" && dir.IsDir && dir.Size <= threshold {
		sum += dir.Size
	}
	for _, child := range dir.Children {
		sum = SumOfDirsOverThreshold(threshold, child, sum)
	}
	return sum
}

func GetDirs(root FileSysItem) []FileSysItem {
	return getDirs(root, []FileSysItem{})
}

func getDirs(current FileSysItem, dirs []FileSysItem) []FileSysItem {
	if current.IsDir {
		dirs = append(dirs, current)
		for _, child := range current.Children {
			dirs = getDirs(child, dirs)
		}
	}
	return dirs
}

func Day07(input string) error {

	root := FileSysItem{
		Name:     "/",
		IsDir:    true,
		Size:     -1,
		Children: []FileSysItem{},
	}

	ParseCmds([]*FileSysItem{&root}, strings.Split(input, "\n")[1:])

	root.CalcSize()

	preview, err := json.MarshalIndent(root, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(preview))

	dirs := GetDirs(root)
	fmt.Println("Found", len(dirs), "dirs")

	sum := 0
	for _, dir := range dirs {
		if dir.Size <= 100000 {
			sum += dir.Size
		}
	}

	fmt.Println("Answer 1:", sum)

	target := 30000000

	current := root.Size

	smallest := current

	for _, dir := range dirs {
		fmt.Println(dir.Name, dir.Size)
		if dir.Size < smallest {
			fmt.Printf("%s: %v - %v = %v <= %v\n", dir.Name, current, dir.Size, current-dir.Size, target)
			if current-dir.Size <= target {
				fmt.Println("New smallest!")
				smallest = dir.Size
			}
		}
	}

	fmt.Println("Answer 2 incomplete:", smallest)

	return nil
}
