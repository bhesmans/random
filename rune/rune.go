package main

import "fmt"

func main() {
	fmt.Println(len("é"))
	fmt.Println(len([]rune("é")))
	fmt.Println(unique("ééééé"))
	fmt.Println(unique("éèplo!@"))
	fmt.Println(unique2("ééééé"))
	fmt.Println(unique2("éèplo!@"))
	fmt.Println(permutation("azerty", "zrteya"))
	fmt.Println(permutation("azerty", "zrtzya"))
	fmt.Println(permutation("azerty", "zrteyaé"))
	fmt.Println(permutation("éé@èè", "@èéèé"))
	fmt.Println(escapeSpaces(" plo pla plu ! "))
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

func permutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	r1 := []rune(s1)
	r2 := []rune(s2)

	balance := make(map[rune]int)

	for i := range r1 {
		balance[r1[i]] = balance[r1[i]] + 1
		balance[r2[i]] = balance[r2[i]] - 1
	}

	for _, v := range balance {
		if v != 0 {
			return false
		}
	}

	return true
}

func runeCount(s string, r rune) int {
	count := 0

	for _, rr := range s {
		if rr == r {
			count++
		}
	}

	return count
}

func escapeSpaces(s string) string {
	sc := runeCount(s, ' ')
	r := []rune(s)
	slen := len(r)
	res := make([]rune, slen+(sc*2))

	for i, j := 0, 0; i < slen; i++ {
		if r[i] != ' ' {
			res[j] = r[i]
			j++
		} else {
			res[j] = '%'
			res[j+1] = '2'
			res[j+2] = '0'
			j += 3
		}
	}

	return string(res)
}
