package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func Countdown(out io.Writer) {
	const finalWord = "Go!"
	const coundownStart = 3
	for i := coundownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(time.Second * 1)
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
