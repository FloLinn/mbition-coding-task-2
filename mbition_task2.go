package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	// setup an IntervalMerger and let it listen to a stream of Intervals
	merger := NewIntervalMerger()
	intervalStream := make(chan Interval, 1)
	go func(c chan Interval) {
		for interval := range c {
			fmt.Printf("insert: %v", interval)
			err := merger.MergeSingleInterval(interval)
			if err != nil {
				fmt.Println("ERROR: ", err)
			} else {
				fmt.Printf(" -> %v\n", merger.Result())
			}
		}
	}(intervalStream)

	ui := NewUI()
	ui.GetUserChoices()

	switch ui.MainChoice {
	case DEFAULT:
		runDefaultIntervals(intervalStream)
	case CUSTOM:
		runCustomIntervals(intervalStream, ui.Count)
	case INFINITE:
		runInfiniteIntervals(intervalStream)
	}

	// let other goroutines finish their logging
	time.Sleep(1 * time.Second)
}

func runDefaultIntervals(intervalStream chan Interval) {
	intervals := []Interval{
		Interval{25, 30},
		Interval{2, 19},
		Interval{14, 23},
		Interval{4, 8},
	}
	for _, interval := range intervals {
		intervalStream <- interval
	}
}

func runCustomIntervals(intervalStream chan Interval, count int) {
	sqrtCount := int(math.Sqrt(float64(count)))
	for i := 0; i < count; i++ {
		b := rand.Intn(count)
		a := b - rand.Intn(sqrtCount)
		if a < 0 {
			a = 0
		}
		intervalStream <- Interval{a, b}
	}
}

func runInfiniteIntervals(intervalStream chan Interval) {
	maxValue := 10000
	maxWidth := 100
	for {
		b := rand.Intn(maxValue)
		a := b - rand.Intn(maxWidth)
		if a < 0 {
			a = 0
		}
		intervalStream <- Interval{a, b}
		time.Sleep(10 * time.Millisecond)
	}
}
