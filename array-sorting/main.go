package main

import (
	"bufio"
	"os"
	"fmt"
	"sort"
	"strings"
	"strconv"
	"sync"
)

func sortPart(part []int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Sorting subarray:", part)
	sort.Ints(part)
}

func merge(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return result
}


func main() {
	// read input
	fmt.Println("Enter integers separated by spaces:")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	fields := strings.Fields(line)
	nums := make([]int, len(fields))

	for i, f := range fields {
		nums[i], _ = strconv.Atoi(f)
	}

	// partition (4 parts)
	n := len(nums)
	size := (n + 3) / 4 // ceiling division

	parts := make([][]int, 0, 4)
	for i := 0; i < n; i += size {
		end := i + size
		if end > n {
			end = n
		}
		parts = append(parts, nums[i:end])
	}

	// sort each part concurrently
	var wg sync.WaitGroup
	wg.Add(len(parts))

	for _, part := range parts {
		go sortPart(part, &wg)
	}

	wg.Wait()

	// merge sorted parts
	sorted := parts[0]
	for i := 1; i < len(parts); i++ {
		sorted = merge(sorted, parts[i])
	}

	// output result
	fmt.Println("Final sorted array:", sorted)
}