package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// Semente para geração de números aleatórios
	rand.Seed(time.Now().UnixNano())

	// Gerar arrays de tamanhos de 10¹ a 10⁵
	for exp := 1; exp <= 7; exp++ {
		size := int(pow(10, exp))
		array := generateArray(size)

		// Nome do arquivo baseado no tamanho do array
		filename := "array_" + strconv.Itoa(size) + ".txt"

		// Salvar array em arquivo
		err := saveArrayToFile(array, filename)
		if err != nil {
			fmt.Printf("Erro ao salvar o arquivo %s: %v\n", filename, err)
			continue
		}

		fmt.Printf("Arquivo %s gerado com sucesso.\n", filename)
	}
}

// Função para gerar um array de números inteiros aleatórios
func generateArray(size int) []int {
	array := make([]int, size)
	for i := range array {
		array[i] = rand.Intn(1000) // Número aleatório entre 0 e 999
	}
	return array
}

// Função para salvar um array em um arquivo
func saveArrayToFile(array []int, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, num := range array {
		_, err := file.WriteString(strconv.Itoa(num) + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

// Função para calcular a potência de base^exp
func pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}
