package main

import (
	"fmt"
)

// В общем логика простая:
// На вход приходит 123
// 123 mod 10 = 3
// 123 / 10 = 12
// s = "3" + "" == "3"
// второй цикл
// 12 mod 10 = 2
// 12 / 10 = 1
// s = "2" + "3" == "23"
// Поставь отладку на цикле и все поймешь
func itoa(i int) (s string) {
	negative := i < 0
	if negative {
		i = 0 - i
	}
	if i == 0 {
		return "0"
	}
	for i > 0 {
		tmp := i % 10
		i = i / 10
		s = string('0'+tmp) + s
	}
	if negative {
		s = "-" + s
	}
	return s
}

func main() {
	type pair struct {
		i int
		s string
	}
	test := []pair{
		{3, "3"},
		{0, "0"},
		{22, "22"},
		{32432523, "32432523"},
		{-3, "-3"},
	}
	for _, t := range test {
		if t.s == itoa(t.i) {
			fmt.Printf("%d - %s\n", t.i, "OK")
		} else {
			fmt.Printf("%d - %s\n", t.i, "FAIL")
		}
	}
}
