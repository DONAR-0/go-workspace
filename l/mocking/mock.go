package mocking

import (
	"fmt"
	"io"
	"iter"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const finalWord = "Go!"

// NOTE Comment
// const countDownStart = 3
const write = "write"
const sleep = "sleep"

// func CountDown(out io.Writer, sleeper Sleeper) {
// 	for i := countDownStart; i > 0; i-- {
// 		_, _ = fmt.Fprintln(out, i)
//
// 		sleeper.Sleep()
// 	}
//
// 	_, _ = fmt.Fprint(out, finalWord)
// }

func CountDown(out io.Writer, sleeper Sleeper) {
	for i := range countDownFrom(3) {
		_, _ = fmt.Fprintln(out, i)

		sleeper.Sleep()
	}

	_, _ = fmt.Fprint(out, finalWord)
}

func countDownFrom(from int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := from; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}
