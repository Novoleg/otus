package main

import (
	"fmt"

	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Текущее время:", time)
}
