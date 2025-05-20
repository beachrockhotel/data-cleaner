package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
			writer.WriteString("|||") // если строка сломана
			writer.WriteString("\n")
			continue
		}

		name := fixName(fields[0])

		result := fmt.Sprintf("%s|%s|%s|%s\n", name)
		writer.WriteString(result)
	}
	writer.Flush()
}

func fixName(input string) string {
	input = strings.ReplaceAll(input, " ", "")
	re := regexp.MustCompile(`^([А-ЯЁа-яё]+)([А-ЯЁа-яё]+)$`)
	matches := re.FindStringSubmatch(input)
	if len(matches) == 3 {
		return capitalize(matches[1]) + " " + capitalize(matches[2])
	}
	return ""
}

func capitalize(word string) string {
	if len(word) == 0 {
		return ""
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}
