package main

import (
	"fmt"
)

func main() {
	a := []string{"apple", "orange", "lemon", "apple", "vine"}

	str, a, err := delete_strings(a, "apple")
	fmt.Println(str) // => "apple"
	fmt.Println(a)   // => "[orange lemon vine]"
	fmt.Println(err) // => "<nil>"

	str, a, err = delete_strings(a, "apple")
	fmt.Println(str) // => ""
	fmt.Println(a)   // => "[orange lemon vine]"
	fmt.Println(err) // => "Couldn't find"
}

func delete_strings(slice []string, s string) (string, []string, error) {
	ret := make([]string, len(slice))
	i := 0
	for _, x := range slice {
		if s != x {
			ret[i] = x
			i++
		}
	}
	if len(ret[:i]) == len(slice) {
		return "", slice, fmt.Errorf("couldn't find")
	}
	return s, ret[:i], nil
}
