package day1

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	data, err := ioutil.ReadFile("input1.txt")
	if err != nil {
		t.Fatalf("cannot read input1.txt file")
	}
	var expenses []int
	for _, s := range strings.Split(string(data), "\n") {
		num, err := strconv.Atoi(s)
		if err != nil {
			t.Fatalf("invalid input")
		}
		expenses = append(expenses, num)
	}
	res := Part1(expenses, 2020)
	if res != 793524 {
		t.Errorf("Wrong answer")
	}
}
