package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var re, _ = regexp.Compile(":(\\d+)")

func main() {

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3
	reader.Comma = ';'

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	newCsv := strings.Join(getCorrectCsv(records), "\n")
	before, after, _ := strings.Cut(os.Args[1], ".")

	newFile, err := os.Create(before + "_fixed." + after)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	_, _ = newFile.WriteString(newCsv)
}

func getCorrectCsv(records [][]string) []string {
	newCsv := make([]string, len(records))
	numbers := make([]int, 2)

	for index, row := range records {
		for _, match := range re.FindAllString(row[2], -1) {
			num, _ := strconv.Atoi(match[1:])
			numbers = append(numbers, num)
		}
		sum := 0
		for _, val := range numbers {
			sum += val
		}
		newCsv[index] = fmt.Sprintf("%s;%d", row[0], sum)
		numbers = numbers[:0]
	}
	return newCsv
}
