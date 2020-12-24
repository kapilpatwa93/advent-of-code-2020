package day2

import (
	"fmt"
	"github.com/kapilpatwa93/advent-of-code-2020/util"
	"os"
	"path"
	"strconv"
	"strings"
	"testing"
)

func getDay2Input() ([]string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("something went wrong")
	}
	data, err := util.ReadFile(path.Join(dir, "input.txt"))
	if err != nil {
		return nil, fmt.Errorf("cannot read input.txt")
	}
	return strings.Split(string(data), "\n"), nil
}

func getDay2InputPart1() ([]Policy1, error) {
	lines, err := getDay2Input()
	if err != nil {
		return nil, err
	}

	var policies []Policy1
	for _, line := range lines {
		splits := strings.Split(line, " ")
		counts := strings.Split(splits[0], "-")
		leastCount, err := strconv.Atoi(counts[0])
		if err != nil {
			return nil, fmt.Errorf("invalid input")
		}
		mostCount, err := strconv.Atoi(counts[1])
		policies = append(policies, Policy1{
			LeastCount: leastCount,
			MostCount:  mostCount,
			Password:   splits[2],
			Character:  splits[1][:1],
		})
	}
	return policies, nil
}

func TestPart1(t *testing.T) {
	policies, err := getDay2InputPart1()
	if err != nil {
		t.Fatal(err)
	}
	res := Part1(policies)
	if res != 636 {
		t.Errorf("Wrong answer")
	}
}

func getDay2InputPart2() ([]Policy2, error) {
	lines, err := getDay2Input()
	if err != nil {
		return nil, err
	}
	var policies []Policy2
	for _, line := range lines {
		splits := strings.Split(line, " ")
		counts := strings.Split(splits[0], "-")
		position1, err := strconv.Atoi(counts[0])
		if err != nil {
			return nil, fmt.Errorf("invalid input")
		}
		position2, err := strconv.Atoi(counts[1])
		policies = append(policies, Policy2{
			Position1: position1,
			Position2: position2,
			Password:  splits[2],
			Character: splits[1][:1],
		})
	}
	return policies, nil
}

func TestPart2(t *testing.T) {
	policies, err := getDay2InputPart2()
	if err != nil {
		t.Fatal(err)
	}
	res := Part2(policies)
	if res != 588 {
		t.Errorf("Wrong answer")
	}
}
