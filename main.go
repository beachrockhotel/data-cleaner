package main

import (
	"fmt"
	"os"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Ошибка при создании выходного файла:", err)
		return
	}
	defer outputFile.Close()

	fmt.Println("Файлы открыты успешно.")
}
