package main

import "fmt"

func main() {
	letters := [...]string{"z", "x", "c", "d", "v"}
	combinations := []string{}
	for i, letter := range letters {
		c := []string{}
		c = append(c, letter)
		for _, letter range letters[i:-1] {
			c = append(c, letter)
		}
		combinations = append(combinations, c)
		fmt.Printf("Created combintation: %s", c)
	}
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

}

// TODO: create a function that returns an array with the index removed
func removeElement(a []string, i int) (result []string, err error){
	if a == nil || len(a) < 2 {
		return nil, err
	}
	start := a[:i-1]
	end := a[i+1:]
	result = append(start, end)
	return result
}
