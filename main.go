package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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
		age := fixAge(fields[1])
		phone := fixPhone(fields[2])

		result := fmt.Sprintf("%s|%s|%s|%s\n", name, age, phone)
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

func fixAge(input string) string {
	input = strings.TrimSpace(input)
	re := regexp.MustCompile(`\d+`)
	match := re.FindString(input)
	if match == "" {
		return ""
	}
	age, err := strconv.Atoi(match)
	if err != nil || age < 1 || age > 120 {
		return ""
	}
	return strconv.Itoa(age)
}

func fixPhone(input string) string {
	re := regexp.MustCompile(`\d`)
	digits := strings.Join(re.FindAllString(input, -1), "")
	if len(digits) == 11 && strings.HasPrefix(digits, "7") {
		return fmt.Sprintf("+7 (%s) %s-%s-%s", digits[1:4], digits[4:7], digits[7:9], digits[9:])
	}
	return ""
}

func capitalize(word string) string {
	if len(word) == 0 {
		return ""
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}
