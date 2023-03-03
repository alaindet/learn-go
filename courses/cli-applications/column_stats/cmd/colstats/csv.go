package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
)

type statsFunc func(data []float64) float64

func sum(data []float64) float64 {
	sum := 0.0
	for _, v := range data {
		sum += v
	}
	return sum
}

func avg(data []float64) float64 {
	return sum(data) / float64(len(data))
}

// Reads a whole .csv file and return a given numerical column as []float64
func csv2float(r io.Reader, column int) ([]float64, error) {

	var data []float64
	cr := csv.NewReader(r)
	cr.ReuseRecord = true

	for i := 0; ; i++ {

		row, err := cr.Read()

		// Reached end of file, exit loop
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("cannot read data from file: %w", err)
		}

		// Skip headers row
		if i == 0 {
			continue
		}

		// Error: Too few columns
		if len(row) <= column {
			err := fmt.Errorf("%w: File has only %d columns", ErrInvalidColumn, len(row))
			return nil, err
		}

		v, err := strconv.ParseFloat(row[column], 64)

		// Error: Not a number
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrNotNumber, err)
		}

		data = append(data, v)
	}

	return data, nil
}
