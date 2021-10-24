package mocking

import (
	"fmt"
	"io"
)

const countDownStart = 3

type Sleeper interface {
	Sleep()
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintf(writer, "%d\n", i)
	}
	sleeper.Sleep()
	fmt.Fprintf(writer, "Go!")
}
