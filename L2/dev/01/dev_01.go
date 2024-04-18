package dev01

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func GetTime() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	fmt.Print(time)
}
