package main

import (
    "fmt"
)

func NoRepeatedNumInArray(arrs []int32) int32 {
	if len(arrs) < 2 {
		return int32(len(arrs))
	}
	count := 1
	tmp := arrs[0]
	for i := 1; i < len(arrs); i++ {
		if tmp != arrs[i] {
			count++
			tmp = arrs[i]
		}
	}
	return int32(count)
}

func main() {
    labels := []int32 {1, 2, 1, 1, 1, 1, 3, 3, 1, 3, 3, 4, 3, 3, 4, 3, 4, 4, 1, 5, 1, 1, 4, 3, 6, 4, 3, 5, 7, 6, 4, 4}
    length := NoRepeatedNumInArray(labels)
    vec := make([][]int32, length+1)
    for i, label := range labels {
        vec[label] = append(vec[label], int32(i))
    }
    fmt.Println(vec)
}
