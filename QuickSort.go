package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
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

func loadArrayFromFile(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var array []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			return nil, fmt.Errorf("erro ao converter número no arquivo %s: %v", filename, err)
		}
		array = append(array, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return array, nil
}

func main() {
	numCores := runtime.NumCPU()

	// Processar arquivos de 10^1 a 10^7
	for exp := 1; exp <= 7; exp++ {
		size := 1
		for i := 0; i < exp; i++ {
			size *= 10
		}

		filename := fmt.Sprintf("array_%d.txt", size)
		array, err := loadArrayFromFile(filename)
		if err != nil {
			fmt.Printf("Erro ao carregar o arquivo %s: %v\n", filename, err)
			continue
		}

		for i := 1; i <= numCores; i++ {
			runtime.GOMAXPROCS(i)
			arrayCopy := make([]int, len(array))
			copy(arrayCopy, array) // Evitar modificar o array original

			var wg sync.WaitGroup
			wg.Add(1)

			start := time.Now()
			go parallelQuicksort(arrayCopy, 0, len(arrayCopy)-1, &wg)
			wg.Wait()
			duration := time.Since(start).Seconds() // Tempo em segundos

			// Imprimir no formato tamanhoDoArray;cores;tempo
			fmt.Printf("%d;%d;%.6f\n", size, i, duration)
		}
	}
}
