package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {

	slicints := make([]int, 40)
	cursor := 0
	fmt.Println("\n==================================================================================")
	fmt.Println("Please enter up to a maximum of 40 numbers but also not less than a minimum of 12.")
	fmt.Println("==================================================================================")

	scanner := bufio.NewScanner(os.Stdin)

	// Loop to request the user to enter a list of integers
	for {

		fmt.Printf("Please input integer number #%d (type 'END' to end entering more numbers): ", cursor+1)
		scanner.Scan() // Get the next value

		inputData := strings.ToUpper(strings.TrimSpace(scanner.Text()))

		if inputData == "END" {
			if cursor >= 12 {
				break
			} else {
				fmt.Println("Remember: The introduction of at least 12 integers is required for the program to perform its function. Its execution is aborted. Please try again.")
				os.Exit(1)
			}
		}

		// Convert to an integer
		val, err := strconv.Atoi(inputData)

		if nil == err {
			// Insert value into the Slice
			slicints[cursor] = val
			cursor++
			if cursor >= len(slicints) {
				break
			}
		} else {
			// Print an error message
			fmt.Println("Value was not an integer (nor 'END' condition). Please try again.")
			continue
		}

	}

	fmt.Println("\n------------------------------------------------------------------------------------------")
	fmt.Printf("\nTotal number of integer entered [%d]\n", cursor)
	fmt.Printf("Initial slice of integers: %v\n", slicints[:cursor])

	// Get the 4 slices
	divider := cursor / 4

	slic01 := slicints[:divider]
	slic02 := slicints[divider : divider*2]
	slic03 := slicints[divider*2 : divider*3]
	slic04 := slicints[divider*3 : cursor]

	fmt.Printf("Initial slice of integers to be sorted in Thread #%d: %v\n", 1, slic01)
	fmt.Printf("Initial slice of integers to be sorted in Thread #%d: %v\n", 2, slic02)
	fmt.Printf("Initial slice of integers to be sorted in Thread #%d: %v\n", 3, slic03)
	fmt.Printf("Initial slice of integers to be sorted in Thread #%d: %v\n", 4, slic04)

	fmt.Println("------------------------------------------------------------------------------------------")

	// Do the concurrent work

	var wg sync.WaitGroup
	wg.Add(4)
	go sliceSort(&wg, 1, slic01)
	go sliceSort(&wg, 2, slic02)
	go sliceSort(&wg, 3, slic03)
	go sliceSort(&wg, 4, slic04)
	wg.Wait()

	fmt.Println("------------------------------------------------------------------------------------------")
	// Show results
	fmt.Printf("Resulting slice of integers sorted in Thread #%d: %v\n", 1, slic01)
	fmt.Printf("Resulting slice of integers sorted in Thread #%d: %v\n", 2, slic02)
	fmt.Printf("Resulting slice of integers sorted in Thread #%d: %v\n", 3, slic03)
	fmt.Printf("Resulting slice of integers sorted in Thread #%d: %v\n", 4, slic04)

	fmt.Printf("\nAfter merging all that slices, the initial slice of integers now contains: %v\n", slicints[:cursor])
	sort.Ints(slicints[:cursor])
	fmt.Printf("\nThe resulting slice of integers, once sorted, now contains: %v\n", slicints[:cursor])

}

// Sorting routine
func sliceSort(wg *sync.WaitGroup, threadID int, theSlice []int) {
	fmt.Printf("On thread #[%d] - Slice received %v\n", threadID, theSlice)
	fmt.Printf("On thread #[%d] - Sorting ...\n", threadID)
	sort.Ints(theSlice)
	fmt.Printf("On thread #[%d] - Slice sorted %v. Exiting thread...\n", threadID, theSlice)
	wg.Done()
}
