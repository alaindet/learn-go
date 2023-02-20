package main

import (
	"encoding/csv"
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

func csv2float(r io.Reader, column int) ([]float64, error) {

	var data []float64
	cr := csv.NewReader(r)
	allData, err := cr.ReadAll()

	if err != nil {
		return nil, fmt.Errorf("cannot read data from file: %w", err)
	}

	for i, row := range allData {

		// Skip headers row
		if i == 0 {
			continue
		}

		// Too few columns?!
		if len(row) <= column {
			err := fmt.Errorf("%w: File has only %d columns", ErrInvalidColumn, len(row))
			return nil, err
		}

		v, err := strconv.ParseFloat(row[column], 64)

		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrNotNumber, err)
		}

		data = append(data, v)
	}

	return data, nil
}
