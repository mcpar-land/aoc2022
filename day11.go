package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Atoi(s string) int64 {
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return int64(n)
}

func Trim(prefix, s, suffix string) string {
	return strings.TrimSuffix(strings.TrimPrefix(s, prefix), suffix)
}

type Op string

const (
	Add Op = "+"
	Sub Op = "-"
	Mul Op = "*"
	Div Op = "/"
)

type Monkey struct {
	Number          int64
	Items           []int64
	Operation       Op
	OperationAmount int64
	OperationOnSelf bool
	TestDivisible   int64
	ThrowTrue       int64
	ThrowFalse      int64
	NInspected      int64
}

type Throw struct {
	To   int64
	Item int64
}

func ParseMonkey(input string) Monkey {
	lines := strings.Split(input, "\n")

	// parse number
	number := Atoi(Trim("Monkey ", lines[0], ":"))

	// parse starting items
	startingItems := []int64{}
	for _, n := range strings.Split(
		strings.TrimPrefix(lines[1], "  Starting items: "), ",",
	) {
		startingItems = append(startingItems, Atoi(n))
	}

	// parse operation
	_op := strings.Split(strings.TrimPrefix(lines[2], "  Operation: new = old "), " ")
	operation := Op(_op[0])
	operationAmount := int64(0)
	if _op[1] != "old" {
		operationAmount = Atoi(_op[1])
	}

	testDivisible := Atoi(strings.TrimPrefix(lines[3], "  Test: divisible by "))

	throwTrue := Atoi(strings.TrimPrefix(lines[4], "    If true: throw to monkey "))
	throwFalse := Atoi(strings.TrimPrefix(lines[5], "    If false: throw to monkey "))

	return Monkey{
		Number:          number,
		Items:           startingItems,
		Operation:       operation,
		OperationAmount: operationAmount,
		OperationOnSelf: _op[1] == "old",
		TestDivisible:   testDivisible,
		ThrowTrue:       throwTrue,
		ThrowFalse:      throwFalse,
		NInspected:      0,
	}
}

func (m Monkey) Inspect(item int64, divByThree bool) Throw {
	opAmt := m.OperationAmount
	if m.OperationOnSelf {
		opAmt = item
	}
	switch m.Operation {
	case Add:
		item += opAmt
		break
	case Mul:
		item *= opAmt
	default:
		panic("Unhandled operation: " + m.Operation)
	}

	if divByThree {
		item /= 3
	}

	throwTo := m.ThrowFalse
	if item%m.TestDivisible == 0 {
		throwTo = m.ThrowTrue
	}

	return Throw{
		Item: item,
		To:   throwTo,
	}
}

func (m *Monkey) Turn(divByThree bool) []Throw {
	throws := []Throw{}

	for _, item := range m.Items {
		throws = append([]Throw{m.Inspect(item, divByThree)}, throws...)
		m.NInspected += 1
	}

	m.Items = []int64{}

	return throws
}

func RunMonkeys(monkeys []Monkey, rounds int, divByThree bool) int64 {
	for i := 0; i < rounds; i++ {
		for m := range monkeys {
			for _, throw := range monkeys[m].Turn(divByThree) {
				monkeys[throw.To].Items = append(monkeys[throw.To].Items, throw.Item)
			}
		}
	}
	inspectedAmounts := []int64{}

	for _, monkey := range monkeys {
		fmt.Println("Monkey", monkey.Number, "inspected", monkey.NInspected)
		inspectedAmounts = append(inspectedAmounts, monkey.NInspected)
	}

	sort.Slice(inspectedAmounts, func(i, j int) bool {
		return inspectedAmounts[i] < inspectedAmounts[j]
	})
	inspectedAmounts = inspectedAmounts[len(inspectedAmounts)-2:]

	return inspectedAmounts[0] * inspectedAmounts[1]
}

func ParseMonkeys(input string) []Monkey {
	monkeys := []Monkey{}

	for _, monkeyInput := range strings.Split(input, "\n\n") {
		monkeys = append(monkeys, ParseMonkey(monkeyInput))
	}

	return monkeys
}

const ROUNDS int = 20

func Day11(input string) error {

	for _, monkey := range ParseMonkeys(input) {
		fmt.Printf("%#v\n", monkey)
	}

	fmt.Println("Answer 1:", RunMonkeys(ParseMonkeys(input), 20, true))

	fmt.Println("Answer 2:", RunMonkeys(ParseMonkeys(input), 20, false))
	fmt.Println("Answer 2:", RunMonkeys(ParseMonkeys(input), 1000, false))

	return nil
}
