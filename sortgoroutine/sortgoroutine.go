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
	var wg sync.WaitGroup	
	arr := GetInput()
	i := Divide(&arr)

	// call the four goroutines and wait for them to complete
	wg.Add(4)
	go SortArray(arr[:i[0]], &wg)
	go SortArray(arr[i[0]:i[1]], &wg)
	go SortArray(arr[i[1]:i[2]], &wg)
	go SortArray(arr[i[2]:], &wg)
	wg.Wait()

	merged := MergeArray(arr[:i[0]], arr[i[0]:i[1]],
		arr[i[1]:i[2]], arr[i[2]:], len(arr))

	fmt.Printf("Final sorted array %v\n", merged)
}

func GetInput() []int {
		// take input of integers from the user
		sli := make([]int, 0, 25)
		var input string
		fmt.Println("Please enter integers separated by space: ")
		
		// create a scanner and read the line of text
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input = scanner.Text()
	
		// convert string of space-separated integers to slice of integers
		fields := strings.Fields(input)
		for _, num := range fields {
			num, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println("Please enter space-separated integers!")
				continue
			}
			sli =append(sli, num)
		}
		return sli
}


// Divides the given array into 4 nearly equal parts and returns indices
func Divide(arr *[]int) *[3]int	{
	var i [3]int
	i[1] = len(*arr)/2
	i[0] = i[1]/2
	i[2] = i[1] + (len(*arr) - i[1])/2
	
	return &i
}

// Sorts the given array
func SortArray(arr []int, wg *sync.WaitGroup) {
	fmt.Printf("Sorting subarray %v\n", arr)
	sort.Ints(arr)
	wg.Done()
}

// Merge the four arrays
func MergeArray(a1, a2, a3, a4 []int, size int) []int {
	return Merge2Arrays(Merge2Arrays(a1, a2, len(a1)+len(a2)),
	Merge2Arrays(a3, a4, len(a3)+len(a4)), size)
}

// Uses merge algorithm similar to that in merge sort
func Merge2Arrays(a1, a2 []int, size int) []int {
	// initial empty array to store merged ints
	merged := make([]int, 0, size)

	// indices to keep track of smallest element in each array
	i, j := 0, 0

	for i < len(a1) && j < len(a2) {
		// find and add smallest element to merged array
		if a1[i] < a2[j] {
			merged = append(merged, a1[i])
			i++
		} else {
			merged = append(merged, a2[j])
			j++
		}
	}

	// add any remaining elements to the merged array
	merged = append(merged, a1[i:]...)
	merged = append(merged, a2[j:]...)

	return merged
}
