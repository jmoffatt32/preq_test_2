# Go Preq Test 2

## Instructions

Run `go build main.go` to build the executable file

Then run `./main x` to execute the program, where x is any integer greater than 0

## Expected Output

If we run `go build main.go && ./main 100000` we may recieve the output:

```
Question 1:
Synchronous Summation Time: 48.458µs
50062287
Asynchronous Summation Time: 64.208µs
50062287

Question 2:
'sort.Slice' Run Time: 7.064125ms
'sort.SliceStable' Run Time: 326.875µs
```

## Question 3

I do think that my prgoram follows the time analysis explained in the Go Package Documentation. As described in the docs, both `Sort` and `Stable` make 1 call to `data.Len` to determine the length of the slice. Additionally, both sort and stable make `O(n * log(n))` calls to `data.Less` to make comparsions between elements. They are differentiated by the number of swaps the make. `Sort` makes `O(n * log(n))`
swaps while `Stable` makes `O(n * log(n) * log(n))` swaps. We see that this results in a much larger time analysis for the `Sort` method, as it must more swaps as `n` increases.
