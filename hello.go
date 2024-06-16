package main

import (
	"errors"
)

func main() {
	//letters := [...]string{"z", "x", "c", "d", "v"}
	//combinations := []string{}
	//for i, letter := range letters {
	//	c := []string{}
	//	c = append(c, letter)
	//	for _, letter := range letters[i:1] {
	//		c = append(c, letter)
	//	}
	//	combinations = append(combinations, c)
	//	fmt.Printf("Created combintation: %s", c)
	//}
	//
	//example := []string{"a", "b", "c"}
	//result, err := RemoveElement(example, 1)
	//if err != nil {
	//	fmt.Printf("error: %s", err)
	//}
	//if len(result) < 1 {
	//	fmt.Printf("stuff2")
	//}
	//if result[0] == "a" {
	//	fmt.Printf("stuff")
	//}
}

type CombinationMaker interface {
	CreateCombinations(a []string)
	RemoveElement([]string, int) (result []string, err error)
}

func CreateCombinations(a []string) (result [][]string, err error) {
	result = make([][]string, len(a))
	result[0] = a
	return result, nil
}

func makeCombinations(a []string, prefix string, r []string) (result []string, err error) {
	if len(a) == 1 {
		prefix = prefix + a[0]
		r = append(r, prefix)
		return r, nil
	}
	if len(a) == 0 {
		return nil, err
	}
	if prefix == "" {
		for _, p := range a {
			prefix = p
			// remove prefix from array and then makeCombinations

		}
	}
	return nil, err
}

func RemoveElement(a []string, i int) (result []string, err error) {
	if a == nil {
		err := errors.New("Cannot split an empty array")
		return nil, err
	}
	if len(a) < 2 {
		err := errors.New("Cannot split an array with a single item")
		return nil, err
	}
	if i >= len(a) {
		err := errors.New("Index out of bounds")
		return nil, err
	}
	start := a[0:i]
	end := a[i+1:]
	result = append(start, end...)
	return result, nil
}
