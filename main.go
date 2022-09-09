package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {

	// Get x from command-line arguments and convert to int
	x, _ := strconv.Atoi(os.Args[1])

	fmt.Println("Question 1:")
	question1(x)
	fmt.Println("\nQuestion 2:")
	question2(x)

}

///----------------Question 1 Code---------------------///

// Helper function to sum a slice of integers.
// Will be used in the synchronous summation
func sumSync(s []int) int {
	sum := 0
	for i := 0; i < cap(s); i++ {
		sum = sum + s[i]
	}
	return sum
}

// Helper function to sum a slice of integers.
// Will be usd in the asynchronous summation
func sumAsync(s []int, c chan int) {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum = sum + s[i]
	}
	// Use the output channel to send the sum to c
	c <- sum
}

func question1(x int) {

	s := make([]int, x)

	// Generate x random integers
	for i := 0; i < x; i++ {
		num := rand.Intn(1000)
		s[i] = num
	}

	// Synchronous Summation will loop over each element in the slice & print the sum.
	start := time.Now()
	sum := sumSync(s)
	end := time.Now()
	fmt.Println("Synchronous Summation Time:", end.Sub(start))
	fmt.Println(sum)

	// Asycnvhronous Summation will split the slice into two halves.
	// Each halve will be passed to a goroutine
	// Lines 61-64 where closely adopted from "https://go.dev/tour/concurrency/2"
	start = time.Now()
	c := make(chan int)
	// Begin to go routines to each sum over half of the slice
	go sumAsync(s[:cap(s)/2], c)
	go sumAsync(s[cap(s)/2:], c)
	s1, s2 := <-c, <-c // recieve each half sum through the channel
	sum = s1 + s2      // compute total sum
	end = time.Now()
	fmt.Println("Asynchronous Summation Time:", end.Sub(start))
	fmt.Println(sum)
}

///----------------Question 2 Code---------------------///

func question2(x int) {

	s := make([]int, x)

	// Generate x random integers
	for i := 0; i < x; i++ {
		num := rand.Intn(1000)
		s[i] = num
	}

	// Sort the slice using Slice from the sort library.
	start := time.Now()
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] }) // less function is closely adopted from example in "https://pkg.go.dev/sort#Slice"
	end := time.Now()
	fmt.Println("'sort.Slice' Run Time:", end.Sub(start))

	// Sort the slice using Slice from the sort library.
	start = time.Now()
	sort.SliceStable(s, func(i, j int) bool { return s[i] < s[j] }) // less function is closely adopted from example in "https://pkg.go.dev/sort#Slice"
	end = time.Now()
	fmt.Println("'sort.SliceStable' Run Time:", end.Sub(start))
}
