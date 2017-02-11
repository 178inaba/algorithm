package main

import "fmt"

type data struct {
	s, w string
	t    []int
}

func main() {
	d := newData("ABC ABCDAB ABCDABCDABDE", "ABCDABD")

	p := d.kmp()

	fmt.Println(d.s)
	for i := 0; i < p; i++ {
		fmt.Print(" ")
	}
	fmt.Println(d.w)
}

func newData(s, w string) data {
	return data{s: s, w: w, t: newTable(w)}
}

func newTable(w string) []int {
	t := make([]int, len(w))
	t[0], t[1] = -1, 0

	for i, j := 2, 0; i < len(w); {
		if w[i-1] == w[j] {
			t[i] = j + 1
			i++
			j++
		} else if j > 0 {
			j = t[j]
		} else {
			t[i] = 0
			i++
		}
	}

	return t
}

func (d data) kmp() int {
	sLen := len(d.s)
	for m, i := 0, 0; m+i < sLen; {
		if d.w[i] == d.s[m+i] {
			i++
			if i == len(d.w) {
				return m
			}
		} else {
			m = m + i - d.t[i]
			if i > 0 {
				i = d.t[i]
			}
		}
	}

	return sLen
}
