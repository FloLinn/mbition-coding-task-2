package main

import (
	"fmt"

	avl_tree "github.com/emirpasic/gods/trees/avltree"
)

// Interval from A to B (included)
type Interval struct {
	A int
	B int
}

// IntervalMerger merges intervals :).
type IntervalMerger struct {
	// note: The Tree Values are Pointers to Intervals.
	// So, we can alter them and avoid re-insertion of nodes when possible
	resultTree *avl_tree.Tree
}

// MergeSingleInterval merges a single Interval into this struct
func (im IntervalMerger) MergeSingleInterval(toMerge Interval) error {
	// we find the highest interval that intersects with toMerge
	// Then we merge it and all smaller intervals that intersect with it
	// once we encounter interval does not intersect toMerge, we are done
	// Going 'downwards' is necessary since we order the result by start-value

	if toMerge.A > toMerge.B {
		return fmt.Errorf("Interval %v with negative length detected", toMerge)
	}

	// find the last Interval in result then overlaps with toMerge
	nodeToCheck, found := im.resultTree.Floor(toMerge.B)

	if !found {
		im.resultTree.Put(toMerge.A, &toMerge)
		return nil
	}

	for nodeToCheck != nil {
		nodeInterval := nodeToCheck.Value.(*Interval)
		// we know: nodeInterval.A <= toMerge.B
		if nodeInterval.B < toMerge.A {
			// no overlap
			im.resultTree.Put(toMerge.A, &toMerge)
			nodeToCheck = nil // done
		} else if nodeInterval.A <= toMerge.A {
			// we don't need to change the tree structure
			if nodeInterval.B < toMerge.B {
				// expand nodeInterval to the right
				nodeInterval.B = toMerge.B
			}
			//else: nodeInterval contains toMerge completely -> no changes needed
			nodeToCheck = nil // done
		} else {
			// nodeInterval.A > toMerge.A
			// toMerge can affect the previous interval in the result
			// so, we include the current nodeInterval into toMerge,
			// remove nodeToCheck and continue this loop with the previous node
			if toMerge.B < nodeInterval.B {
				toMerge = Interval{toMerge.A, nodeInterval.B}
			} else {
				toMerge = Interval{toMerge.A, toMerge.B}
			}
			nodeToCheck = nodeToCheck.Prev()
			im.resultTree.Remove(nodeInterval.A)
			if nodeToCheck == nil {
				// no prev node present .. just insert toMerge
				im.resultTree.Put(toMerge.A, &toMerge)
			}
		}
	}
	return nil
}

// Result returns all merged intervals
func (im IntervalMerger) Result() (result []Interval) {
	intervalPointers := im.resultTree.Values()
	// don't return pointers, return by value (prevent unwanted sideeffects)
	result = make([]Interval, len(intervalPointers))
	for i, v := range intervalPointers {
		result[i] = *(v.(*Interval))
	}
	return
}

// PrintTree is only used for debugging
func (im IntervalMerger) PrintTree() {
	fmt.Println(im.resultTree.String())
	iterator := im.resultTree.Iterator()
	for iterator.Next() {
		fmt.Println(iterator.Value())
	}
}

// NewIntervalMerger constructor
func NewIntervalMerger() IntervalMerger {
	return IntervalMerger{avl_tree.NewWithIntComparator()}
}
