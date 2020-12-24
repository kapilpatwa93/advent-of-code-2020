package day3

import (
	"fmt"
	"github.com/kapilpatwa93/advent-of-code-2020/util"
	"os"
	"path"
	"strings"
	"testing"
)

func getInput() ([][]string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("something went wrong")
	}
	data, err := util.ReadFile(path.Join(dir, "input.txt"))
	if err != nil {
		return nil, fmt.Errorf("cannot read input.txt")
	}
	var areaMap [][]string
	for _, line := range strings.Split(string(data), "\n") {
		row := strings.Split(line, "")
		areaMap = append(areaMap, row)
	}
	return areaMap, nil
}

func TestPart1(t *testing.T) {
	areaMap, err := getInput()
	if err != nil {
		t.Fatal(err)
	}
	rightInc := 3
	bottomInc := 1
	tree := "#"
	res := Part1(areaMap, rightInc, bottomInc, tree)
	fmt.Println(res)
}
