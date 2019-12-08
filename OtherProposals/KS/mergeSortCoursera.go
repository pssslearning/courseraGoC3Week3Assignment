package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {

	fmt.Print("Enter series of integer numbers: ")
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')

	if err != nil {
		log.Fatal("invalid input", err)
	}
	s, _ := parseInts(line)

	s1 := s[:len(s)/4]
	s2 := s[len(s)/4 : len(s)/2]
	s3 := s[len(s)/2 : len(s)*3/4]
	s4 := s[3*len(s)/4 : len(s)]

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("Sorting ", s1)
		sort.Ints(s1)
		defer wg.Done()
	}()
	wg.Add(1)
	go func() {
		fmt.Println("Sorting ", s2)
		sort.Ints(s2)
		defer wg.Done()
	}()
	wg.Add(1)
	go func() {
		fmt.Println("Sorting ", s3)
		sort.Ints(s3)
		defer wg.Done()
	}()
	wg.Add(1)
	go func() {
		fmt.Println("Sorting ", s4)
		sort.Ints(s4)
		defer wg.Done()
	}()
	wg.Wait()

	var result1 []int
	var result2 []int

	wg.Add(1)
	go func() {
		result1 = mergeSort(s1, s2)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		result2 = mergeSort(s3, s4)
		wg.Done()
	}()
	wg.Wait()

	final_result := mergeSort(result1, result2)

	fmt.Println("Sorted result ", final_result)

}

func mergeSort(s1 []int, s2 []int) []int {
	k, i, j := 0, 0, 0
	final_size := len(s1) + len(s2)
	final := make([]int, final_size, final_size)
	for i < len(s1) && j < len(s2) {
		if s1[i] < s2[j] {
			final[k] = s1[i]
			i++
			k++
		} else {
			final[k] = s2[j]
			j++
			k++
		}
	}
	for i < len(s1) {
		final[k] = s1[i]
		i++
		k++

	}

	for j < len(s2) {
		final[k] = s2[j]
		j++
		k++
	}
	return final
}

func parseInts(s string) ([]int, error) {
	var (
		fields  = strings.Fields(s)
		numbers = make([]int, len(fields))
		err     error
	)
	for i, f := range fields {
		numbers[i], err = strconv.Atoi(f)
		if err != nil {
			return nil, err
		}
	}
	return numbers, nil
}
