package main

import "fmt"

func main() {
	fmt.Println(len("é"))
	fmt.Println(len([]rune("é")))
	fmt.Println(unique("ééééé"))
	fmt.Println(unique("éèplo!@"))
	fmt.Println(unique2("ééééé"))
	fmt.Println(unique2("éèplo!@"))
}

func unique(s string) bool {
	present := make(map[rune]bool)
	for _, r := range s {
		if present[r] {
			return false
		}
		present[r] = true
	}

	return true
}

func unique2(s string) bool {
	for i, r := range s {
		if i == len([]rune(s))-1 {
			return true
		}
		for _, rr := range ([]rune(s))[i+1:] {
			if r == rr {
				return false
			}
		}
	}

	return true
}
