package main

import (
	"fmt"
	"testing"
)

// #################################################
// BASIC INITIALIZATION TEST

func TestIntervalMergerEmpty(t *testing.T) {
	merger := NewIntervalMerger()
	result := merger.Result()
	if len(result) != 0 {
		t.Errorf("non empty result on freshly initialized IntervalMerger: %v", result)
	}
}

func TestIntervalMergerFirstInterval(t *testing.T) {
	merger := NewIntervalMerger()
	err := merger.MergeSingleInterval(Interval{1, 2})
	if err != nil {
		t.Errorf("Got unexpected error %v", err)
	}
	result := merger.Result()
	if actual, expected := fmt.Sprintf("%v", result), "[{1 2}]"; actual != expected {
		t.Errorf("wrong result on first merged Interval: got %v, expected %v", actual, expected)
	}
}

func TestIntervalMergerBadInterval(t *testing.T) {
	merger := NewIntervalMerger()
	err := merger.MergeSingleInterval(Interval{2, 1})
	if err == nil {
		t.Errorf("no Error was thrown on bad interval")
	}
}

// #################################################
// MERGE TEST FOR SINGLE INTERVALS

func TestIntervalMergerInsertDisjunctRight(t *testing.T) {
	merger := NewIntervalMerger()
	merger.MergeSingleInterval(Interval{10, 11})
	merger.MergeSingleInterval(Interval{15, 16})
	result := merger.Result()
	if actual, expected := fmt.Sprintf("%v", result), "[{10 11} {15 16}]"; actual != expected {
		t.Errorf("wrong result on second merged Interval: got %v, expected %v", actual, expected)
	}
}

func TestIntervalMergerInsertOverlapRight(t *testing.T) {
	merger := NewIntervalMerger()
	merger.MergeSingleInterval(Interval{10, 15})
	merger.MergeSingleInterval(Interval{15, 16})
	result := merger.Result()
	if actual, expected := fmt.Sprintf("%v", result), "[{10 16}]"; actual != expected {
		t.Errorf("wrong result on insert Interval: got %v, expected %v", actual, expected)
	}
	merger.MergeSingleInterval(Interval{13, 20})
	result = merger.Result()
	if actual, expected := fmt.Sprintf("%v", result), "[{10 20}]"; actual != expected {
		t.Errorf("wrong result on insert Interval: got %v, expected %v", actual, expected)
	}
}

func TestIntervalMergerInsertCovered(t *testing.T) {
	merger := NewIntervalMerger()
	merger.MergeSingleInterval(Interval{10, 15})
	merger.MergeSingleInterval(Interval{12, 14})
	result := merger.Result()
	if actual, expected := fmt.Sprintf("%v", result), "[{10 15}]"; actual != expected {
		t.Errorf("wrong result on insert Interval: got %v, expected %v", actual, expected)
	}
	merger.MergeSingleInterval(Interval{10, 15})
	result = merger.Result()
	if actual, expected := fmt.Sprintf("%v", result), "[{10 15}]"; actual != expected {
		t.Errorf("wrong result on insert Interval: got %v, expected %v", actual, expected)
	}
}

func TestIntervalMergerInsertCovering(t *testing.T) {
	merger := NewIntervalMerger()
	merger.MergeSingleInterval(Interval{10, 15})
	merger.MergeSingleInterval(Interval{8, 16})
	result := merger.Result()
	if actual, expected := fmt.Sprintf("%v", result), "[{8 16}]"; actual != expected {
		t.Errorf("wrong result on insert Interval: got %v, expected %v", actual, expected)
	}
}

func TestIntervalMergerInsertOverlapLeft(t *testing.T) {
	merger := NewIntervalMerger()
	merger.MergeSingleInterval(Interval{10, 15})
	merger.MergeSingleInterval(Interval{8, 14})
	result := merger.Result()
	if actual, expected := fmt.Sprintf("%v", result), "[{8 15}]"; actual != expected {
		t.Errorf("wrong result on insert Interval: got %v, expected %v", actual, expected)
	}
	merger.MergeSingleInterval(Interval{2, 8})
	result = merger.Result()
	if actual, expected := fmt.Sprintf("%v", result), "[{2 15}]"; actual != expected {
		t.Errorf("wrong result on insert Interval: got %v, expected %v", actual, expected)
	}
}

func TestIntervalMergerInsertDisjunctLeft(t *testing.T) {
	merger := NewIntervalMerger()
	merger.MergeSingleInterval(Interval{10, 15})
	merger.MergeSingleInterval(Interval{3, 7})
	result := merger.Result()
	if actual, expected := fmt.Sprintf("%v", result), "[{3 7} {10 15}]"; actual != expected {
		t.Errorf("wrong result on insert Interval: got %v, expected %v", actual, expected)
	}
}

// #################################################
// MERGE TEST FOR INTERVALS OF LENGTH 0

func TestIntervalMergerInsertEmptyInterval(t *testing.T) {
	merger := NewIntervalMerger()
	merger.MergeSingleInterval(Interval{10, 10})
	result := merger.Result()
	if actual, expected := fmt.Sprintf("%v", result), "[{10 10}]"; actual != expected {
		t.Errorf("wrong result on insert empty Interval: got %v, expected %v", actual, expected)
	}
	merger.MergeSingleInterval(Interval{10, 10})
	result = merger.Result()
	if actual, expected := fmt.Sprintf("%v", result), "[{10 10}]"; actual != expected {
		t.Errorf("wrong result on insert empty Interval: got %v, expected %v", actual, expected)
	}
	merger.MergeSingleInterval(Interval{9, 9})
	result = merger.Result()
	if actual, expected := fmt.Sprintf("%v", result), "[{9 9} {10 10}]"; actual != expected {
		t.Errorf("wrong result on insert empty Interval: got %v, expected %v", actual, expected)
	}
	merger.MergeSingleInterval(Interval{9, 10})
	result = merger.Result()
	if actual, expected := fmt.Sprintf("%v", result), "[{9 10}]"; actual != expected {
		t.Errorf("wrong result on insert empty Interval: got %v, expected %v", actual, expected)
	}
	merger.MergeSingleInterval(Interval{9, 9})
	result = merger.Result()
	if actual, expected := fmt.Sprintf("%v", result), "[{9 10}]"; actual != expected {
		t.Errorf("wrong result on insert empty Interval: got %v, expected %v", actual, expected)
	}
	merger.MergeSingleInterval(Interval{10, 10})
	result = merger.Result()
	if actual, expected := fmt.Sprintf("%v", result), "[{9 10}]"; actual != expected {
		t.Errorf("wrong result on insert empty Interval: got %v, expected %v", actual, expected)
	}
}

// #################################################
// MERGE TESTS INSERT COVERING MULTIPLE INTERVALS

func TestIntervalMergerInsertCoveringMultipleIntervals_1(t *testing.T) {
	merger := NewIntervalMerger()
	merger.MergeSingleInterval(Interval{10, 15})
	merger.MergeSingleInterval(Interval{20, 30})
	merger.MergeSingleInterval(Interval{35, 40})
	merger.MergeSingleInterval(Interval{5, 45})
	result := merger.Result()
	if actual, expected := fmt.Sprintf("%v", result), "[{5 45}]"; actual != expected {
		t.Errorf("wrong result on insert big Interval: got %v, expected %v", actual, expected)
	}
}

func TestIntervalMergerInsertCoveringMultipleIntervals_2(t *testing.T) {
	merger := NewIntervalMerger()
	merger.MergeSingleInterval(Interval{10, 15})
	merger.MergeSingleInterval(Interval{20, 30})
	merger.MergeSingleInterval(Interval{35, 40})
	merger.MergeSingleInterval(Interval{15, 35})
	result := merger.Result()
	if actual, expected := fmt.Sprintf("%v", result), "[{10 40}]"; actual != expected {
		t.Errorf("wrong result on insert big Interval: got %v, expected %v", actual, expected)
	}
}

func TestIntervalMergerInsertCoveringMultipleIntervals_3(t *testing.T) {
	merger := NewIntervalMerger()
	merger.MergeSingleInterval(Interval{10, 15})
	merger.MergeSingleInterval(Interval{20, 30})
	merger.MergeSingleInterval(Interval{35, 40})
	merger.MergeSingleInterval(Interval{20, 38})
	result := merger.Result()
	if actual, expected := fmt.Sprintf("%v", result), "[{10 15} {20 40}]"; actual != expected {
		t.Errorf("wrong result on insert big Interval: got %v, expected %v", actual, expected)
	}
}

// #################################################
// MERGE TESTS RESULT DECOUPLED FROM INPUT OUTPUT (no POINTERS)

func TestIntervalMergerDecoupling(t *testing.T) {
	merger := NewIntervalMerger()
	interval := Interval{10, 15}
	merger.MergeSingleInterval(interval)
	// change input
	interval.A = 1
	interval.B = 2
	result := merger.Result()
	if actual, expected := fmt.Sprintf("%v", result), "[{10 15}]"; actual != expected {
		t.Errorf("changing input has sideeffects on result: got %v, expected %v", actual, expected)
	}
	// change output
	result[0].A = 1
	result[0].B = 2
	result = merger.Result()
	if actual, expected := fmt.Sprintf("%v", result), "[{10 15}]"; actual != expected {
		t.Errorf("changing input has sideeffects on result: got %v, expected %v", actual, expected)
	}
}
