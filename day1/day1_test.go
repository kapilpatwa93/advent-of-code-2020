package day1

import (
	"fmt"
	"github.com/kapilpatwa93/advent-of-code-2020/util"
	"os"
	"path"
	"strconv"
	"strings"
	"testing"
)

func getDay1Input() ([]int, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("something went wrong")
	}
	data, err := util.ReadFile(path.Join(dir, "input1.txt"))
	if err != nil {
		return nil, fmt.Errorf("cannot read input1.txt")
	}
	var expenses []int
	for _, s := range strings.Split(string(data), "\n") {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("invalid input")
		}
		expenses = append(expenses, num)
	}
	return expenses, nil
}

func TestPart1(t *testing.T) {
	expenses, err := getDay1Input()
	if err != nil {
		t.Fatal(err)
	}
	res := Part1(expenses, 2020)
	if res != 793524 {
		t.Errorf("Wrong answer")
	}
}

func TestPart2(t *testing.T) {
	expenses, err := getDay1Input()
	if err != nil {
		t.Fatal(err)
	}
	res := Part2(expenses,2020)
	fmt.Println(res)
}
