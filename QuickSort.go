package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func parallelQuicksort(array []int, arrStart int, arrEnd int, wg *sync.WaitGroup) {
	defer wg.Done()

	if arrStart < arrEnd {
		pivot := partition(array, arrStart, arrEnd)

		wg.Add(2)
		go parallelQuicksort(array, arrStart, pivot-1, wg)
		go parallelQuicksort(array, pivot+1, arrEnd, wg)
	}
}

func partition(array []int, arrStart int, arrEnd int) int {
	pivot := array[arrEnd]
	i := arrStart - 1

	for j := arrStart; j < arrEnd; j++ {
		if array[j] <= pivot {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}

	array[i+1], array[arrEnd] = array[arrEnd], array[i+1]
	return i + 1
}

func populateArray(size int, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())

	array := make([]int, size)

	for i := 0; i < size; i++ {
		array[i] = rand.Intn(maxValue)
	}

	return array
}

func main() {
	numCores := runtime.NumCPU()
	k := 1
	s := 10
	for k <= 5 {
		fmt.Printf("Tamanho do array: %d\n", s)
		fmt.Printf("Cores;Tempo\n")
		for i := 1; i <= numCores; i++ {
			runtime.GOMAXPROCS(i)
			array := populateArray(s, 500000000)
			var wg sync.WaitGroup
			wg.Add(1)
			start := time.Now()
			go parallelQuicksort(array, 0, len(array)-1, &wg)
			wg.Wait()
			duration := time.Since(start)
			fmt.Printf("%d;%v\n", i, duration)
		}
		s = s * 10
	}
}
