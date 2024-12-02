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

func parallelMergeSort(array []int, wg *sync.WaitGroup) []int {
	defer wg.Done()
	if len(array) < 2 {
		return array
	}
	wg.Add(2)
	first := parallelMergeSort(array[:len(array)/2], wg)
	second := parallelMergeSort(array[len(array)/2:], wg)
	return merge(first, second)
}

func merge(a []int, b []int) []int {
		final := []int{}
		i := 0
		j := 0
		for i < len(a) && j < len(b) {
				if a[i] < b[j] {
						final = append(final, a[i])
						i++
				} else {
						final = append(final, b[j])
						j++
				}
		}
		for ; i < len(a); i++ {
				final = append(final, a[i])
		}
		for ; j < len(b); j++ {
				final = append(final, b[j])
		}
		return final
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
			go parallelMergeSort(arrayCopy, &wg)
			wg.Wait()
			duration := time.Since(start).Seconds() // Tempo em segundos

			// Imprimir no formato tamanhoDoArray;cores;tempo
			fmt.Printf("%d;%d;%.6f\n", size, i, duration)
		}
	}
}
