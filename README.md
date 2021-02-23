# mbition coding task 2

### how to run
This program is written in go.

If you have go installed, use
`go run mbition_task2` for the program
and
`go test` to run the test suite.

### general notice

+ Although I'm german, I'll write this doc in english.
+ I'm new to golang, but I want to learn it, so this thing is written in Go.

### basic properties of the problem

+ Merging just two intervals is trivial and can be done in constant time 
+ The merge function is commutative, we don't need a specific computation order
+ The merge function is idempotent, so merging the same interval multiple times does not affect the result.
+ Intervals can be sorted (e.g. by their start value)

let

`n == initial number of intervals`

`rn == maximum number of intervals in the result`

### in-place algorithm
We could implement this algorithm as in-place (the input itself is transformed),
which would require O(n^2) if the input is not being sorted initially (without proof here).

If we sort the input first, we then can use a sweep algorithm (just along the list indexes), which successively merges adjacent intervals.
Sorting the input requires O(n^log(n)) time and the sweep would take O(n).
So this algorithm would take at least **O(n^log(n)) time and O(n) space**.

(note: A Divide-and-Conquer solution is not practical here, without explanation)

### inline-algorithm
Another solution would be an inline-algorithm, in which intervals are pulled successively from the input and merged into the current result.
Here, we can be sure, that the current result has no overlapping intervals at any time.

We will keep the output sorted (by start of interval).
This way, we can find the 'insertion point' for the next pulled interval in O(log(rn)) time.

Inserting this interval can take up to O(rn) time, but since the result-size will shrink accordingly, this cost in amortized (on the whole input list) to O(1)
This algorithm therefor **takes O(n * log(rn)) time**.

The space consumtion depends on the structure of the result.
We need a dynamically changing, sorted structure, so I choose an AVL-Tree (tree with guaranteed performance) with a space consumtion of O(rn).
The needed **space** therefor is O(rn) + O(n) = **O(n)**.

In the worst-case, when the input has no intersecting intervals, n == rn,
the inline-algorithm has no performance advantage over the in-place algorithm.
In practice however, we can assume that rn will be quite smaller than n.
Also, we don't need the input as a whole in the first place, this could also be a stream.
In this case, we only need **O(rn) space**.

A disadvantage of an inline-algorithm is, that we cannot speed up by a parallel algorithm.

I assume that mbtion would rather like the ability of input streaming than spawning a lot of threads,
so I go with the **inline algorithm**.

### timesheet

+ The decision on the algorithm took 10 minutes. I know problems like this from my studies
+ writing the description above took about 1 hour. I think I deleted about the same amount of text again :)
+ writing the implementation of the algorithm took ... well ... about 4 hours! This was my first go app, so I stumbled over pointer errors, syntax, missing generics etc. But it was fun and I learned a lot.
+ writing the ui took again about 2 hours. Since you didn't ask for it, you can just ignore this. (I just had some fun learning go)
+ writing the tests and tidiying up took one hour. Go's testing framework is surprisingly easy.




