package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"math/rand"
	"sort"
)

func question1() {

	// Get x from command-line arguments and convert to int
	x, _ := strconv.Atoi(os.Args[1])

	// Start Timer for Array
	arrStart := time.Now()
	// Initilize Array Variable
	var a [100]int
	// Iterate from 0 to x and declare each Array index
	for i :=0; i < x; i++ {
		a[i] = i
	}
	// End timer for Array
	arrEnd := time.Now()
	arrTime := arrEnd.Sub(arrStart)
	// Print Results
	fmt.Println("Array Time:", arrTime)
	


	// Start Timer for Slice
	slcStart := time.Now()
	// Initilize Map Variable
	s := make([]int, x)
	// Iterate from 0 to x and declare each slice index
	for i :=0; i < x; i++ {
		s[i] = i
	}
	// End timer for Map
	slcEnd := time.Now()
	slcTime := slcEnd.Sub(slcStart)
	// Print Results
	fmt.Println("Slice Time:", slcTime)


	// Start Timer for Map
	mapStart := time.Now()
	// Initilize Map Variable
	m := make(map[int]int)
	// Iterate from 0 to x and declare each map field
	for i :=0; i < x; i++ {
		m[i] = i
	}
	// End timer for Map
	mapEnd := time.Now()
	mapTime := mapEnd.Sub(mapStart)
	// Print Results
	fmt.Println("Map Time:", mapTime)


	//Increment Each Array Element
	arrStart = time.Now()
	for i :=0; i < x; i++ {
		a[i] = a[i] + 1
	}
	arrEnd = time.Now()
	arrTime = arrEnd.Sub(arrStart)
	fmt.Println("Array Increment Time:", arrTime)

	//Increment Each Slice Element
	slcStart = time.Now()
	for i :=0; i < x; i++ {
		s[i] = s[i] + 1
	}
	slcEnd = time.Now()
	slcTime = slcEnd.Sub(slcStart)
	fmt.Println("Slice Increment Time:", slcTime)

	//Increment Each Map Element's Field
	mapStart = time.Now()
	for i :=0; i < x; i++ {
		s[i] = s[i] + 1
	}
	mapEnd = time.Now()
	mapTime = mapEnd.Sub(mapStart)
	fmt.Println("Map Increment Time:", mapTime)
}

func question2() {

	// Get x from command-line arguments and convert to int
	x, _ := strconv.Atoi(os.Args[1])

	var a [100]int
	s := make([]int, x)

	

	// Generate x random integers
	for i :=0; i < x; i++ {
		num := rand.Intn(1000)
		a[i] = num
		s[i] = num
	}
	
	// Sort and time the array
	arrStart := time.Now()
	sort.Ints(a[:x])
	arrEnd := time.Now()
	fmt.Println("Array Sort Time:", arrEnd.Sub(arrStart))

	// Sort and time the slice
	slcStart := time.Now()
	sort.Ints(s)
	slcEnd := time.Now()
	fmt.Println("Slice Sort Time:", slcEnd.Sub(slcStart))

	// Stable and time the array
	arrStart = time.Now()
	sort.SliceStable(a[:x], func(i, j int) bool {return a[i] < a[j]})
	arrEnd = time.Now()
	fmt.Println("Array Sort Time:", arrEnd.Sub(arrStart))

	// Stable and time the slice
	slcStart = time.Now()
	sort.SliceStable(s[:x], func(i, j int) bool {return s[i] < s[j]})
	slcEnd = time.Now()
	fmt.Println("Slice Sort Time:", slcEnd.Sub(slcStart))
}

func main() {
	fmt.Println("Question 1:")
	question1()

	fmt.Println("Question 2:")
	question2()

}