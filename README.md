# courseraGoC3Week3Assignment
Assigment Week #3- Concurrency in Go - THREADS IN GO - Course https://www.coursera.org/learn/golang-concurrency/home/week/3


## Assigment Week #3 - Concurrency in Go - THREADS IN GO  

- [Assignment](https://www.coursera.org/learn/golang-concurrency/peer/meeiu/module-3-activity)


- [Course: Concurrency in Go](https://www.coursera.org/learn/golang-concurrency/home/welcome)
  
## Instructions

The goal of this activity is to explore the use of threads by creating a program for sorting integers that uses four goroutines to create four sub-arrays and then merge the arrays into a single array.

Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.

## My assignment

Source code at `concurrentintsort/concurrentintsort.go`

## Sample compilation and first test execution

```sh
devel1@vbxdeb10mate:~$ cd $GOPATH/src/github.com/pssslearning/courseraGoC3Week3Assignment/concurrentintsort/
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/courseraGoC3Week3Assignment/concurrentintsort$ go build concurrentintsort.go 
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/courseraGoC3Week3Assignment/concurrentintsort$ ./concurrentintsort 

==================================================================================
Please enter up to a maximum of 40 numbers but also not less than a minimum of 12.
==================================================================================
Please input integer number #1 (type 'END' to end entering more numbers): -1
Please input integer number #2 (type 'END' to end entering more numbers): 234
Please input integer number #3 (type 'END' to end entering more numbers): 453
Please input integer number #4 (type 'END' to end entering more numbers): -1000
Please input integer number #5 (type 'END' to end entering more numbers): 257
Please input integer number #6 (type 'END' to end entering more numbers): 90
Please input integer number #7 (type 'END' to end entering more numbers): 58
Please input integer number #8 (type 'END' to end entering more numbers): 45
Please input integer number #9 (type 'END' to end entering more numbers): -987
Please input integer number #10 (type 'END' to end entering more numbers): -34
Please input integer number #11 (type 'END' to end entering more numbers): 0
Please input integer number #12 (type 'END' to end entering more numbers): 12
Please input integer number #13 (type 'END' to end entering more numbers): 45
Please input integer number #14 (type 'END' to end entering more numbers): -234
Please input integer number #15 (type 'END' to end entering more numbers): 125
Please input integer number #16 (type 'END' to end entering more numbers): 789
Please input integer number #17 (type 'END' to end entering more numbers): 123
Please input integer number #18 (type 'END' to end entering more numbers): END

------------------------------------------------------------------------------------------

Total number of integer entered [17]
Initial slice of integers: [-1 234 453 -1000 257 90 58 45 -987 -34 0 12 45 -234 125 789 123]
Initial slice of integers to be sorted in Thread #1: [-1 234 453 -1000]
Initial slice of integers to be sorted in Thread #2: [257 90 58 45]
Initial slice of integers to be sorted in Thread #3: [-987 -34 0 12]
Initial slice of integers to be sorted in Thread #4: [45 -234 125 789 123]
------------------------------------------------------------------------------------------
On thread #[4] - Slice received [45 -234 125 789 123]
On thread #[4] - Sorting ...
On thread #[4] - Slice sorted [-234 45 123 125 789]. Exiting thread...
On thread #[3] - Slice received [-987 -34 0 12]
On thread #[3] - Sorting ...
On thread #[3] - Slice sorted [-987 -34 0 12]. Exiting thread...
On thread #[2] - Slice received [257 90 58 45]
On thread #[2] - Sorting ...
On thread #[2] - Slice sorted [45 58 90 257]. Exiting thread...
On thread #[1] - Slice received [-1 234 453 -1000]
On thread #[1] - Sorting ...
On thread #[1] - Slice sorted [-1000 -1 234 453]. Exiting thread...
------------------------------------------------------------------------------------------
Resulting slice of integers sorted in Thread #1: [-1000 -1 234 453]
Resulting slice of integers sorted in Thread #2: [45 58 90 257]
Resulting slice of integers sorted in Thread #3: [-987 -34 0 12]
Resulting slice of integers sorted in Thread #4: [-234 45 123 125 789]

After merging all that slices, the initial slice of integers now contains: [-1000 -1 234 453 45 58 90 257 -987 -34 0 12 -234 45 123 125 789]

The resulting slice of integers, once sorted, now contains: [-1000 -987 -234 -34 -1 0 12 45 45 58 90 123 125 234 257 453 789]
```
