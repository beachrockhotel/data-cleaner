package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "|")

		if len(fields) != 4 {
			writer.WriteString("|||") // если формат неверный
			writer.WriteString("\n")
			continue
		}

		// Пока просто возвращаем поля как есть
		name := fields[0]
		age := fields[1]
		phone := fields[2]
		email := fields[3]

		result := fmt.Sprintf("%s|%s|%s|%s\n", name, age, phone, email)
		writer.WriteString(result)
	}
	writer.Flush()
}
