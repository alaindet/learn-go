package main

import (
	"bufio"
	"os"
	"strconv"
)

func readCapitalsData(dataPath string) (map[string]int, error) {
	file, err := os.Open(dataPath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	capitals := map[string]int{}

	for scanner.Scan() {
		cityName := scanner.Text()
		scanner.Scan() // newline?
		population, _ := strconv.Atoi(scanner.Text())
		capitals[cityName] = population
	}

	return capitals, nil
}
